package transformer

import "testing"

func TestGitHttps(t *testing.T) {
	var testCaseFile = NewCmdTestCase("git_https file", "git_https file test",
		"futl://../../tests/git_https/input/values.yaml",
		"../../tests/git_https/output/git_https_test.txt", "", false,
	)
	ExecuteTests(t, *testCaseFile)
	var testCaseDir = NewCmdTestCase("git_https dir", "git_https dir test",
		"futl://../../tests/git_https/input/values_dir.yaml",
		"../../tests/git_https/output/git_https_test_dir.txt", "", false,
	)
	ExecuteTests(t, *testCaseDir)
}
