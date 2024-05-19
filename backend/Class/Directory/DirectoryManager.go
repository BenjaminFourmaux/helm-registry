package Directory

import (
	"archive/tar"
	"backend/Class/Database"
	"backend/Class/Logger"
	"backend/Class/Utils"
	"backend/Class/Utils/env"
	"backend/Entity"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func UpdateIndex() {
	filePath := env.INDEX_FILE_PATH

	Logger.Info("Updating Index")

	// Step 1. Get registry info from Database
	rows, errSql := Database.GetALlChartsOrderedByName()
	if errSql != nil {
		Logger.Error("Unable to get data from Database")
	}

	// Step 2. Build the file
	index := Entity.Index{
		APIVersion: "v1",
		Entries:    make(map[string][]Entity.ChartEntry),
		Generated:  time.Now(),
	}

	// Step 3. Foreach rows
	// TODO: refactor to use a PARSER from Utils
	for rows.Next() {
		var entry Entity.ChartDTO

		if err := rows.Scan(&entry.Id, &entry.Name, &entry.Description, &entry.Version, &entry.Created, &entry.Digest,
			&entry.Home, &entry.Sources, &entry.Urls); err != nil {
			Logger.Error("Deserialization data -> dto")
		}

		// Create entry
		chartEntry := Entity.ChartEntry{
			Version:     entry.Version,
			Created:     entry.Created,
			Name:        entry.Name,
			Description: Utils.NullToString(entry.Description),
			Digest:      entry.Digest,
			Home:        Utils.NullToString(entry.Home),
			Sources:     strings.Split(Utils.NullToString(entry.Sources), ";"),
			Urls:        strings.Split(entry.Urls, ";"),
		}

		// Add entry in file content
		index.Entries[entry.Name] = append(index.Entries[entry.Name], chartEntry)
	}

	// Step 4. Check if change needed
	yamlFile := &Entity.Index{}
	err := yaml.Unmarshal(ReadFile(filePath), yamlFile)
	if err != nil {
		Logger.Error("Unable to unmarshal the index file")
	}

	if CheckChange(yamlFile, &index) {
		index.Generated = time.Now()
		yamlData, _ := yaml.Marshal(&index)

		// Step 5. Save index YAML file
		SaveFile(filePath, yamlData)

		Logger.Success("Index successfully updated")
	} else {
		Logger.Info("Index - No change needed")
	}
}

func ReadFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		Logger.Error("Unable to open file")
	}

	return file
}

func SaveFile(filePath string, data []byte) {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		Logger.Error("Fail to save file : " + filePath)
	}
}

// CheckChange Compare oldYaml with newYaml and return true if there are a change
func CheckChange(oldYaml *Entity.Index, newYaml *Entity.Index) bool {
	// Remove 'generated' field
	oldYaml.Generated = time.Time{}
	newYaml.Generated = time.Time{}

	return !reflect.DeepEqual(*oldYaml, *newYaml)
}

// IsTGZArchive Return true if the file (or path+file) extension is .tgz
func IsTGZArchive(path string) bool {
	return filepath.Ext(path) == ".tgz"
}

// IsAChartPackage Check if in the zip has the requirement to be a Helm Chart (Chart.yaml)
func IsAChartPackage(fileReader *tar.Reader) bool {
	fmt.Println(fileReader)
	for {
		header, err := fileReader.Next()
		if err == io.EOF {
			break
		}
		if header.Typeflag == tar.TypeReg {
			if header.Name == "Chart.yaml" || header.Name == "Chart.yml" {
				return true
			}
		}
	}
	return false
}

// IsChartFile Return true if the filename match with Helm chart file naming rule
func IsChartFile(filename string) bool {
	return filename == "Chart.yaml" || filename == "Chart.yml" || filename == "chart.yaml" || filename == "chart.yml"
}

// IsFilenameInDirectoryFiles Return true if the filename is on the directory (list of present filename)
func IsFilenameInDirectoryFiles(filename string, list []string) bool {
	for _, item := range list {
		if item == filename {
			return true
		}
	}
	return false
}

// CreateDirIfNotExist Create a new dir if not exist
func CreateDirIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		err := os.MkdirAll(path, 0755)

		if err != nil {
			Logger.Error("Error creating directory : " + path)
		} else {
			Logger.Success("Creating new directory on : " + path)
		}
	}
}
