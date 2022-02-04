package reader

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

const (
	DirPrefix = "dir"
)

func ReadDir(dir string) (InputValue, error) {
	path := strings.TrimPrefix(dir, DirPrefix+"://")
	file, err := os.Open(path)
	if err != nil {
		return InputValue{}, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return InputValue{}, err
	}
	if !fileInfo.IsDir() {
		return InputValue{}, errors.Errorf("Specified path is not a directory")
	}

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return InputValue{}, err
	}
	result := InputValue{Kind: InputKindDir, Value: make(map[string][]byte)}
	for _, fileInfo := range fileInfos {
		parsedFile, err := ReadFile(path + "/" + fileInfo.Name())
		if err != nil {
			return InputValue{}, err
		}
		result.Value[fileInfo.Name()] = parsedFile.Value[InputKindFile]
	}

	return result, nil
}
