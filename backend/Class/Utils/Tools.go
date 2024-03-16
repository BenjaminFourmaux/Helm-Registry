package Utils

import (
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

func GenerateChartUrls(filename string) []string {
	return []string{
		"/charts/" + filename,
	}
}
