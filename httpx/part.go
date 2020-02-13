package httpx

import (
	"bytes"
	"io"
)

const (
	formPart = 0
	filePart = 2
)

type Part struct {
	_type int
	// form part
	name  string
	value string

	// file part
	filename string
	reader   io.Reader
}

func NewFormPart(filedName string, value string) *Part {
	return &Part{
		_type: formPart,
		name:  filedName,
		value: value,
	}
}

func NewFilePart(filedName string, filename string, reader io.Reader) *Part {
	return &Part{
		_type:    filePart,
		name:     filedName,
		filename: filename,
		reader:   reader,
	}
}

func NewBytesFilePart(filedName string, filename string, data []byte) *Part {
	return &Part{
		_type:    filePart,
		name:     filedName,
		filename: filename,
		reader:   bytes.NewReader(data),
	}
}
