package iox

import (
	"bufio"
	"io"
)

// ReadLines read line by line from reader, and return a error if read failed.
// the reader will be leaved unclosed
func ReadLines(reader io.Reader, consume func(line string)) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		consume(scanner.Text())
	}
	return scanner.Err()
}

// ReadAllLines read all data till EOF, return a lines slice
// the reader will be leaved unclosed
func ReadAllLines(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
