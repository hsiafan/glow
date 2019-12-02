package charsetx

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type Encoding int

// convert string to bytes with specific encoding
func EncodeString(str string, encoding encoding.Encoding) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), encoding.NewEncoder())
	return ioutil.ReadAll(reader)
}

// convert bytes to string with specific encoding
func DecodeString(data []byte, encoding encoding.Encoding) (string, error) {
	reader := transform.NewReader(bytes.NewReader(data), encoding.NewDecoder())
	utf8Data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(utf8Data), err
}
