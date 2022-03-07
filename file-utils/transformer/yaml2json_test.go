package transformer

import "testing"

func TestYaml2JsonTransform(t *testing.T) {
	var testCase = NewCmdTestCase("yaml2json", "yaml2json test",
		"futl://../../tests/yaml_json/yaml2json/input/values.yaml",
		"../../tests/yaml_json/yaml2json/output/expected_json.yaml", "", false,
	)
	ExecuteTests(t, *testCase)
}
