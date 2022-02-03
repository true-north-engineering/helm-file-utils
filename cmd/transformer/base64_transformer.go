package transformer

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const (
	B64EncPrefix = "base64enc://"
)

func B64ENCTransform(filePath string) (interface{}, error) {
	filePath = strings.TrimPrefix(filePath, B64EncPrefix)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.Errorf("file %s does not exist", filePath)
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.Errorf("file %s cannot be read", filePath)
	}
	encodedFile := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(encodedFile, file)
	return encodedFile, nil
}
