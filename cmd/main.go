package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/true-north-engineering/helm-file-utils/cmd/transformer"
)

//var version = "Version is not provided"

func main() {
	if len(os.Args) < 5 {
		fmt.Println("error while running file utils plugin, filepath argument is not correctly specified.")
		os.Exit(1)
	}
	filePath, err := url.Parse(os.Args[4])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	transformByProtocol, err := transformer.DetermineTransformer(filePath.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	parsedValue, err := transformByProtocol(filePath.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch parsedValue.(type) {
	case []byte:
		fmt.Println(string(parsedValue.([]byte)))
	default:
		log.Fatal("Resulting value is not a single file")
		os.Exit(1)
	}
}
