package iox

import (
	"bytes"
	"io"
)

const defaultBufSize = 4096

// LineReader is a read which can read line by line, without limitation for line length.
type LineReader struct {
	reader     io.Reader
	buf        []byte // the buffer
	start, end int    // current buffer has data, new write will be start and end pos
	err        error  // read error
}

// NewLineReaderSize create new LineReader with initial buf size
func NewLineReaderSize(r io.Reader, initialSize int) *LineReader {
	if initialSize <= 0 {
		panic("buf size should larger than 0")
	}
	return &LineReader{
		reader: r,
		buf:    make([]byte, initialSize),
	}
}

// NewLineReader create new LineReader
func NewLineReader(r io.Reader) *LineReader {

	return &LineReader{
		reader: r,
		buf:    make([]byte, defaultBufSize),
	}
}

func (r *LineReader) readLine() ([]byte, error) {
	searchStart := r.start
	for {
		if pos := bytes.IndexRune(r.buf[searchStart:r.end], '\n'); pos >= 0 {
			pos = searchStart + pos
			b := r.buf[r.start:pos]
			r.start = pos + 1
			return b, nil
		}

		if r.err != nil {

			var line []byte
			if r.err == io.EOF {
				line = r.buf[r.start:r.end]
			}
			r.start = r.end
			searchStart = r.end
			return line, r.err
		}

		// slide existing data to beginning.
		if r.start > 0 {
			copy(r.buf, r.buf[r.start:r.end])
			r.end = r.end - r.start
			r.start = 0
		}
		searchStart = r.end

		// expand buf if needed
		if r.end >= len(r.buf) {
			newBuf := make([]byte, len(r.buf)*2)
			copy(newBuf, r.buf)
			r.buf = newBuf
		}

		// read data
		n, err := r.reader.Read(r.buf[r.end:])
		r.end += n
		if err != nil {
			r.err = err
		}
	}
}

// ForEachLine read all lines in reader, and call consume function, The line index pass to function start from 0.
// If and err occurred during read, return an error. If read all data succeed till and io.EOF, nil error will be
// returned.
func (r *LineReader) ForEachLine(consume func(line string)) error {
	for {
		line, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		consume(line)
	}
}

// ReadAllLine read all lines in reader, return lines, and an error if any error occurred while read.
// If read all data succeed till and io.EOF, nil error will be returned.
func (r *LineReader) ReadAllLines() ([]string, error) {
	var lines []string
	for {
		line, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				return lines, nil
			}
			return lines, err
		}
		lines = append(lines, line)
	}
}

// ReadLine read and return one line, not including the end-of-line bytes.
// ReadLine either returns a non-nil line or it returns an error, never both.
// No indication or error is given if the input ends without a final line end.
// An io.EOF error would be returned if already read to the end of reader.
func (r *LineReader) ReadLine() (string, error) {
	line, err := r.readLine()

	if len(line) > 0 && line[len(line)-1] == '\r' {
		line = line[:len(line)-1]
	}

	if len(line) == 0 {
		return "", err
	}

	return string(line), nil
}
