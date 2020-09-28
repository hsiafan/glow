package httpx

import (
	"bytes"
	"io"
)

const (
	formPart = 0
	filePart = 2
)

// Part is one part of multi-part encoded body, can be key-value param form part, or file part.
type Part struct {
	_type int
	// form part
	name  string
	value string

	// file part
	filename string
	reader   io.Reader
}

// NewFormPart create one new key-value param Part.
func NewFormPart(filedName string, value string) *Part {
	return &Part{
		_type: formPart,
		name:  filedName,
		value: value,
	}
}

// NewFilePart create one new file Part.
func NewFilePart(filedName string, filename string, reader io.Reader) *Part {
	return &Part{
		_type:    filePart,
		name:     filedName,
		filename: filename,
		reader:   reader,
	}
}

// NewBytesFilePart create one new file Part from binary array.
func NewBytesFilePart(filedName string, filename string, data []byte) *Part {
	return &Part{
		_type:    filePart,
		name:     filedName,
		filename: filename,
		reader:   bytes.NewReader(data),
	}
}
