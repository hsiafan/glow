package filex

import (
	"github.com/hsiafan/glow/stringx"
	"io"
	"os"

	"github.com/hsiafan/glow/iox"
	"golang.org/x/text/encoding"
)

// ReadAllToStringWithEncoding read and return all data as string in file
func ReadAllToStringWithEncoding(path string, enc encoding.Encoding) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	decoded, err := enc.NewDecoder().Bytes(data)
	if err != nil {
		return "", err
	}
	return string(decoded), err
}

// ReadAllToString read and return all data as string in file
func ReadAllToString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return stringx.FromBytes(data), err
}

// ReadAllBytes read and return all data in file
func ReadAllBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ReadAllToWriter read all data from file, into a writer
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

// ReadAllLinesWithEncoding read all data from a file till EOF, return a lines slice.
func ReadAllLinesWithEncoding(path string, enc encoding.Encoding) ([]string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer iox.Close(reader)
	return iox.ReadAllLinesWithEncoding(reader, enc)
}

// ReadAllLines read all data from a file till EOF, return a lines slice.
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

// IsFile return true if path is exists and is regular file
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

// IsDirectory return true if path is exists and is directory
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

// WriteBytes write all data to file, and then close. If file already exists, will be override
func WriteBytes(path string, data []byte) error {
	fi, err := os.Create(path)
	if err != nil {
		return err
	}
	defer iox.Close(fi)
	_, err = fi.Write(data)
	return err
}

// WriteString write all string content to file, and then close. If file already exists, will be override
func WriteString(path string, str string) error {
	return WriteBytes(path, stringx.ToBytes(str))
}

// WriteStringWithEncoding write all string content to file using specific encoding, and then close.
// If file already exists, will be override
func WriteStringWithEncoding(path string, str string, enc encoding.Encoding) error {
	data, err := enc.NewEncoder().Bytes(stringx.ToBytes(str))
	if err != nil {
		return err
	}
	return WriteBytes(path, data)
}

// WriteAllFromReader write all data from a reader, to a file.
func WriteAllFromReader(path string, r io.Reader) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer iox.Close(f)
	_, err = io.Copy(f, r)
	return err
}
