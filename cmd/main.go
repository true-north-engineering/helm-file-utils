package main

import (
	"fmt"
	"github.com/true-north-engineering/helm-file-utils/cmd/base64"
	"github.com/true-north-engineering/helm-file-utils/cmd/futl"
	"log"
	"os"
	"strings"
)

var version = "Version is not provided"

const (
	base64encPrefix = "base64enc://"
	futlPrefix = "futl://"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatalf("error while running file utils plugin, filepath argument is not correctly specified.")
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
			log.Fatal(err)
		}
		fmt.Println(encodedFile)
	default:
		log.Fatalf("error while parsing filepath %s with file utils plugin", filePath)
	}
}
