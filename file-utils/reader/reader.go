package reader

import (
	"errors"
	"strings"
)

var InputSchemesMap = map[string]bool{
	dirPrefix:   true,
	FilePrefix:  true,
	HttpsPrefix: true,
	HttpPrefix:  true,
	SshPrefix:   true,
	GitHttpsPrefix: true,
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
	switch {
	case strings.HasPrefix(filePath, dirPrefix):
		return ReadDir, nil
	case strings.HasPrefix(filePath, FilePrefix):
		return ReadFile, nil
	case strings.HasPrefix(filePath, HttpsPrefix):
		return ReadHttps, nil
	case strings.HasPrefix(filePath, HttpPrefix):
		return ReadHttps, nil
	case strings.HasPrefix(filePath, SshPrefix):
		return ReadSsh, nil
	case strings.HasPrefix(filePath, GitHttpsPrefix):
		return ReadGitHttps, nil
	default:
		return nil, errors.New("invalid reader scheme")
	}
}
