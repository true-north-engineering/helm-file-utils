package parser

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	FUTLPrefix = "futl://"
	FUTLTag    = "!futl"
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

	var parsedYaml interface{}
	err = yaml.Unmarshal(file, &FutlTagProcessor{&parsedYaml})
	if err != nil {
		return "", err
	}
	outputYaml, err := yaml.Marshal(parsedYaml)
	if err != nil {
		return "", err
	}
	return string(outputYaml), nil
}

type FutlTagProcessor struct {
	source interface{}
}

func (i *FutlTagProcessor) UnmarshalYAML(node *yaml.Node) error {
	resolved, err := resolveFutlTags(node)
	if err != nil {
		return err
	}
	return resolved.Decode(i.source)
}

func resolveFutlTags(node *yaml.Node) (*yaml.Node, error) {
	if node.Tag == FUTLTag {
		fileURL := node.Value
		parserFunc, err := DetermineParser(fileURL)
		if err != nil {
			return nil, err
		}
		value, err := parserFunc(fileURL)
		if err != nil {
			return nil, err
		}
		node.Value = value
	}
	if node.Kind == yaml.SequenceNode || node.Kind == yaml.MappingNode {
		var err error
		for i := range node.Content {
			node.Content[i], err = resolveFutlTags(node.Content[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return node, nil
}
