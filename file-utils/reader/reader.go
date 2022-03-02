package reader

import (
	"errors"
	"strings"
)

var InputSchemesMap = map[string]bool{
	dirPrefix:  true,
	FilePrefix: true,
	SshPrefix:  true,
}

const (
	InputKindFile = "file"
	InputKindDir  = "dir"
)

type InputKind string

type InputValue struct {
	Kind  InputKind
	Value map[string][]byte
}

type Factory func(filePath string) (InputValue, error)

func DetermineReader(filePath string) (Factory, error) {
	if strings.HasPrefix(filePath, dirPrefix) {
		return ReadDir, nil
	} else if strings.HasPrefix(filePath, FilePrefix) {
		return ReadFile, nil
	} else if strings.HasPrefix(filePath, SshPrefix) {
		return ReadSsh, nil
	} else {
		return nil, errors.New("invalid reader scheme")
	}
}
