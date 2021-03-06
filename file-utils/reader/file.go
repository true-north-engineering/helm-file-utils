package reader

import (
	"io/ioutil"
	"strings"
)

const (
	FilePrefix = "file"
	NoPrefix   = ""
)

// ReadFile Reader protocol that reads content from a file no matter the extension.
func ReadFile(filePath string) (InputValue, error) {
	file, err := ioutil.ReadFile(strings.TrimPrefix(filePath, FilePrefix+"://"))
	result := InputValue{Kind: InputKindFile, Value: make(map[string][]byte)}
	result.Value[InputKindFile] = file
	return result, err
}
