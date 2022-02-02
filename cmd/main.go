package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/true-north-engineering/helm-file-utils/cmd/base64"
	"github.com/true-north-engineering/helm-file-utils/cmd/futl"
)

const (
	base64encPrefix = "base64enc://"
	filePrefix      = "file://"
	futlPrefix      = "futl://"
)

var version = "Version is not provided"

func main() {
	if len(os.Args) < 5 {
		fmt.Println("error while running file utils plugin, filepath argument is not correctly specified.")
		os.Exit(1)
	}
	filePath := os.Args[4]
	switch {
	case strings.HasPrefix(filePath, base64encPrefix):
		encodedFile, err := base64.EncodeFile(strings.TrimPrefix(filePath, base64encPrefix))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encodedFile)
	case strings.HasPrefix(filePath, futlPrefix):
		encodedFile, err := futl.ParseFile(strings.TrimPrefix(filePath, futlPrefix))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(encodedFile)
	case strings.HasPrefix(filePath, filePrefix):
		file, err := ioutil.ReadFile(strings.TrimPrefix(filePath, filePrefix))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(file))
	default:
		fmt.Printf("error while parsing filepath %s with file utils plugin", filePath)
	}
}
