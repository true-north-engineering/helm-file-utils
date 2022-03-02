package reader

import (
	"io/ioutil"
	"net/http"
)

const (
	HttpsPrefix = "https"
	HttpPrefix  = "http"
)

func ReadHttps(httpPath string) (InputValue, error) {

	resp, err := http.Get(httpPath)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := InputValue{Kind: InputKindFile, Value: make(map[string][]byte)}
	result.Value[InputKindFile] = body
	return result, err
}
