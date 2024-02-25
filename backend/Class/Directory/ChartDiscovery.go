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
				Logger.Info("Event trigger - " + event.Op.String() + " on" + event.Name)

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
		Logger.Info("Action - delete")
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
			if header.Name == "Chart.yaml" || header.Name == "Chart.yml" {
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
