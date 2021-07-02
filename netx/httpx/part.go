package httpx

import (
	"bytes"
	"io"
	"os"
)

type formType int

const (
	formPart formType = 0
	filePart formType = 2
)

// Part is one part of multi-part encoded body, can be key-value form field, or form file part.
type Part struct {
	_type formType
	// form part
	name  string // the name of this part
	value string // the value, used for form filed

	// file part
	filename string                    // the file name, used for form file
	reader   func() (io.Reader, error) // provide the reader to get file content, used for form file
}

// NewFormPart create one new key-value param Part.
func NewFormPart(name string, value string) *Part {
	return &Part{
		_type: formPart,
		name:  name,
		value: value,
	}
}

// NewFilePart create one new file Part.
// param readerProvider provide the reader contains the content of this part, should allow call multi times.
func NewFilePart(name string, filename string, readerProvider func() (io.Reader, error)) *Part {
	return &Part{
		_type:    filePart,
		name:     name,
		filename: filename,
		reader:   readerProvider,
	}
}

// NewBytesFilePart create one new file Part from binary array.
func NewBytesFilePart(name string, filename string, data []byte) *Part {
	return &Part{
		_type:    filePart,
		name:     name,
		filename: filename,
		reader: func() (io.Reader, error) {
			return bytes.NewReader(data), nil
		},
	}
}

// NewFSFilePart create one new file Part from a file in file system.
func NewFSFilePart(name string, filename string, path string) *Part {
	return &Part{
		_type:    filePart,
		name:     name,
		filename: filename,
		reader: func() (io.Reader, error) {
			return os.Open(path)
		},
	}
}
