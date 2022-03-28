package tests

import (
	"testing"
)

func TestExecutor(t *testing.T) {
	var testCase = NewCmdTestCase("executor", "base64 decoder",
		"futl://../../tests/futl/input/values_futl.yaml",
		"../../tests/futl/output/futl_output.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
