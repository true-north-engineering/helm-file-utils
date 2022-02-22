package transformer

import (
	"testing"
)

func TestExecutor(t *testing.T) {
	var testCase = NewCmdTestCase("executor", "base64 decoder",
		"futl://../../tests/futl/input/values.yaml",
		"../../tests/futl/output/futl-test.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
