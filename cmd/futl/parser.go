package futl

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/cmd/base64"
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
		var value string
		var err error
		switch {
		case strings.HasPrefix(x, "base64enc://"):
			value, err = base64.EncodeFile(strings.TrimPrefix(x, "base64enc://"))

		case strings.HasPrefix(x, "file://"):
			var file []byte
			file, err = ioutil.ReadFile(strings.TrimPrefix(x, "file://"))
			value = string(file)
		default:
			value = x
		}
		if err != nil {
			log.Fatal(err)
		}
		return value
	}
	return i
}
