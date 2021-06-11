package jsonx

import (
	"encoding/json"
	"github.com/hsiafan/glow/stringx"
)

// Marshal marshal value to json string
func Marshal(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return stringx.FromBytes(data), nil
}

// MarshalIndent is like Marshal but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
func MarshalIndent(v interface{}, prefix, indent string) (string, error) {
	data, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return "", err
	}
	return stringx.FromBytes(data), nil
}

// UnMarshal unmarshal json string to value
func UnMarshal(jsonStr string, v interface{}) error {
	return json.Unmarshal(stringx.ToBytes(jsonStr), v)
}
