package transformer

import (
	"errors"
	reader "github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"strings"
)

var schemesMap = map[string]bool{
	B64EncPrefix: true,
	B64DecPrefix: true,
	FUTLPrefix:   true,
}

type Factory func(inputValue reader.InputValue) (reader.InputValue, error)

type URI struct {
	TransformSchemes []string
	InputURL         string
}

func ParseURI(uri string) (URI, error) {
	result := URI{}

	uriFragments := strings.Split(uri, "://")
	if len(uriFragments) != 2 {
		return result, errors.New("invalid format URL")
	}

	schemes := strings.Split(uriFragments[0], "+")
	for i := len(schemes) - 1; i >= 0; i-- {
		if i == len(schemes)-1 {
			if !schemesMap[schemes[i]] && !reader.InputSchemesMap[schemes[i]] {
				return result, errors.New("invalid combination of protocol schemes")
			}
		} else {
			if !schemesMap[schemes[i]] || reader.InputSchemesMap[schemes[i]] {
				return result, errors.New("invalid combination of protocol schemes")
			}
		}

		if schemesMap[schemes[i]] {
			result.TransformSchemes = append(result.TransformSchemes, schemes[i])
		} else if reader.InputSchemesMap[schemes[i]] {
			result.InputURL = schemes[i] + "://" + uriFragments[1]
		}
	}
	if result.InputURL == "" {
		result.InputURL = reader.FilePrefix + "://" + uriFragments[1]
	}
	return result, nil
}

func DetermineTransformer(scheme string) (Factory, error) {
	switch {
	case scheme == B64EncPrefix:
		return B64ENCTransform, nil
	case scheme == B64DecPrefix:
		return B64DECTransform, nil
	case scheme == FUTLPrefix:
		return FUTLTransform, nil
	}
	return nil, errors.New("transform scheme not found")
}
