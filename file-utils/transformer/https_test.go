package transformer

import "testing"

func TestHttps(t *testing.T) {
	var testCase = NewCmdTestCase("https", "https test",
		"futl://../../tests/https/input/values.yaml",
		"../../tests/https/output/https-test.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
