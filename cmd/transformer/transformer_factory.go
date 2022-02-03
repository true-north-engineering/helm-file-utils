package transformer

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/cmd/reader"
)

type Factory func(filePath string) (interface{}, error)

func DetermineTransformer(filePath string) (Factory, error) {
	switch {
	case strings.HasPrefix(filePath, B64EncPrefix):
		return B64ENCTransform, nil
	case strings.HasPrefix(filePath, FUTLPrefix):
		return FUTLTransorm, nil
	case strings.HasPrefix(filePath, reader.FilePrefix):
		return reader.ReadFile, nil
	case strings.HasPrefix(filePath, reader.DirPrefix):
		return reader.ReadDir, nil
	}
	if !strings.Contains(filePath, "://") {
		return reader.ReadFile, nil
	}
	return nil, errors.Errorf("error while parsing filepath %s with file utils plugin", filePath)
}
