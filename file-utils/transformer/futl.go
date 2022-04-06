package transformer

import (
	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"gopkg.in/yaml.v3"
)

const (
	FUTLPrefix = "futl"
	FUTLTag    = "!futl"
)

func FUTLTransform(inputValue reader.InputValue) (reader.InputValue, error) {
	if inputValue.Kind != reader.InputKindFile {
		return reader.InputValue{}, errors.New("wrong input type into futl transformer")
	}
	file := inputValue.Value[reader.InputKindFile]

	var parsedYaml interface{}
	err := yaml.Unmarshal(file, &FUTLTagProcessor{&parsedYaml})
	if err != nil {
		return reader.InputValue{}, err
	}
	outputYaml, err := yaml.Marshal(parsedYaml)
	if err != nil {
		return reader.InputValue{}, err
	}
	result := reader.InputValue{
		Kind: reader.InputKindFile,
		Value: map[string][]byte{
			reader.InputKindFile: outputYaml,
		},
	}
	return result, nil
}

type FUTLTagProcessor struct {
	source interface{}
}

func (i *FUTLTagProcessor) UnmarshalYAML(node *yaml.Node) error {
	resolved, err := resolveFUTLTags(node)
	if err != nil {
		return err
	}
	return resolved.Decode(i.source)
}

func resolveFUTLTags(node *yaml.Node) (*yaml.Node, error) {
	if node.Tag == FUTLTag {
		fileURL := node.Value

		transformedValue, err := ExecuteTransformations(fileURL)
		if err != nil {
			return nil, err
		}
		if transformedValue.Kind == reader.InputKindFile {
			node.Value = string(transformedValue.Value[reader.InputKindFile])
		} else if transformedValue.Kind == reader.InputKindDir {
			node = &yaml.Node{Kind: yaml.MappingNode}

			for fileName, fileValue := range transformedValue.Value {
				fileNameNode := &yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: fileName,
				}
				fileValueNode := &yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: string(fileValue),
				}
				node.Content = append(node.Content, fileNameNode, fileValueNode)
			}
		}
	}
	if node.Kind == yaml.SequenceNode || node.Kind == yaml.MappingNode {
		var err error
		for i := range node.Content {
			node.Content[i], err = resolveFUTLTags(node.Content[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return node, nil
}
