package tests

import (
	"testing"
)

func TestGitHttps(t *testing.T) {
	var testCaseFile = NewCmdTestCase("git_https", "git_https test",
		"futl://../../tests/git_https/input/values_git_https.yaml",
		"../../tests/git_https/output/git_https_output.txt", "", false,
	)
	ExecuteTests(t, *testCaseFile)
}
