package transformer

import (
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
)

func ExecuteTransformations(fileURIString string) (reader.InputValue, error) {

	fileURI, err := ParseURI(fileURIString)
	if err != nil {
		return reader.InputValue{}, err
	}

	readByProtocol, err := reader.DetermineReader(fileURI.InputURL)

	if err != nil {
		return reader.InputValue{}, err
	}

	inputValue, err := readByProtocol(fileURI.InputURL)

	if err != nil {
		return reader.InputValue{}, err
	}

	for _, transformScheme := range fileURI.TransformSchemes {

		transformByProtocol, err := DetermineTransformer(transformScheme)
		if err != nil {
			return reader.InputValue{}, err
		}

		parsedValue, err := transformByProtocol(inputValue)
		if err != nil {
			return reader.InputValue{}, err
		}

		inputValue = parsedValue
	}
	return inputValue, nil
}
