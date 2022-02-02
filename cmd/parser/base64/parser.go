package base64

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const (
	ENCPrefix = "base64enc://"
)

// ParseFile encode a byte array to base64 string
// receives:
// - filePath: path to file that will be encoded
// returns string base64 encoded and error.
func ParseFile(filePath string) (string, error) {
	filePath = strings.TrimPrefix(filePath, ENCPrefix)
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
