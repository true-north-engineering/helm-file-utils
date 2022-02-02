package file

import (
	"io/ioutil"
	"strings"
)

const (
	Prefix = "file://"
)

func ParseFile(filePath string) (string, error) {
	filePath = strings.TrimPrefix(filePath, Prefix)
	file, err := ioutil.ReadFile(strings.TrimPrefix(filePath, Prefix))
	return string(file), err
}
