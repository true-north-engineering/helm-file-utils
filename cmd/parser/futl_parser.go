package parser

import (
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	FUTLPrefix = "futl://"
)

func ParseFile(filePath string) (string, error) {
	filePath = strings.TrimPrefix(filePath, FUTLPrefix)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.Errorf("file %s does not exist", filePath)
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.Errorf("file %s cannot be read", filePath)
	}

	var yamlBody interface{}
	err = yaml.Unmarshal(file, &yamlBody)

	yamlBody = expandFile(yamlBody)

	outputYaml, err := yaml.Marshal(yamlBody)
	return string(outputYaml), nil
}

func expandFile(i interface{}) interface{} {
	switch x := i.(type) {
	case map[string]interface{}:
		for k, v := range x {
			x[k] = expandFile(v)
		}
	case []interface{}:
		for i, v := range x {
			x[i] = expandFile(v)
		}
	case string:
		var value string

		parserFunc, _ := DetermineParser(x)
		if parserFunc == nil {
			return x
		}

		value, err := parserFunc(x)
		if err != nil {
			log.Fatal(err)
		}
		return value
	}
	return i
}
