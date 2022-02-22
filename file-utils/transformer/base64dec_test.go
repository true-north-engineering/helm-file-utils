package transformer

import (
	"testing"
)

func TestBase64Dec(t *testing.T) {
	var testCase = NewCmdTestCase("base64dec", "base64 decoder",
		"futl://../../tests/base64/dec/input/values.yaml",
		"../../tests/base64/dec/output/dec-test.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
