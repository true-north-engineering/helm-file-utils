package reader

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

const (
	DirPrefix = "dir://"
)

func ReadDir(filePath string) (interface{}, error) {
	filePath = strings.TrimPrefix(filePath, DirPrefix)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() {
		return nil, errors.Errorf("Specified path is not a directory")
	}

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]byte)
	for _, fileInfo := range fileInfos {
		parsedFile, err := ReadFile(filePath + "/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		result[fileInfo.Name()] = parsedFile.([]byte)
	}

	return result, nil
}
