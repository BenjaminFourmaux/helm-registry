package Utils

import "path/filepath"

func ConvertWindowsPathToUnix(windowsPath string) string {
	return filepath.ToSlash(windowsPath)
}

func GenerateChartUrls(filename string) []string {
	return []string{
		"/charts/" + filename,
	}
}
