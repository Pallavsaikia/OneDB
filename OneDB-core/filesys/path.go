package filesys

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func IsValidPath(path string) bool {
	isAbsolute := filepath.IsAbs(path)
	if isAbsolute {
		_, err := os.Stat(path)
		return err == nil
	}
	return false
}

func CreateFileAndPathIfNotExist(filePath string) (*os.File, error) {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("error creating directories: %v", err)
	}
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error creating directories: %v", err)
	}
	return file, nil
}

func CreatePathFromStringArray(paths []string) string {
	str := ""
	for i, path := range paths {
		if i == 0 {
			str = path
		} else {
			str = str + "\\" + path
		}

	}
	// fmt.Println(str)
	return str
}
