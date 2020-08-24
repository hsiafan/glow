package vlog

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

type wrappedError struct {
	message string
	cause   error
}

func (we *wrappedError) Error() string {
	return we.message + ": " + we.cause.Error()
}

func wrapError(message string, err error) error {
	return &wrappedError{message: message, cause: err}
}

func concatBytes(items ...[]byte) []byte {
	var l = 0
	for _, item := range items {
		l += len(item)
	}

	var result = make([]byte, l)
	var copied = 0
	for _, item := range items {
		copy(result[copied:], item)
		copied += len(item)
	}
	return result
}

func parseSize(sizeStr string) (int64, error) {
	if sizeStr == "" {
		return 0, errors.New("empty string can not convert to size")
	}
	lastChar := sizeStr[len(sizeStr)-1]
	lastChar = byte(unicode.ToLower(rune(lastChar)))
	if lastChar == 'k' || lastChar == 'm' || lastChar == 'g' || lastChar == 't' {
		v, err := strconv.ParseFloat(sizeStr[:len(sizeStr)-1], 64)
		if err != nil {
			return 0, err
		}
		switch lastChar {
		case 'k':
			return int64(v * 1024), nil
		case 'm':
			return int64(v * 1024 * 1024), nil
		case 'g':
			return int64(v * 1024 * 1024 * 1024), nil
		case 't':
			return int64(v * 1024 * 1024 * 1024 * 1024), nil
		}
	}
	return strconv.ParseInt(sizeStr, 10, 64)
}

// for files

func makeParentDirs(path string) error {
	dir, _ := filepath.Split(path)
	if len(dir) > 0 {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}
	return nil
}

//open file, create file and parent dirs if need
func openFile(path string) (*os.File, error) {
	if err := makeParentDirs(path); err != nil {
		return nil, wrapError("make file log parent dir failed", err)
	}
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}
