package base64enc

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var name = "helm-file-utils"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "base64enc",
		Short: "Run base64 encoding",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.GetViper().BindPFlags(cmd.Flags()); err != nil {
				return errors.Wrap(err, "failed to bind flags")
			}
			return nil
		},
		RunE: cmdF(),
	}

	return cmd
}

func cmdF() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {

		fmt.Print(args)
		filePath := strings.TrimPrefix(args[3], "base64://")

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return errors.Errorf("file %s does not exist", filePath)
		}
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return errors.Errorf("file %s cannot be read", filePath)
		}
		encodedFile := base64.StdEncoding.EncodeToString(file)
		if err != nil {
			return err
		}
		fmt.Print(encodedFile)
		return nil
	}
}
