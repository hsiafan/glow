package iox

import (
	"github.com/hsiafan/glow/unsafex"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
)

// ReadAllStringWithEncoding read and return all data as string in reader, with encoding the data used
func ReadAllStringWithEncoding(r io.Reader, encoding encoding.Encoding) (string, error) {
	r = encoding.NewDecoder().Reader(r)
	return ReadAllString(r)
}

// ReadAllString read and return all data as string in reader
func ReadAllString(r io.Reader) (string, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return unsafex.BytesToString(data), err
}

// ReadAll read and return all data in reader
func ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

// ForEachLine read all lines in reader with specific encoding, and call consume function,
// The line index pass to function start from 0.
// If and err occurred during read, return an error. If read all data succeed till and io.EOF, nil error will be
// returned.
func ForEachLineWithEncoding(r io.Reader, encoding encoding.Encoding, consume func(line string)) error {
	r = encoding.NewDecoder().Reader(r)
	return ForEachLine(r, consume)
}

// ForEachLine read all lines in reader, and call consume function, The line index pass to function start from 0.
// If and err occurred during read, return an error. If read all data succeed till and io.EOF, nil error will be
// returned.
func ForEachLine(r io.Reader, consume func(line string)) error {
	lr := NewLineReader(r)
	return lr.ForEachLine(consume)
}

// ReadAllLinesWithEncoding read all data as string line with specific encoding till EOF, return a lines slice.
// The reader will be left unclosed
func ReadAllLinesWithEncoding(r io.Reader, enc encoding.Encoding) ([]string, error) {
	r = enc.NewDecoder().Reader(r)
	return ReadAllLines(r)
}

// ReadAllLines read all data as string lines, till EOF, return a lines slice.
// The reader will be left unclosed
func ReadAllLines(r io.Reader) ([]string, error) {
	lr := NewLineReader(r)
	return lr.ReadAllLines()
}

var emptyR = &emptyReader{}

type emptyReader struct {
}

func (e *emptyReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

// EmptyReader return a empty reader
func EmptyReader() io.Reader {
	return emptyR
}
