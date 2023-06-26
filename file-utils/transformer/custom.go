package transformer

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
)

const (
	CustomPrefix = "custom"
)

func CustomTransform(inputValue reader.InputValue, args string) (reader.InputValue, error) {
	commandLine := strings.Split(args, " ")
	result := reader.InputValue{Kind: inputValue.Kind, Value: make(map[string][]byte)}

	if inputValue.Kind == reader.InputKindFile {
		cmd := exec.Command(commandLine[0], commandLine[1:]...)
		inputFile := inputValue.Value[reader.InputKindFile]
		var cmdOutput bytes.Buffer
		var cmdStdErr bytes.Buffer

		cmd.Stdin = strings.NewReader(string(inputFile))
		cmd.Stdout = &cmdOutput
		cmd.Stderr = &cmdStdErr
		cmdErr := cmd.Run()
		if cmdErr != nil {
			return reader.InputValue{}, errors.WithMessage(cmdErr, cmdStdErr.String())
		}
		result.Value[reader.InputKindFile] = cmdOutput.Bytes()
	} else if inputValue.Kind == reader.InputKindDir {
		inputFiles := inputValue.Value

		for fileName, fileValue := range inputFiles {
			cmd := exec.Command(commandLine[0], commandLine[1:]...)
			var cmdOutput bytes.Buffer
			var cmdStdErr bytes.Buffer

			cmd.Stdin = strings.NewReader(string(fileValue))
			cmd.Stdout = &cmdOutput
			cmd.Stderr = &cmdStdErr
			cmdErr := cmd.Run()
			if cmdErr != nil {
				return reader.InputValue{}, errors.WithMessage(cmdErr, cmdStdErr.String())
			}
			result.Value[fileName] = cmdOutput.Bytes()
		}
	}

	return result, nil
}
