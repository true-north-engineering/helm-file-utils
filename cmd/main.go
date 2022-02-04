package main

import (
	"fmt"
	"log"
	"os"

	"github.com/true-north-engineering/helm-file-utils/cmd/reader"
	"github.com/true-north-engineering/helm-file-utils/cmd/transformer"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal("error while running file utils plugin, filepath argument is not correctly specified.")
	}
	result, err := transformer.ExecuteTransformations(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	if result.Kind != reader.InputKindFile {
		log.Fatal("error dir scheme available only inside futl parsed yaml file")
	}
	fmt.Println(string(result.Value[reader.InputKindFile]))
}
