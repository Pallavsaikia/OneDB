package filesys

import (
	"fmt"
	"runtime"
	"path"
	"path/filepath"
)

func GetFilePath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "Unable to determine file location"
	}
	absolutePath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Sprintf("Error getting absolute path: %v", err)
	}
	return absolutePath
}


func GetFileLocation() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("unable to determine file location")
	}
	dir := path.Dir(filename)
	return dir, nil
}
