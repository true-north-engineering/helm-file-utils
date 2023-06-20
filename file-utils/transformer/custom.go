package transformer

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
)

const (
	CustomPrefix = "custom"
)

func CustomTransform(inputValue reader.InputValue, args string) (reader.InputValue, error) {
	commandLine := strings.Split(args, " ")
	cmd := exec.Command(commandLine[0], commandLine[1:]...)
	result := reader.InputValue{Kind: inputValue.Kind, Value: make(map[string][]byte)}

	if inputValue.Kind == reader.InputKindFile {
		inputFile := inputValue.Value[reader.InputKindFile]
		var cmdOutput bytes.Buffer
		var cmdStdErr bytes.Buffer

		cmd.Stdin = strings.NewReader(string(inputFile))
		cmd.Stdout = &cmdOutput
		cmd.Stderr = &cmdStdErr
		cmdErr := cmd.Run()
		if cmdErr != nil {
			return reader.InputValue{}, cmdErr
		}
		result.Value[reader.InputKindFile] = cmdOutput.Bytes()
	} else if inputValue.Kind == reader.InputKindDir {
		inputFiles := inputValue.Value
		var cmdOutput bytes.Buffer
		var cmdStdErr bytes.Buffer

		for fileName, fileValue := range inputFiles {
			cmd.Stdin = strings.NewReader(string(fileValue))
			cmd.Stdout = &cmdOutput
			cmd.Stderr = &cmdStdErr
			cmdErr := cmd.Run()
			if cmdErr != nil {
				return reader.InputValue{}, cmdErr
			}
			result.Value[fileName] = cmdOutput.Bytes()
		}
	}

	return result, nil
}