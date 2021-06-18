package jsonx

import (
	"bytes"
	"encoding/json"
	"github.com/hsiafan/glow/stringx"
)

// Marshal marshal value to json string, without html escape
func Marshal(v interface{}) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	data := buf.Bytes()
	l := len(data)
	if l > 0 && data[l-1] == '\n' {
		data = data[:l-1]
	}
	return stringx.FromBytes(data), err
}

// MarshalIndent is like Marshal but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
func MarshalIndent(v interface{}, indent string) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", indent)
	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	data := buf.Bytes()
	l := len(data)
	if l > 0 && data[l-1] == '\n' {
		data = data[:l-1]
	}
	return stringx.FromBytes(data), err
}

// UnMarshal unmarshal json string to value
func UnMarshal(jsonStr string, v interface{}) error {
	return json.Unmarshal(stringx.ToBytes(jsonStr), v)
}
