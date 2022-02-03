package file

import (
	"io/ioutil"
	"strings"
)

const (
	Prefix = "file://"
)

func ParseFile(filePath string) (interface{}, error) {
	filePath = strings.TrimPrefix(filePath, Prefix)
	file, err := ioutil.ReadFile(filePath)
	return file, err
}
