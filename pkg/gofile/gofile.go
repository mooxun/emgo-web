package gofile

import (
	"io/ioutil"
	"os"
)

// Create File
func CreateFile(filePath string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filePath, data, perm)
}

// FileExists file_exists()
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}