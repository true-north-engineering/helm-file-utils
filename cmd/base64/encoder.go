package base64

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// EncodeFile encode a byte array to base64 string
// receives:
// - filePath: path to file that will be encoded
// returns string base64 encoded and error.
func EncodeFile(filePath string) (string, error) {
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
