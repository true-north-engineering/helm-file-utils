package main

import (
	"fmt"
	"github.com/true-north-engineering/helm-file-utils/cmd/parser"
	"os"
)

//var version = "Version is not provided"

func main() {
	if len(os.Args) < 5 {
		fmt.Println("error while running file utils plugin, filepath argument is not correctly specified.")
		os.Exit(1)
	}
	filePath := os.Args[4]
	transformByProtocol, err := parser.DetermineParser(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	parsedValue, err := transformByProtocol(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(parsedValue)
}
