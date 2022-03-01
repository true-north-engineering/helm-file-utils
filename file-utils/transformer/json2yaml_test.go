package transformer

import "testing"

func TestJson2YamlTransform(t *testing.T) {
	var testCase = NewCmdTestCase("json2yaml", "json2yaml test",
		"futl://../../tests/yaml_json/json2yaml/input/values.yaml",
		"../../tests/yaml_json/json2yaml/output/expected_yaml.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
