package reader

import (
	"strings"

	"github.com/pkg/errors"
)

type Factory func(filePath string) (interface{}, error)

func DetermineReader(filePath string) (Factory, string, error) {
	switch {
	case strings.Contains(filePath, DirPrefix):
		return ReadDir, "", nil
	case strings.Contains(filePath, FilePrefix):
		return ReadFile, "", nil
	}
	if !strings.Contains(filePath, "://") {
		return ReadFile, "", nil
	}
	return nil, "", errors.Errorf("error while parsing filepath %s with file utils plugin", filePath)
}
