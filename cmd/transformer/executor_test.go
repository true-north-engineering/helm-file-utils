package transformer

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/cmd/reader"
	"io/ioutil"
	"testing"
)

func TestFUTLDownloaderPlugin(t *testing.T) {
	tests := []cmdTestCase{
		{
			name:   "chart with template with external dir",
			input:  "futl://../../tests/futl/input/values.yaml",
			golden: "../../tests/futl/output/futl-test.txt",
		},
	}
	runTestCmd(t, tests)
}

func runTestCmd(t *testing.T, tests []cmdTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("running cmd : %s", tt.input)
			out, err := executeActionCommandC(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("expected error, got '%v'", err)
			}
			if tt.golden != "" {
				err := assertGoldenString(out, tt.golden)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func executeActionCommandC(input string) (string, error) {
	result, err := ExecuteTransformations(input)
	return string(result.Value[reader.InputKindFile]), err
}

// cmdTestCase describes a test case that works with releases.
type cmdTestCase struct {
	name      string
	input     string
	golden    string
	protocol  string
	wantError bool
}

// assertGoldenString asserts that the given string matches the contents of the given file.
func assertGoldenString(actual, filename string) error {
	if err := compare([]byte(actual), filename); err != nil {
		return err
	}
	return nil
}

func compare(actual []byte, filename string) error {
	actual = normalize(actual)
	expected, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, "unable to read testdata %s", filename)
	}
	expected = normalize(expected)
	if !bytes.Equal(expected, actual) {
		return errors.Errorf("does not match golden file %s\n\nWANT:\n'%s'\n\nGOT:\n'%s'", filename, expected, actual)
	}
	return nil
}

func normalize(in []byte) []byte {
	return bytes.Replace(in, []byte("\r\n"), []byte("\n"), -1)
}
