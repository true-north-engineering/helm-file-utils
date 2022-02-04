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
	for i := len(schemes) - 1; i >= 0; i-- {
		if i == len(schemes)-1 {
			if !TransformerSchemes[schemes[i]] && !reader.InputSchemes[schemes[i]] {
				return result, errors.New("Invalid combination of protocol schemes")
			}
		} else {
			if !TransformerSchemes[schemes[i]] || reader.InputSchemes[schemes[i]] {
				return result, errors.New("Invalid combination of protocol schemes")
			}
		}

		if TransformerSchemes[schemes[i]] {
			result.TransformSchemes = append(result.TransformSchemes, schemes[i])
		} else if reader.InputSchemes[schemes[i]] {
			result.InputURL = schemes[i] + "://" + uriFragments[1]
		}
	}
	if result.InputURL == "" {
		result.InputURL = "file://" + uriFragments[1]
	}
	return result, nil
}
