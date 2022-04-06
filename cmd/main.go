package main

import (
	"encoding/binary"
	"fmt"
	"github.com/true-north-engineering/helm-file-utils/file-utils/reader"
	"github.com/true-north-engineering/helm-file-utils/file-utils/transformer"
	"log"
	"os"
)

var version = "develop"

func main() {

	if ok := versionCmd(os.Args); ok {
		return
	}

	if len(os.Args) < 5 {
		log.Fatal("error while running file utils plugin, filepath argument is not correctly specified.")
	}
	result, err := transformer.ExecuteTransformations(os.Args[4])

	if err != nil {
		os.Stderr.Write([]byte("Krafna"))
		os.Exit(1)
	}
	if result.Kind != reader.InputKindFile {
		log.Fatal("error dir scheme available only inside futl parsed yaml file")
	}
	err = binary.Write(os.Stdout, binary.LittleEndian, result.Value[reader.InputKindFile])

	if err != nil {
		log.Fatal("failed to write transformed file")
	}
}

// versionCmd print version if command `version` is passed as an argument and return true otherwise return false.
func versionCmd(args []string) bool {
	if len(args) >= 2 && args[1] == "version" {
		fmt.Printf("Version: %s\n", version)
		return true
	}
	return false
}
