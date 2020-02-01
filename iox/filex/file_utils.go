package filex

import (
	"github.com/hsiafan/glow/iox"
	"os"
)

// Read line by line from a file, and return a error if read failed.
func ReadLines(path string, consume func(line string)) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer iox.CloseQuite(reader)
	return iox.ReadLines(reader, consume)
}

// Read all data from a file till EOF, return a lines slice.
func ReadAllLines(path string) ([]string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer iox.CloseQuite(reader)
	return iox.ReadAllLines(reader)
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

// Return true if path is exists and is regular file
func IsFile(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return true, nil
	default:
		return false, nil
	}
}

// Return true if path is exists and is directory
func IsDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	default:
		return false, nil
	}
}
