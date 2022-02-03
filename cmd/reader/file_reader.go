package reader

import (
	"io/ioutil"
	"strings"
)

const (
	FilePrefix = "file://"
)

func ReadFile(filePath string) (interface{}, error) {
	filePath = strings.TrimPrefix(filePath, FilePrefix)
	file, err := ioutil.ReadFile(filePath)
	return file, err
}
