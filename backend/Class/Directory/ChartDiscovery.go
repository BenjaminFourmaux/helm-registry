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
	"strconv"
)

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

				ActionTrigger(event)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error", err)
			}
		}
	}()
}

// ActionTrigger Chose an action by the event operation
func ActionTrigger(event fsnotify.Event) {
	switch event.Op.String() {
	case "CREATE":
		if IsATGZFile(event.Name) {
			Logger.Info("Action - insert")
			InsertDBFromNewFile(event.Name)
		}
	case "REMOVE":
		if IsATGZFile(event.Name) {
			Logger.Info("Action - delete")
			DeleteDBFromRemoveFile(event.Name)
		}
	}
	// Update index.yaml file after action triggering and database change
	UpdateIndex()
}

// InsertDBFromNewFile Send to BD info of a new chart creating in the repository directory
func InsertDBFromNewFile(filepath string) {
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
			if Utils.GetFilenameFromPath(header.Name) == "Chart.yaml" || Utils.GetFilenameFromPath(header.Name) == "Chart.yml" {
				// Read the content of the file and unmarshal it in yaml format
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, tarReader); err != nil {
					Logger.Error("Error when reading Chart.yaml file")
					return
				}
				var dataFile Entity.ChartFile
				err := yaml.Unmarshal(buf.Bytes(), &dataFile)
				if err != nil {
					Logger.Error("Error in the YAML file, unable to deserialize")
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

// DeleteDBFromRemoveFile Delete on the DB when a .tar file is removed
func DeleteDBFromRemoveFile(filepath string) {
	result := Database.GetChartByFilename(Utils.GetFilenameFromPath(filepath))
	
	fmt.Println(result)

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
