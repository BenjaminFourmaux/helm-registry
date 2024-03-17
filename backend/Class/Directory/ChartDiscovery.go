package Directory

import (
	"archive/tar"
	"backend/Class/Database"
	"backend/Class/Logger"
	"backend/Class/Utils"
	"backend/Class/Utils/env"
	"backend/Entity"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

// Discovery Browse all compressed file in the charts directory and check if all charts are in database.
// Otherwise, add them
func Discovery() {
	Logger.Info("Discovering charts")

	files, err := os.ReadDir(env.CHARTS_DIR)
	if err != nil {
		Logger.Error("Unable to open Charts Directory")
		Logger.Raise(err.Error())
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".tgz" {
			// 1. Open .tgz archive
			archive, err := os.Open(filepath.Join(env.CHARTS_DIR, file.Name()))
			if err != nil {
				Logger.Error("Unable to open tar archive")
				Logger.Raise(err.Error())
			}
			defer archive.Close()

			uncompressedFile, err := gzip.NewReader(archive)
			tarReader := tar.NewReader(uncompressedFile)

			// 2. Browse zip file content
			for {
				header, err := tarReader.Next()
				if err == io.EOF {
					break
				}
				if header.Typeflag == tar.TypeReg {
					if IsChartFile(Utils.GetFilenameFromPath(header.Name)) {
						// 3. Extract chart infos from chart YAML file
						var buf bytes.Buffer
						if _, err := io.Copy(&buf, tarReader); err != nil {
							Logger.Error("Error when reading Chart.yaml file")
							return
						}
						var dataFile Entity.ChartFile
						err := yaml.Unmarshal(buf.Bytes(), &dataFile)
						if err != nil {
							Logger.Error("Error in the YAML file, unable to deserialize it")
						}

						// 4. Create the DTO entity with the data from file
						urls := Utils.GenerateChartUrls(Utils.GetFilenameFromPath(file.Name()))
						var dto = Utils.ParserChartToDTO(dataFile, urls)

						// 5. Check if chart already exist in the database
						if Database.IfChartExist(dto) {
							// 6.a Update chart in db
							var chartId = Utils.ParserRowToChartDTO(Database.GetChartByCriteria(dto)).Id
							Database.UpdateChart(chartId, dto)

						} else {
							// 6.b Insert to the database
							Database.InsertChart(dto)
						}

						break
					}
				}
			}

		}
	}
}

// RepositoryDirectoryWatcher Initialize the Directory Watcher Listener and call appropriate functions when Events throw
func RepositoryDirectoryWatcher() {
	// Create a Watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		Logger.Error("Unable to create folder watcher")
	}

	// Add the watching folder
	err = watcher.Add(env.CHARTS_DIR)

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

// actionTrigger Chose an action by the event operation
func actionTrigger(event fsnotify.Event) {
	switch event.Op.String() {
	case "CREATE":
		if IsATGZFile(event.Name) {
			Logger.Info("Action - insert")

			insertDBFromNewFile(event.Name)

			// Update index.yaml file after action triggering and database change
			UpdateIndex()
		}
	case "REMOVE":
		if IsATGZFile(event.Name) {
			Logger.Info("Action - delete")
			deleteDBFromRemoveFile(event.Name)
		}
	}
	// Update index.yaml file after action triggering and database change
	UpdateIndex()
}

// insertDBFromNewFile Send to BD info of a new chart creating in the repository directory
func insertDBFromNewFile(filepath string) {
	// Open tar archive
	file, err := os.Open(Utils.ConvertWindowsPathToUnix(filepath))
	if err != nil {
		Logger.Error("Error unable to open .tgz archive")
		return
	}

	defer file.Close()

	// uncompressed file
	uncompressedFile, err := gzip.NewReader(file)
	if err != nil {
		Logger.Error("Error uncompressed archive file")
	}

	// Create the archive reader
	tarReader := tar.NewReader(uncompressedFile)

	// Check if is a Helm Chart package
	//if !IsAChartPackage(tarReader) {
	//return
	//}

	// Browse archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if header.Typeflag == tar.TypeReg {
			if IsChartFile(Utils.GetFilenameFromPath(header.Name)) {
				// Read the content of the file and unmarshal it in yaml format
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, tarReader); err != nil {
					Logger.Error("Error when reading Chart.yaml file")
					return
				}
				var dataFile Entity.ChartFile
				err := yaml.Unmarshal(buf.Bytes(), &dataFile)
				if err != nil {
					Logger.Error("Error in the YAML file, unable to deserialize it")
				}

				// Create the DTO entity with the data from file
				urls := Utils.GenerateChartUrls(Utils.GetFilenameFromPath(file.Name()))
				var dto = Utils.ParserChartToDTO(dataFile, urls)

				// Insert to the database
				Database.InsertChart(dto)
				break
			}
		}
	}
}

// deleteDBFromRemoveFile Delete on the DB when a .tar file is removed
func deleteDBFromRemoveFile(filepath string) {
	result := Database.GetChartByFilename(Utils.GetFilenameFromPath(filepath))

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
