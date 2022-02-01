package futl

import (
	"encoding/base64"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func ParseFile(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.Errorf("file %s does not exist", filePath)
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.Errorf("file %s cannot be read", filePath)
	}
	encodedFile := base64.StdEncoding.EncodeToString(file)
	return encodedFile, nil
}
