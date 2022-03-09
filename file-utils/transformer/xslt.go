package transformer

import (
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
)

const (
	XsltPrefix = "xslt"
)

func XsltTransform(inputValue reader.InputValue) (reader.InputValue, error) {
	result := reader.InputValue{Kind: inputValue.Kind, Value: make(map[string][]byte)}

	return result, nil
}
