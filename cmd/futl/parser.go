package futl

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func ParseFile(filePath string) (string, error) {
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
		if strings.HasPrefix(x, "futl+base64enc://") {
			return "AAA"
		}
	}
	return i
}
