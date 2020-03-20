package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/hsiafan/glow/httpx/mimetype"
	"github.com/hsiafan/glow/iox"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// Http body, with content type
type Body interface {
	// return content type of this body
	ContentType() string

	// Encoding for body. may be nil
	Encoding() encoding.Encoding

	// get reader from body
	GetReader() (io.Reader, error)
}

type hasContentType struct {
	contentType string
}

func (h *hasContentType) ContentType() string {
	return h.contentType
}

type hasEncoding struct {
	enc encoding.Encoding
}

func (e *hasEncoding) Encoding() encoding.Encoding {
	return e.enc
}

var _ Body = (*ReaderBody)(nil)

// Reader body
type ReaderBody struct {
	hasContentType
	hasEncoding
	reader io.Reader
}

func (r *ReaderBody) GetReader() (io.Reader, error) {
	return r.reader, nil
}

// Create new Body from reader
func NewBody(reader io.Reader, contentType string) *ReaderBody {
	return NewBodyWithEncoding(reader, contentType, nil)
}

// Create new Body from reader
func NewBodyWithEncoding(reader io.Reader, contentType string, enc encoding.Encoding) *ReaderBody {
	return &ReaderBody{
		hasContentType: hasContentType{contentType},
		hasEncoding:    hasEncoding{enc},
		reader:         reader,
	}
}

var _ Body = (*BytesBody)(nil)

// Byte Array as Body
type BytesBody struct {
	hasEncoding
	hasContentType
	data []byte
}

// Create new Body from bytes
func NewBytesBody(data []byte, contentType string) *BytesBody {
	return NewBytesBodyWithEncoding(data, contentType, nil)
}

// Create new Body from bytes
func NewBytesBodyWithEncoding(data []byte, contentType string, enc encoding.Encoding) *BytesBody {
	return &BytesBody{
		hasEncoding:    hasEncoding{enc},
		hasContentType: hasContentType{contentType},
		data:           data,
	}
}

func (b *BytesBody) Encoding() encoding.Encoding {
	return nil
}

func (b *BytesBody) GetReader() (io.Reader, error) {
	return bytes.NewReader(b.data), nil
}

var _ Body = (*StringBody)(nil)

// String as Body
type StringBody struct {
	hasContentType
	hasEncoding
	content string
}

// Create new Body from bytes
func NewStringBody(content string, contentType string) *StringBody {
	return NewStringBodyWithEncoding(content, contentType, unicode.UTF8)
}

// Create new Body from bytes
func NewStringBodyWithEncoding(content string, contentType string, enc encoding.Encoding) *StringBody {
	return &StringBody{
		hasEncoding:    hasEncoding{enc},
		hasContentType: hasContentType{contentType},
		content:        content,
	}
}

func (s *StringBody) GetReader() (io.Reader, error) {
	reader := strings.NewReader(s.content)
	if s.enc == nil || s.enc == unicode.UTF8 {
		return reader, nil
	}
	return s.enc.NewDecoder().Reader(reader), nil
}

var _ Body = (*JSONBody)(nil)

// Body marshal value as json
type JSONBody struct {
	hasEncoding
	value interface{}
}

// Create new Body from value, marshall to json
func NewJSONBody(value interface{}) *JSONBody {
	return NewJSONBodyWithEncoding(value, unicode.UTF8)
}

// Create new Body from value, marshall to json
func NewJSONBodyWithEncoding(value interface{}, enc encoding.Encoding) *JSONBody {
	return &JSONBody{
		hasEncoding: hasEncoding{enc},
		value:       value,
	}
}

func (j *JSONBody) ContentType() string {
	return mimetype.JSON
}

func (j *JSONBody) GetReader() (io.Reader, error) {
	data, err := json.Marshal(j.value)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)
	if j.enc == nil || j.enc == unicode.UTF8 {
		return reader, nil
	}
	return j.enc.NewDecoder().Reader(reader), nil
}

var _ Body = (*FormBody)(nil)

// URL encoded form body
type FormBody struct {
	hasEncoding
	params []*Param
}

// Create new form encoded Body from params
func NewFormBody(params ...*Param) *FormBody {
	return NewFormBodyWithEncoding(params, unicode.UTF8)
}

// Create new form encoded Body from params
func NewFormBodyWithEncoding(params []*Param, enc encoding.Encoding) *FormBody {
	return &FormBody{
		hasEncoding: hasEncoding{enc},
		params:      params,
	}
}

func (f *FormBody) ContentType() string {
	return mimetype.FormEncoded
}

func (f *FormBody) GetReader() (io.Reader, error) {
	if len(f.params) == 0 {
		return iox.EmptyReader(), nil
	}
	var buf strings.Builder
	err := EncodeParamsTo(f.params, f.enc, buf)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(buf.String()), nil
}

var _ Body = (*MultiPartBody)(nil)

type MultiPartBody struct {
	parts   []*Part
	reader  *io.PipeReader
	writer  *io.PipeWriter
	mwriter *multipart.Writer
}

// Create new multi part body
func NewMultiPartBody(parts []*Part) *MultiPartBody {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	return &MultiPartBody{
		parts:   parts,
		reader:  pr,
		mwriter: writer,
	}
}

func (m *MultiPartBody) Encoding() encoding.Encoding {
	return nil
}

func (m *MultiPartBody) ContentType() string {
	return m.mwriter.FormDataContentType()
}

func (m *MultiPartBody) GetReader() (io.Reader, error) {
	go func() {
		defer func() {
			for _, part := range m.parts {
				if part._type == filePart {
					if closer, ok := part.reader.(io.Closer); ok {
						_ = closer.Close()
					}
				}
			}
		}()
		for _, part := range m.parts {
			switch part._type {
			case formPart:
				err := m.mwriter.WriteField(part.name, part.value)
				if err != nil {
					_ = m.writer.CloseWithError(err)
					return
				}
			case filePart:
				partWriter, err := m.mwriter.CreateFormFile(part.name, part.filename)
				if err != nil {
					_ = m.writer.CloseWithError(err)
					return
				}

				_, err = io.Copy(partWriter, part.reader)
				if err != nil {
					_ = m.writer.CloseWithError(err)
					return
				}
			default:
				panic("unknown part type: " + strconv.Itoa(part._type))
			}
		}
	}()
	return m.reader, nil
}
