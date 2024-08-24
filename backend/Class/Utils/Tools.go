package Utils

import (
	"backend/Entity"
	"path/filepath"
	"strings"
)

func ConvertWindowsPathToUnix(windowsPath string) string {
	// return filepath.ToSlash(windowsPath) doesn't work
	return strings.Replace(windowsPath, "\\", "/", -1)
}

func GetFilenameFromPath(path string) string {
	return filepath.Base(ConvertWindowsPathToUnix(path))
}

func GenerateChartPath(filename string) string {
	return "/charts/" + filename
}

func IsChartAlreadyExist(charts []Entity.ChartDTO, chartToCheck Entity.ChartDTO) (bool, int) {
	for _, chart := range charts {
		if chart.Name == chartToCheck.Name && chart.Version == chartToCheck.Version && NullToString(chart.Path) == NullToString(chartToCheck.Path) {
			return true, chart.Id
		}
	}
	return false, 0
}
