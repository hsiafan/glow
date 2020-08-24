package filex

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/hsiafan/glow/iox"
	"github.com/hsiafan/glow/unsafex"
	"golang.org/x/text/encoding"
)

// Read and return all data as string in file
func ReadAllToStringWithEncoding(path string, enc encoding.Encoding) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	decoded, err := enc.NewDecoder().Bytes(data)
	if err != nil {
		return "", err
	}
	return string(decoded), err
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
func ReadAllBytes(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Read all data from file, into a writer
func ReadAllToWriter(path string, w io.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	return err
}

// ForEachLineWithEncoding read line by line from a file with specific, and return a error if read failed.
func ForEachLineWithEncoding(path string, enc encoding.Encoding, consume func(line string)) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer iox.Close(reader)
	return iox.ForEachLineWithEncoding(reader, enc, consume)
}

// ForEachLine read line by line from a file, and return a error if read failed.
func ForEachLine(path string, consume func(line string)) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer iox.Close(reader)
	return iox.ForEachLine(reader, consume)
}

// Read all data from a file till EOF, return a lines slice.
func ReadAllLinesWithEncoding(path string, enc encoding.Encoding) ([]string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer iox.Close(reader)
	return iox.ReadAllLinesWithEncoding(reader, enc)
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
func WriteBytes(path string, data []byte) error {
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
	return WriteBytes(path, unsafex.StringToBytes(str))
}

// Write all string content to file using specific encoding, and then close. If file already exists, will be override
func WriteStringWithEncoding(path string, str string, enc encoding.Encoding) error {
	data, err := enc.NewEncoder().Bytes(unsafex.StringToBytes(str))
	if err != nil {
		return err
	}
	return WriteBytes(path, data)
}

// Write all data from a reader, to a file.
func WriteAllFromReader(path string, r io.Reader) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer iox.Close(f)
	_, err = io.Copy(f, r)
	return err
}
