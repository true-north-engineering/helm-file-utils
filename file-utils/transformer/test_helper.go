package transformer

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"io/ioutil"
	"testing"
)

// cmdTestCase describes a test case that works with releases.
type cmdTestCase struct {
	name        string
	description string
	input       string
	golden      string
	protocol    string
	wantError   bool
}

func NewCmdTestCase(name string, description string, input string, golden string, protocol string, wantError bool) *cmdTestCase {
	return &cmdTestCase{name: name, description: description, input: input, golden: golden, protocol: protocol, wantError: wantError}
}

func ExecuteTests(t *testing.T, testCase cmdTestCase) {
	tests := []cmdTestCase{testCase}
	runTestCmd(t, tests)
}

func runTestCmd(t *testing.T, tests []cmdTestCase) {

	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			t.Logf("running cmd : %s", tt.input)
			t.Logf("running test : %s", tt.name)

			out, err := executeActionCommand(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("expected error, got '%v'", err)
			} else if err != nil {
				t.Errorf("got error '%v'", err)
				return
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

func executeActionCommand(input string) ([]byte, error) {
	result, err := ExecuteTransformations(input)
	return result.Value[reader.InputKindFile], err
}

// AssertGoldenString asserts that the given string matches the contents of the given file.
func assertGoldenString(actual []byte, filename string) error {
	if err := compare(actual, filename); err != nil {
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
