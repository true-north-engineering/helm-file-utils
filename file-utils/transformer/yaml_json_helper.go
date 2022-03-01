package transformer

import (
	"fmt"
	"path/filepath"
)

// Allows YAML to be encoded into JSON format easily.
func cleanYaml(yamlData map[interface{}]interface{}) map[string]interface{} {

	cleanYamlMapping := make(map[string]interface{})

	for key, value := range yamlData {

		// Type assertion on the key within the yaml, key will be type of interface{}
		// so it must be asserted to ensure it is string.
		assertedKey := key.(string)
		cleanYamlMapping[assertedKey] = value

		assertedMapVal, isInterfaceMapType := value.(map[interface{}]interface{})
		assertedSliceVal, isInterfaceSliceType := value.([]interface{})

		// If the value is also another map, then you need to retreive that value, adding it into the outer map.
		if isInterfaceMapType {
			cleanInnerMap := cleanYaml(assertedMapVal)
			cleanYamlMapping[assertedKey] = cleanInnerMap
		}

		// If the item is a interface slice, we need to check whether it contains a map[interface{}]interface{} type, if so we can convert it.
		if isInterfaceSliceType {
			for _, item := range assertedSliceVal {

				itemAsserted, isInnerMap := item.(map[interface{}]interface{})

				if isInnerMap {
					cleanInnerMap := cleanYaml(itemAsserted)
					cleanYamlMapping[assertedKey] = cleanInnerMap
				}

			}

		}
	}

	return cleanYamlMapping
}

// IsYAMLFile checks whether a specified file is YAML.
func IsYAMLFile(filePath string) (bool, error) {

	fileType := filepath.Ext(filePath)

	if fileType == ".yml" || fileType == ".yaml" {
		return true, nil
	}
	fmt.Println("nije yaml")
	fmt.Println(fileType)
	return false, fmt.Errorf("given file is not yaml file")
}

// IsJSONFile checks whether a specified file is JSON.
func IsJSONFile(filePath string) (bool, error) {

	fileType := filepath.Ext(filePath)
	if fileType == ".json" {
		return true, nil
	}

	return false, fmt.Errorf("given file is not json file")
}
