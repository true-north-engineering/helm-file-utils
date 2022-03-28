package tests

import (
	"testing"
)

func TestBase64Dec(t *testing.T) {
	var testCase = NewCmdTestCase("base64dec", "base64 decoder",
		"futl://../../tests/base64/dec/input/values_base64dec.yaml",
		"../../tests/base64/dec/output/base64dec_output.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
