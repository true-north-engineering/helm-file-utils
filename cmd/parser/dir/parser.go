package dir

import (
	"os"
	"strings"

	"github.com/pkg/errors"
	fileParser "github.com/true-north-engineering/helm-file-utils/cmd/parser/file"
)

const (
	Prefix = "dir://"
)

func ParseDir(filePath string) (interface{}, error) {
	filePath = strings.TrimPrefix(filePath, Prefix)
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
		parsedFile, err := fileParser.ParseFile(filePath + "/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		result[fileInfo.Name()] = parsedFile.([]byte)
	}

	return result, nil
}
