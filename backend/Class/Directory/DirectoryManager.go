package Directory

import (
	"archive/tar"
	"backend/Class/Logger"
	"backend/Class/Utils/env"
	"backend/Entity"
	"fmt"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/repo"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"time"
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
CheckChange Compare oldYaml with newYaml and return true if there are a change
*/
func CheckChange(oldYaml *Entity.Index, newYaml *Entity.Index) bool {
	// Remove 'generated' field
	oldYaml.Generated = time.Time{}
	newYaml.Generated = time.Time{}

	return !reflect.DeepEqual(*oldYaml, *newYaml)
}

/*
IsTGZArchive Return true if the file (or path+file) extension is .tgz
*/
func IsTGZArchive(path string) bool {
	return filepath.Ext(path) == ".tgz"
}

/*
IsAChartPackage Check if in the zip has the requirement to be a Helm Chart (Chart.yaml)
*/
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
