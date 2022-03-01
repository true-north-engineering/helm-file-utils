package transformer

import "testing"

func TestYaml2JsonJson2YamlTransform(t *testing.T) {
	var testCase = NewCmdTestCase("yaml2json_json2yaml", "yaml2jsonjson2yaml test",
		"futl://../../tests/yaml_json/yaml2json_json2yaml/input/values.yaml",
		"../../tests/yaml_json/yaml2json_json2yaml/output/expected_json.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}