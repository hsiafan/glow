package iox

import (
	"bufio"
	"github.com/hsiafan/glow/unsafex"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
)

// Read and return all data as string in reader
func ReadAllToStringWithEncoding(reader io.Reader, encoding encoding.Encoding) (string, error) {
	reader = encoding.NewDecoder().Reader(reader)
	return ReadAllToString(reader)
}

// Read and return all data as string in reader
func ReadAllToString(reader io.Reader) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return unsafex.BytesToString(data), err
}

// Read and return all data in reader
func ReadAll(reader io.Reader) ([]byte, error) {
	return ioutil.ReadAll(reader)
}

// ReadLines read line by line from reader with specific encoding, and return a error if read failed.
// The reader will be leaved unclosed
func ReadLinesWithEncoding(reader io.Reader, encoding encoding.Encoding, consume func(line string)) error {
	reader = encoding.NewDecoder().Reader(reader)
	return ReadLines(reader, consume)
}

// ReadLines read line by line from reader, and return a error if read failed.
// The reader will be leaved unclosed
func ReadLines(reader io.Reader, consume func(line string)) error {
	// scanner can only handle line blow max len, which is bufio.MaxScanTokenSize(65536)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		consume(scanner.Text())
	}
	return scanner.Err()
}

// ReadAllLines read all data as string line with specific encoding till EOF, return a lines slice.
// The reader will be leaved unclosed
func ReadAllLinesWithEncoding(reader io.Reader, encoding encoding.Encoding) ([]string, error) {
	reader = encoding.NewDecoder().Reader(reader)
	return ReadAllLines(reader)
}

// ReadAllLines read all data as string lines, till EOF, return a lines slice.
// The reader will be leaved unclosed
func ReadAllLines(reader io.Reader) ([]string, error) {
	var lines []string
	err := ReadLines(reader, func(line string) {
		lines = append(lines, line)
	})
	return lines, err
}
