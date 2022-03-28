package tests

import (
	"testing"
)

func TestBase64Enc(t *testing.T) {
	var testCase = NewCmdTestCase("base64enc", "base64 encoder",
		"futl://../../tests/base64/enc/input/values_base64enc.yaml",
		"../../tests/base64/enc/output/base64enc_output.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
