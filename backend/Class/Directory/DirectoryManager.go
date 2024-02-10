package Directory

import (
	"backend/Class/Database"
	"backend/Class/Logger"
	"backend/Entity"
	"gopkg.in/yaml.v2"
	"os"
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

	// Step 3. Build the file
	index := Entity.Index{
		APIVersion: "v1",
		Entries:    make(map[string][]Entity.ChartEntry),
		Generated:  time.Now(),
	}

	// Step 2. Foreach rows
	for rows.Next() {
		var entry Entity.DTORegistry

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

	yamlData, _ := yaml.Marshal(&index)

	// Step 3. Save index YAML file
	SaveFile(filePath, yamlData)

	Logger.Success("Index YAML file successfully updated")
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
