package transformer

import (
	"encoding/json"
	"fmt"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"gopkg.in/yaml.v3"
)

const (
	Yaml2JsonPrefix = "yaml2json"
)

// Yaml2JsonTransform Transformer that transforms given .yaml or .yml file into .json file.
func Yaml2JsonTransform(inputValue reader.InputValue) (reader.InputValue, error) {

	result := reader.InputValue{Kind: inputValue.Kind, Value: make(map[string][]byte)}

	if inputValue.Kind == reader.InputKindFile {
		jsonfile, err := YAMLToJSONFull(inputValue.Value[reader.InputKindFile])

		if err != nil {
			return reader.InputValue{}, err
		}
		result.Value[reader.InputKindFile] = jsonfile
	} else if inputValue.Kind == reader.InputKindDir {
		inputFiles := inputValue.Value
		for fileName, fileValue := range inputFiles {
			jsonfile, err := YAMLToJSONFull(fileValue)
			if err != nil {
				fmt.Println(err)
				continue
			}
			result.Value[fileName] = jsonfile
		}
	}

	return result, nil
}

// ConvertYAMLToJSON will convert raw YAML into a JSON encoded byte array, this is ready to be written to a file.
func ConvertYAMLToJSON(yamlData map[interface{}]interface{}) ([]byte, error) {

	cleanedYaml := cleanYaml(yamlData)

	output, err := json.MarshalIndent(cleanedYaml, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error converting yaml to json: %s", err.Error())
	}

	return output, nil
}

// UnmarshalYAMLFile will return a YAML configuration file in it's raw Golang form.
func UnmarshalYAMLFile(fileData []byte) (map[interface{}]interface{}, error) {

	yamlData := make(map[interface{}]interface{})

	err := yaml.Unmarshal(fileData, &yamlData)

	return yamlData, err

}

// YAMLToJSONFull is a wrapper function around the other underlying functions
// for ease of use. Simply, a file is specified and the conversion is handled internally.
func YAMLToJSONFull(fileData []byte) ([]byte, error) {

	yamlData, err := UnmarshalYAMLFile(fileData)
	if err != nil {
		return nil, err
	}

	jsonOutput, err := ConvertYAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}

	return jsonOutput, err
}
