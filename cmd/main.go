package main

import (
	"fmt"
	"github.com/true-north-engineering/helm-file-utils/cmd/base64"
	"log"
	"os"
	"strings"
)

var version = "Version is not provided"

const (
	base64enc = "base64enc://"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatalf("error while running file utils plugin, filepath argument is not correctly specified.")
		os.Exit(1)
	}
	filePath := os.Args[4]
	switch {
	case strings.HasPrefix(filePath, base64enc):
		encodedFile, err := base64.EncodeFile(strings.TrimPrefix(filePath, base64enc))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(encodedFile)
	default:
		log.Fatalf("error while parsing filepath %s with file utils plugin", filePath)
	}
}
