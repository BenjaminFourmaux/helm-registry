package Directory

import (
	"archive/tar"
	"backend/Class/Database"
	"backend/Class/Logger"
	"backend/Entity"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

func UpdateIndex() {
	filePath := os.Getenv("INDEX_FILE_PATH")

	// Step 1. Get registry info from Database
	rows, errSql := Database.GetALlChartsOrderedByName()
	if errSql != nil {
		Logger.Error("Enable to get data from Database")
	}

	// Step 2. Build the file
	index := Entity.Index{
		APIVersion: "v1",
		Entries:    make(map[string][]Entity.ChartEntry),
		Generated:  time.Now(),
	}

	// Step 3. Foreach rows
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
			Description: entry.Description,
			Digest:      entry.Digest,
			Home:        entry.Home,
			Sources:     strings.Split(entry.Sources, ";"),
			Urls:        strings.Split(entry.Urls, ";"),
		}

		// Add entry in file content
		index.Entries[entry.Name] = append(index.Entries[entry.Name], chartEntry)
	}

	// Step 4. Check if change needed
	yamlFile := &Entity.Index{}
	err := yaml.Unmarshal(ReadFile(filePath), yamlFile)
	if err != nil {
		Logger.Error("Impossible to unmarshal the index file")
	}

	if CheckChange(yamlFile, &index) {
		index.Generated = time.Now()
		yamlData, _ := yaml.Marshal(&index)

		// Step 5. Save index YAML file
		SaveFile(filePath, yamlData)

		Logger.Success("Index YAML file successfully updated")
	}
}

func ReadFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		Logger.Error("Enable to open file")
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

func IsATGZFile(path string) bool {
	return strings.HasSuffix(path, ".tgz")
}

// IsAChartPackage Check if in the zip has the requirement to be a Helm Chart (Chart.yaml)
func IsAChartPackage(fileReader *tar.Reader) bool {
	Logger.Debug("issou")
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
