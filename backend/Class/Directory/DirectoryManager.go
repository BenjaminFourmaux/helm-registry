package Directory

import (
	"backend/Class/Logger"
	"backend/Class/Utils/env"
	"backend/Entity"
	"errors"
	"fmt"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/repo"
	"os"
	"path/filepath"
)

/*
UpdateIndex Update index.yaml using Helm SDK
*/
func UpdateIndex() {
	Logger.Info("Updating index.yaml file")

	// Get all charts in REPOSITORY_DIR (/charts), extract their information and build an index
	indexFile, err := repo.IndexDirectory(env.REPOSITORY_DIR, fmt.Sprintf("%s://%s:%d/charts", env.Scene, env.Hostname, env.Port))
	if err != nil {
		Logger.Error("When getting charts")
		Logger.Raise(err)
	}

	err = indexFile.WriteFile(env.INDEX_FILE_PATH, os.FileMode(0777))
	if err != nil {
		Logger.Error("Writing index.yaml file")
		Logger.Raise(err)
	}
}

/*
GetDigestFromIndexFile Get index information from REPOSITORY_DIR and get digest from selected chart
*/
func GetDigestFromIndexFile(chart *chart.Chart) string {
	indexFile, err := repo.IndexDirectory(env.REPOSITORY_DIR, fmt.Sprintf("%s://%s:%d/charts", env.Scene, env.Hostname, env.Port))
	if err != nil {
		Logger.Error("When getting charts")
		Logger.Raise(err)
	}

	if indexFile.Has(chart.Name(), chart.Metadata.Version) {
		indexInfo, _ := indexFile.Get(chart.Name(), chart.Metadata.Version)
		return indexInfo.Digest
	}
	return ""
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

/*
IsTGZArchive Return true if the file (or path+file) extension is .tgz
*/
func IsTGZArchive(path string) bool {
	return filepath.Ext(path) == ".tgz"
}

/*
IsFileExist Check if a file exist
*/
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

/*
IsAChartPackage Check if the file is a Chart archive using LoadArchive from Helm SDK
*/
func IsAChartPackage(pathFile string) bool {
	archive, err := os.Open(pathFile)
	if err != nil {
		Logger.Error("Can't open archive")
		Logger.Raise(err)
		return false
	}

	_, err = loader.LoadArchive(archive)
	defer archive.Close()
	if err != nil {
		Logger.Debug(err.Error())
		return false
	}

	return true
}

/*
IsChartFile Return true if the filename match with Helm chart file naming rule
*/
func IsChartFile(filename string) bool {
	return filename == "Chart.yaml" || filename == "Chart.yml" || filename == "chart.yaml" || filename == "chart.yml"
}

/*
IsFilenameInDirectoryFiles Return true if the filename is on the directory (list of present filename)
*/
func IsFilenameInDirectoryFiles(filename string, list []string) bool {
	for _, item := range list {
		if item == filename {
			return true
		}
	}
	return false
}

/*
IsOnList Search on ChartDTO list if a chart from file exist
*/
func IsOnList(chart *chart.Chart, list []Entity.ChartDTO) bool {
	for _, item := range list {
		if item.Name == chart.Name() && item.Version == chart.Metadata.Version {
			return true
		}
	}
	return false
}

/*
GetOnList Get on ChartDTO list match chart
*/
func GetOnList(chart *chart.Chart, list []Entity.ChartDTO) Entity.ChartDTO {
	for _, item := range list {
		if item.Name == chart.Name() && item.Version == chart.Metadata.Version {
			return item
		}
	}
	return Entity.ChartDTO{}
}
