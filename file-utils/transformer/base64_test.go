package transformer

import (
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"testing"
)

func TestBase64(t *testing.T) {
	tests := []cmdTestCase{
		{
			name:   "chart with template with external dir",
			input:  "futl://../../tests/futl/input/values.yaml",
			golden: "../../tests/futl/output/futl-test.txt",
		},
	}
	runTestBase64(t, tests)
}

func runTestBase64(t *testing.T, tests []cmdTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("running cmd : %s", tt.input)
			out, err := executeActionCommandBase64(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("expected error, got '%v'", err)
			}
			if tt.golden != "" {
				err := assertGoldenStringBase64(out, tt.golden)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func executeActionCommandBase64(input string) ([]byte, error) {
	result, err := ExecuteTransformations(input)
	return result.Value[reader.InputKindFile], err
}

func assertGoldenStringBase64(actual []byte, filename string) error {
	if err := compare(actual, filename); err != nil {
		return err
	}
	return nil
}
