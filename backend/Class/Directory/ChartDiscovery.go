package Directory

import (
	"backend/Class/Database"
	"backend/Class/Logger"
	"backend/Class/Utils"
	"backend/Class/Utils/env"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"helm.sh/helm/v3/pkg/chart/loader"
	"os"
	"path/filepath"
	"strconv"
)

/*
Discovery Browse all compressed file in the charts directory and check if all charts are in database.
Otherwise, add them
*/
func Discovery() {
	Logger.Info("Discovering charts")

	// 0. Get files from REPOSITORY_DIR
	files, err := os.ReadDir(env.REPOSITORY_DIR)
	if err != nil {
		Logger.Error("Unable to open Charts Directory")
		Logger.Raise(err)
	}

	// 1. Check for deleted chart file
	checkRemovedChartFile(files)

	// 2. Get all charts in db to him with charts in directory
	chartsRows, _ := Database.GetAllCharts()
	chartsInDB := Utils.ParserRowsToChartDTO(chartsRows)

	// 3. Browse zip file content
	for _, file := range files {
		if !file.IsDir() && IsTGZArchive(file.Name()) {
			archive, err := os.Open(filepath.Join(env.REPOSITORY_DIR, file.Name()))
			if err != nil {
				Logger.Error("Unable to open tar archive")
				Logger.Raise(err)
			}

			chartInDir, err := loader.LoadArchive(archive)
			if err != nil {
				Logger.Raise(err)
			}

			path := Utils.GenerateChartPath(file.Name())
			chartDigest := GetDigestFromIndexFile(chartInDir)

			// 4. Check if chartInDir is a newer chart from the db
			if IsOnList(chartInDir, chartsInDB) {
				// Chart already exist in db. Check for changes
				chartInDB := GetOnList(chartInDir, chartsInDB)
				if chartDigest != chartInDB.Digest {
					// Need update chart
					chartDTO := Utils.ParserChartToDTO(chartInDir, chartDigest, path)
					Database.UpdateChart(chartInDB.Id, chartDTO)
				}
			} else {
				// The database is empty, add the chart
				chartDTO := Utils.ParserChartToDTO(chartInDir, chartDigest, path)
				Database.InsertChart(chartDTO)
			}
			defer archive.Close()
		}
	}
}

/*
RepositoryDirectoryWatcher Initialize the Directory Watcher Listener and call appropriate functions when Events throw
*/
func RepositoryDirectoryWatcher() {
	// Create a Watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		Logger.Error("Unable to create folder watcher")
	}

	// Add the watching folder
	err = watcher.Add(env.REPOSITORY_DIR)

	// Trigger event
	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				Logger.Info("Event trigger - " + event.Op.String() + " on " + event.Name)

				actionTrigger(event)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error", err)
			}
		}
	}()
}

/*
actionTrigger Chose an action by the event operation
*/
func actionTrigger(event fsnotify.Event) {
	switch event.Op.String() {
	case "CREATE":
		if IsTGZArchive(event.Name) && IsAChartPackage(event.Name) {
			Logger.Info("Action - insert")
			insertDBFromNewFile(event.Name)
		}
	case "REMOVE":
		if IsTGZArchive(event.Name) {
			Logger.Info("Action - delete")
			deleteDBFromRemoveFile(event.Name)
		}
	case "RENAME":
		if IsTGZArchive(event.Name) {
			// Check if file was removed
			if !IsFileExist(event.Name) {
				Logger.Info("Action - delete")
				deleteDBFromRemoveFile(event.Name)
			} else {
				// No effect (CREATE trigger before)
				// This case spawn in GitHub Action Worker (Ubuntu) but not in my pc (Windows)
				// I think rename is a specific case for Unix folder manager (move via os.rename GO)
			}
		}
	}
	// Update index.yaml file after action triggering and database change
	UpdateIndex()
}

/*
checkRemovedChartFile Browse all files of the charts directory, get only .tgz file and add them into a list of file
name and check if in the db, all charts entry has equivalent in directory. Otherwise, delete him
*/
func checkRemovedChartFile(files []os.DirEntry) {
	Logger.Info("Discovering charts - Check for removed chart files")

	var filenames []string
	var chartsIdToDelete []int

	// 1. Browse all files in directory and add them to a list of file present
	for _, file := range files {
		if !file.IsDir() && IsTGZArchive(file.Name()) {
			filenames = append(filenames, file.Name())
		}
	}

	// 2. Get all charts in db
	chartsInDB, _ := Database.GetAllCharts()
	listALlChartsDTO := Utils.ParserRowsToChartDTO(chartsInDB)

	// 3. Browse all charts in db and get chart id not in directory
	for _, chart := range listALlChartsDTO {
		var isOnDirectory = false

		// 4. Browse chart's urls
		if IsFilenameInDirectoryFiles(filepath.Base(Utils.NullToString(chart.Path)), filenames) {
			isOnDirectory = true
		}

		// 5. Add id of chart if not in directory (a deleted file chart)
		if !isOnDirectory {
			chartsIdToDelete = append(chartsIdToDelete, chart.Id)
		}
	}

	// 6. Delete removed chart
	err := Database.DeleteCharts(chartsIdToDelete)
	if err != nil {
		Logger.Error("When deleted removed charts")
		Logger.Raise(err)
	}
}

/*
insertDBFromNewFile Send to BD info of a new chart creating in the repository directory
*/
func insertDBFromNewFile(filepath string) {
	// Open tar archive
	file, err := os.Open(Utils.ConvertWindowsPathToUnix(filepath))
	if err != nil {
		Logger.Error("Error unable to open .tgz archive")
		return
	}

	chartInDir, err := loader.LoadArchive(file)
	if err != nil {
		// If the file is not a chart archive, err must be "not a chart archive"
		Logger.Raise(err)
		return
	}

	path := Utils.GenerateChartPath(Utils.GetFilenameFromPath(file.Name()))
	chartDigest := GetDigestFromIndexFile(chartInDir)

	// Insert chart in db
	chartDTO := Utils.ParserChartToDTO(chartInDir, chartDigest, path)
	Database.InsertChart(chartDTO)

	defer file.Close()
}

/*
deleteDBFromRemoveFile Delete on the DB when a .tar file is removed
*/
func deleteDBFromRemoveFile(filepath string) {
	result := Database.GetChartByFilename(Utils.GenerateChartPath(Utils.GetFilenameFromPath(filepath)))

	if result.Err() != nil {
		Logger.Warning(result.Err().Error())
	} else {
		var chartToDelete = Utils.ParserRowToChartDTO(result)

		Logger.Info("Delete Chart id: " + strconv.Itoa(chartToDelete.Id))

		_, err := Database.DeleteChart(chartToDelete.Id)
		if err != nil {
			Logger.Error("Unable to delete Chart by id: " + strconv.Itoa(chartToDelete.Id))
		}
	}
}
