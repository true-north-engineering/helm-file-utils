package main

import (
	"github.com/spf13/cobra"
	"github.com/true-north-engineering/helm-file-utils/cmd/base64enc"
	"log"
)

var version = "Version is not provided"

var cmd = &cobra.Command{
	Use:   "",
	Short: "Helm file utils plugin",
	RunE: func(cmd *cobra.Command, aargs []string) error {
		return cmd.Help()
	},
}

func main() {
	cmd.AddCommand(base64enc.Command())
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
