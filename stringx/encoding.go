package stringx

import (
	"github.com/hsiafan/glow/unsafex"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// Encode str to bytes using specific encoding
func Encode(str string, enc encoding.Encoding) ([]byte, error) {
	if enc == unicode.UTF8 {
		return []byte(str), nil
	}
	return enc.NewEncoder().Bytes(unsafex.StringToBytes(str))
}

// Decode decode bytes to str using specific encoding
func Decode(data []byte, enc encoding.Encoding) (string, error) {
	if enc == unicode.UTF8 {
		return string(data), nil
	}
	bytes, err := enc.NewDecoder().Bytes(data)
	if err != nil {
		return "", err
	}
	return unsafex.BytesToString(bytes), err
}
