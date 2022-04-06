package transformer

import (
	"encoding/base64"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
)

const (
	B64EncPrefix = "base64enc"
)

// B64ENCTransform Transformer that decodes given base64 encoded data into string format.
func B64ENCTransform(inputValue reader.InputValue) (reader.InputValue, error) {
	result := reader.InputValue{Kind: inputValue.Kind, Value: make(map[string][]byte)}
	if inputValue.Kind == reader.InputKindFile {
		inputFile := inputValue.Value[reader.InputKindFile]
		encodedFile := make([]byte, base64.StdEncoding.EncodedLen(len(inputFile)))
		base64.StdEncoding.Encode(encodedFile, inputFile)
		result.Value[reader.InputKindFile] = encodedFile
	} else if inputValue.Kind == reader.InputKindDir {
		inputFiles := inputValue.Value
		for fileName, fileValue := range inputFiles {
			encodedFile := make([]byte, base64.StdEncoding.EncodedLen(len(fileValue)))
			base64.StdEncoding.Encode(encodedFile, fileValue)
			result.Value[fileName] = encodedFile
		}
	}
	return result, nil
}
