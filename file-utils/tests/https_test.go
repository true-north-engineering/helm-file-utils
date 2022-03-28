package tests

import (
	"testing"
)

func TestHttps(t *testing.T) {
	var testCase = NewCmdTestCase("https", "https test",
		"futl://../../tests/https/input/values_https.yaml",
		"../../tests/https/output/https_output.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
