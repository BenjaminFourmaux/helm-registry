package Utils

import (
	"path/filepath"
	"strings"
)

func ConvertWindowsPathToUnix(windowsPath string) string {
	// return filepath.ToSlash(windowsPath) doesn't work
	return strings.Replace(windowsPath, "\\", "/", -1)
}

/*
GetFilenameFromPath From a file relative path, get le file name (with extension)
tested & approved
*/
func GetFilenameFromPath(path string) string {
	return filepath.Base(ConvertWindowsPathToUnix(path))
}

func GenerateChartPath(filename string) string {
	return "/charts/" + filename
}
