package tests

import (
	"testing"
)

func TestJson2YamlTransform(t *testing.T) {
	var testCase = NewCmdTestCase("json2yaml", "json2yaml test",
		"futl://../../tests/yaml_json/json2yaml/input/values_json2yaml.yaml",
		"../../tests/yaml_json/json2yaml/output/json2yaml_output.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
