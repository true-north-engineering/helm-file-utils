package transformer

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
	reader "github.com/true-north-engineering/helm-file-utils/file-utils/reader"
)

var schemesMap = map[string]bool{
	B64EncPrefix:    true,
	B64DecPrefix:    true,
	Json2YamlPrefix: true,
	Yaml2JsonPrefix: true,
	FUTLPrefix:      true,
	XsltPrefix:      true,
	CustomPrefix:    true,
}

type Factory func(inputValue reader.InputValue, args string) (reader.InputValue, error)

type URI struct {
	TransformSchemes    []string
	TransformSchemesArg []string
	InputURL            string
}

func ParseURI(uri string) (URI, error) {
	result := URI{}

	uriFragments := strings.Split(uri, "://")
	if len(uriFragments) > 2 || len(uriFragments) < 1 {
		return result, errors.New("invalid format URL")
	}

	if len(uriFragments) == 1 {
		result.InputURL = reader.FilePrefix + "://" + uriFragments[0]
		return result, nil
	}

	fragmentWithArgsRegexp := regexp.MustCompile(`(\w+)\(([^\)]*)\)`)
	schemes := strings.Split(uriFragments[0], "+")
	for i := len(schemes) - 1; i >= 0; i-- {

		fragmentWithArgsMatch := fragmentWithArgsRegexp.FindStringSubmatch(schemes[i])
		scheme := schemes[i]
		schemeArgs := ""
		if len(fragmentWithArgsMatch) == 3 {
			scheme = fragmentWithArgsMatch[1]
			schemeArgs = fragmentWithArgsMatch[2]
		}

		if i == len(schemes)-1 {
			if !schemesMap[scheme] && !reader.InputSchemesMap[scheme] {
				return result, errors.New("invalid combination of protocol schemes")
			}
		} else {
			if !schemesMap[scheme] || reader.InputSchemesMap[scheme] {
				return result, errors.New("invalid combination of protocol schemes")
			}
		}

		if schemesMap[scheme] {
			result.TransformSchemes = append(result.TransformSchemes, scheme)
			result.TransformSchemesArg = append(result.TransformSchemesArg, schemeArgs)
		} else if reader.InputSchemesMap[scheme] {
			result.InputURL = scheme + "://" + uriFragments[1]
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
	case scheme == Json2YamlPrefix:
		return Json2YamlTransform, nil
	case scheme == Yaml2JsonPrefix:
		return Yaml2JsonTransform, nil
	case scheme == FUTLPrefix:
		return FUTLTransform, nil
	case scheme == XsltPrefix:
		return XsltTransform, nil
	case scheme == CustomPrefix:
		return CustomTransform, nil
	}

	return nil, errors.New("transform scheme not found")
}
