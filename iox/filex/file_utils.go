package filex

import (
	"github.com/hsiafan/glow/encodingx"
	"github.com/hsiafan/glow/iox"
	"github.com/hsiafan/glow/unsafex"
	"golang.org/x/text/encoding"
	"io/ioutil"
	"os"
)

// Read and return all data as string in file
func ReadAllToStringWithEncoding(path string, encoding encoding.Encoding) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return encodingx.Decode(data, encoding)
}

// Read and return all data as string in file
func ReadAllToString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return unsafex.BytesToString(data), err
}

// Read and return all data in file
func ReadAll(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Read line by line from a file with specific, and return a error if read failed.
func ReadLinesWithEncoding(path string, encoding encoding.Encoding, consume func(line string)) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer iox.Close(reader)
	return iox.ReadLinesWithEncoding(reader, encoding, consume)
}

// Read line by line from a file, and return a error if read failed.
func ReadLines(path string, consume func(line string)) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer iox.Close(reader)
	return iox.ReadLines(reader, consume)
}

// Read all data from a file till EOF, return a lines slice.
func ReadAllLinesWithEncoding(path string, encoding encoding.Encoding) ([]string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer iox.Close(reader)
	return iox.ReadAllLinesWithEncoding(reader, encoding)
}

// Read all data from a file till EOF, return a lines slice.
func ReadAllLines(path string) ([]string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer iox.Close(reader)
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

// Write all data to file, and then close. If file already exists, will be override
func Write(path string, data []byte) error {
	fi, err := os.Create(path)
	if err != nil {
		return err
	}
	defer iox.Close(fi)
	_, err = fi.Write(data)
	return err
}

// Write all string content to file, and then close. If file already exists, will be override
func WriteString(path string, str string) error {
	return Write(path, unsafex.StringToBytes(str))
}

// Write all string content to file using specific encoding, and then close. If file already exists, will be override
func WriteStringWithEncoding(path string, str string, encoding encoding.Encoding) error {
	data, err := encodingx.Encode(str, encoding)
	if err != nil {
		return err
	}
	return Write(path, data)
}
