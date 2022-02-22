package main

import (
	"encoding/binary"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"github.com/true-north-engineering/helm-file-utils/file-utils/transformer"
	"log"
	"os"
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
	binary.Write(os.Stdout, binary.LittleEndian, result.Value[reader.InputKindFile])
}
