package transformer

import (
	"encoding/base64"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
)

const (
	B64DecPrefix = "base64dec"
)

func B64DECTransform(inputValue reader.InputValue) (reader.InputValue, error) {
	result := reader.InputValue{Kind: inputValue.Kind, Value: make(map[string][]byte)}

	if inputValue.Kind == reader.InputKindFile {
		inputFile := inputValue.Value[reader.InputKindFile]
		decodedFile, err := base64.StdEncoding.DecodeString(string(inputFile))
		if err != nil {
			return reader.InputValue{}, err
		}
		result.Value[reader.InputKindFile] = decodedFile
	} else if inputValue.Kind == reader.InputKindDir {
		inputFiles := inputValue.Value
		for fileName, fileValue := range inputFiles {
			decodedFile, err := base64.StdEncoding.DecodeString(string(fileValue))
			if err != nil {
				return reader.InputValue{}, err
			}
			result.Value[reader.InputKindFile] = decodedFile
			result.Value[fileName] = decodedFile
		}
	}
	return result, nil
}
