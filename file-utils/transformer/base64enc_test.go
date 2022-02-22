package transformer

import (
	"testing"
)

func TestBase64Enc(t *testing.T) {
	var testCase = NewCmdTestCase("base64enc", "base64 encoder",
		"futl://../../tests/base64/enc/input/values.yaml",
		"../../tests/base64/enc/output/enc-test.txt", "", false,
	)
	ExecuteTests(t, *testCase)
}
