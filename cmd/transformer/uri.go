package transformer

import (
	"errors"
	"strings"

	"github.com/true-north-engineering/helm-file-utils/cmd/reader"
)

type URI struct {
	TransformSchemes []string
	InputURL         string
}

func ParseURI(uri string) (URI, error) {
	result := URI{}

	uriFragments := strings.Split(uri, "://")
	if len(uriFragments) != 2 {
		return result, errors.New("Invalid format URL")
	}

	schemes := strings.Split(uriFragments[0], "+")
	for i := 0; i < len(schemes); i++ {
		if TransformerSchemes[schemes[i]] {
			result.TransformSchemes = append(result.TransformSchemes, schemes[i])
		} else {
			if i != len(schemes)-1 {
				return result, errors.New("Invalid transformer scheme")
			} else if !reader.InputSchemes[schemes[i]] {
				return result, errors.New("Invalid transformer scheme")
			} else {
				result.InputURL = schemes[i] + "://" + uriFragments[1]
			}
		}
	}
	if result.InputURL == "" {
		result.InputURL = "file://" + uriFragments[1]
	}
	return result, nil
}
