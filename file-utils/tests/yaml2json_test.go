package tests

import (
	"testing"
)

func TestYaml2JsonTransform(t *testing.T) {
	var testCase = NewCmdTestCase("yaml2json", "yaml2json test",
		"futl://../../tests/yaml_json/yaml2json/input/values_yaml2json.yaml",
		"../../tests/yaml_json/yaml2json/output/yaml2json_output.yaml", "", false,
	)
	ExecuteTests(t, *testCase)
}
