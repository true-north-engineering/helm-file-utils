package tests

import (
	"testing"
)

func TestMultiple(t *testing.T) {
	var testCase = NewCmdTestCase("multiple test", "testing multiple transformations chained",
		"futl://../../tests/multiple/input/values_multiple.yaml",
		"../../tests/multiple/output/multiple_output.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
