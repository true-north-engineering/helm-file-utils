package transformer

import (
	"encoding/json"
	"fmt"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"gopkg.in/yaml.v3"
)

const (
	Json2YamlPrefix = "json2yaml"
)

// Json2YamlTransform Transformer that transforms given .json file into .yaml file.
func Json2YamlTransform(inputValue reader.InputValue) (reader.InputValue, error) {

	result := reader.InputValue{Kind: inputValue.Kind, Value: make(map[string][]byte)}

	if inputValue.Kind == reader.InputKindFile {
		yamlfile, err := JSONToYAMLFull(inputValue.Value[reader.InputKindFile])

		if err != nil {
			return reader.InputValue{}, err
		}
		result.Value[reader.InputKindFile] = yamlfile
	} else if inputValue.Kind == reader.InputKindDir {
		inputFiles := inputValue.Value
		for fileName, fileValue := range inputFiles {
			yamlfile, err := JSONToYAMLFull(fileValue)
			if err != nil {
				fmt.Println(err)
				continue
			}
			result.Value[fileName] = yamlfile
		}
	}

	return result, nil
}

func ConvertJson2Yaml(jsonData map[string]interface{}) ([]byte, error) {

	output, err := yaml.Marshal(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling YAML: %s", err.Error())
	}

	return output, nil
}

// UnmarshalJSONFile will return a JSON file in it's raw Golang form.
func UnmarshalJSONFile(fileData []byte) (map[string]interface{}, error) {

	var jsonData map[string]interface{}

	err := json.Unmarshal(fileData, &jsonData)

	return jsonData, err
}

// JSONToYAMLFull is a wrapper function around the other underlying functions
// for ease of use. Simply, a file is specified and the conversion is handled internally.
func JSONToYAMLFull(fileData []byte) ([]byte, error) {

	jsonData, err := UnmarshalJSONFile(fileData)
	if err != nil {
		return nil, err
	}

	yamlOutput, err := ConvertJson2Yaml(jsonData)
	if err != nil {
		return nil, err
	}

	return yamlOutput, err
}
