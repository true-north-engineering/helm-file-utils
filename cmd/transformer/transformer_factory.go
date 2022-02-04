package transformer

import (
	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/cmd/reader"
)

var TransformerSchemes = map[string]bool{
	B64EncPrefix: true,
	FUTLPrefix:   true,
}

type Factory func(inputValue reader.InputValue) (reader.InputValue, error)

func DetermineTransformer(scheme string) (Factory, error) {
	switch {
	case scheme == B64EncPrefix:
		return B64ENCTransform, nil
	case scheme == FUTLPrefix:
		return FUTLTransorm, nil
	}
	return nil, errors.Errorf("Transform scheme not found")
}
