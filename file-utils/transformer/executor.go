package transformer

import (
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"log"
	"os"
)

func ExecuteTransformations(fileURIString string) (reader.InputValue, error) {
	fileURI, err := ParseURI(fileURIString)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	readByProtocol, err := reader.DetermineReader(fileURI.InputURL)
	inputValue, err := readByProtocol(fileURI.InputURL)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
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
