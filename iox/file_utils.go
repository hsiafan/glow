package iox

import (
	"os"
)

// ReadLinesFromFile read line by line from a file, and return a error if read failed.
func ReadLinesFromFile(path string, consume func(line string)) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()
	return ReadLines(reader, consume)
}

// ReadAllLines read all data from a file till EOF, return a lines slice.
// the reader will be leaved unclosed
func ReadAllLinesFromFile(path string) ([]string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return ReadAllLines(reader)
}

// Exists check if file exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
