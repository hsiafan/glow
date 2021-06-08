package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/hsiafan/glow/iox"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// Body is interface for Http body, with content type
type Body interface {
	// MimeType return content mime type of this body
	MimeType() string

	// Encoding for body. may be nil
	Encoding() encoding.Encoding

	// GetReader get a reader for reading body data
	GetReader() (io.Reader, error)
}

type hasMimeType struct {
	mimeType string
}

func (h *hasMimeType) MimeType() string {
	return h.mimeType
}

type hasEncoding struct {
	enc encoding.Encoding
}

func (e *hasEncoding) Encoding() encoding.Encoding {
	return e.enc
}

var _ Body = (*ReaderBody)(nil)

// ReaderBody is a http body contains a io.Reader
type ReaderBody struct {
	hasMimeType
	hasEncoding
	reader io.Reader
}

func (r *ReaderBody) GetReader() (io.Reader, error) {
	return r.reader, nil
}

// NewBody create new Body from reader
func NewBody(reader io.Reader, contentType string) *ReaderBody {
	return NewBodyWithEncoding(reader, contentType, nil)
}

// NewBodyWithEncoding create new Body from reader
func NewBodyWithEncoding(reader io.Reader, contentType string, enc encoding.Encoding) *ReaderBody {
	return &ReaderBody{
		hasMimeType: hasMimeType{contentType},
		hasEncoding: hasEncoding{enc},
		reader:      reader,
	}
}

var _ Body = (*BytesBody)(nil)

// BytesBody is a http body contains byte array as content
type BytesBody struct {
	hasEncoding
	hasMimeType
	data []byte
}

// NewBytesBody create new Body from bytes, providing mimetype
func NewBytesBody(data []byte, contentType string) *BytesBody {
	return NewBytesBodyWithEncoding(data, contentType, nil)
}

// NewBytesBodyWithEncoding create new Body from bytes, and provide mimetype, and encoding
func NewBytesBodyWithEncoding(data []byte, contentType string, enc encoding.Encoding) *BytesBody {
	return &BytesBody{
		hasEncoding: hasEncoding{enc},
		hasMimeType: hasMimeType{contentType},
		data:        data,
	}
}

func (b *BytesBody) Encoding() encoding.Encoding {
	return nil
}

func (b *BytesBody) GetReader() (io.Reader, error) {
	return bytes.NewReader(b.data), nil
}

var _ Body = (*StringBody)(nil)

// StringBody is a http body has string value
type StringBody struct {
	hasMimeType
	hasEncoding
	content string
}

// NewStringBody create new Body from string content
func NewStringBody(content string, contentType string) *StringBody {
	return NewStringBodyWithEncoding(content, contentType, unicode.UTF8)
}

// NewStringBodyWithEncoding create new Body from string
func NewStringBodyWithEncoding(content string, contentType string, enc encoding.Encoding) *StringBody {
	return &StringBody{
		hasEncoding: hasEncoding{enc},
		hasMimeType: hasMimeType{contentType},
		content:     content,
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

// JSONBody is http body, marshal value as json
type JSONBody struct {
	hasEncoding
	value interface{}
}

// NewJSONBody create new Body from value, marshall to json
func NewJSONBody(value interface{}) *JSONBody {
	return NewJSONBodyWithEncoding(value, unicode.UTF8)
}

// NewJSONBodyWithEncoding create new Body from value, marshall to json
func NewJSONBodyWithEncoding(value interface{}, enc encoding.Encoding) *JSONBody {
	return &JSONBody{
		hasEncoding: hasEncoding{enc},
		value:       value,
	}
}

func (j *JSONBody) MimeType() string {
	return MimetypeJson
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

// FormBody is url-encoded-www form body
type FormBody struct {
	hasEncoding
	params []*Param
}

// NewFormBody create new form encoded Body from params
func NewFormBody(params ...*Param) *FormBody {
	return NewFormBodyWithEncoding(params, unicode.UTF8)
}

// NewFormBodyWithEncoding create new form encoded Body from params
func NewFormBodyWithEncoding(params []*Param, enc encoding.Encoding) *FormBody {
	return &FormBody{
		hasEncoding: hasEncoding{enc},
		params:      params,
	}
}

func (f *FormBody) MimeType() string {
	return MimetypeFormEncoded
}

func (f *FormBody) GetReader() (io.Reader, error) {
	if len(f.params) == 0 {
		return iox.EmptyReader(), nil
	}
	var buf strings.Builder
	err := EncodeParamsTo(buf, f.enc, f.params...)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(buf.String()), nil
}

var _ Body = (*MultiPartBody)(nil)

// MultiPartBody is http multi-part form body
type MultiPartBody struct {
	parts   []*Part
	reader  *io.PipeReader
	writer  *io.PipeWriter
	mwriter *multipart.Writer
}

// NewMultiPartBody create new multi part body
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

func (m *MultiPartBody) MimeType() string {
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

// A empty http body
type emptyBody struct {
}

func (e *emptyBody) MimeType() string {
	return ""
}

func (e *emptyBody) Encoding() encoding.Encoding {
	return nil
}

func (e *emptyBody) GetReader() (io.Reader, error) {
	return nil, nil
}

// EmptyBody return an empty http Body
func EmptyBody() Body {
	return &emptyBody{}
}
