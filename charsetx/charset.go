package charsetx

import (
	"golang.org/x/text/encoding"
)

type Encoding int

// convert string to bytes with specific encoding
func EncodeString(str string, encoding encoding.Encoding) ([]byte, error) {
	return encoding.NewEncoder().Bytes([]byte(str))
}

// convert bytes to string with specific encoding
func DecodeString(data []byte, encoding encoding.Encoding) (string, error) {
	data, err := encoding.NewDecoder().Bytes(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
