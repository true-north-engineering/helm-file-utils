package transformer

import (
	"testing"
)

func TestBase64EncDec(t *testing.T) {
	var testCase = NewCmdTestCase("base64encdec", "base64 encoder decoder",
		"futl://../../tests/base64/encdec/input/values.yaml",
		"../../tests/base64/encdec/output/encdec-test.txt", "", true,
	)
	ExecuteTests(t, *testCase)
}
