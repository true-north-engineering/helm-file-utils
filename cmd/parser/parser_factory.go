package parser

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/cmd/parser/base64"
	"github.com/true-north-engineering/helm-file-utils/cmd/parser/file"
)

type Factory func(filePath string) (string, error)

func DetermineParser(filePath string) (Factory, error) {
	switch {
	case strings.HasPrefix(filePath, base64.ENCPrefix):
		return base64.ParseFile, nil
	case strings.HasPrefix(filePath, FUTLPrefix):
		return ParseFile, nil
	case strings.HasPrefix(filePath, file.Prefix):
		return file.ParseFile, nil
	}
	if !strings.Contains(filePath, "://") {
		return file.ParseFile, nil
	}
	return nil, errors.Errorf("error while parsing filepath %s with file utils plugin", filePath)
}
