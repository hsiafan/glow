package httpx

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/hsiafan/glow/iox"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/encoding/unicode"
)

// ResponseContext carry http response and errors.
type ResponseContext struct {
	Response *http.Response
	Err      error
}

// ResponseHeader is a http response without body...
type ResponseHeader struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0

	// Header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// Header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	Header http.Header

	// Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	Trailer http.Header
}

// DiscardBody read and discard all response body
func (r *ResponseContext) DiscardBody() (*ResponseHeader, int64, error) {
	return r.WriteToWriter(ioutil.Discard)
}

// ReadAll read all response body, to bytes
func (r *ResponseContext) ReadAll() (*ResponseHeader, []byte, error) {
	if r.Err != nil {
		return nil, nil, r.Err
	}
	defer iox.Close(r.Response.Body)
	data, err := ioutil.ReadAll(r.Response.Body)
	return r.toResponseHeader(), data, err
}

// ReadAllString read all response body, to string
func (r *ResponseContext) ReadAllString() (*ResponseHeader, string, error) {
	if r.Err != nil {
		return nil, "", r.Err
	}
	enc := r.GetEncoding()
	var reader io.Reader = r.Response.Body
	defer iox.Close(r.Response.Body)

	if enc != nil && enc != unicode.UTF8 {
		reader = enc.NewDecoder().Reader(reader)
	}

	content, err := iox.ReadAllString(reader)
	return r.toResponseHeader(), content, err
}

// GetEncoding get encoding from response header. If header not set charset for content-type, return nil.
func (r *ResponseContext) GetEncoding() encoding.Encoding {
	contentType := r.Response.Header.Get(HeaderContenttype)
	if contentType == "" {
		return nil
	}
	for _, item := range strings.Split(contentType, ";") {
		item = strings.TrimSpace(item)
		param := ParseParam(item)
		if strings.EqualFold(param.Name, "charset") {
			enc, err := htmlindex.Get(param.Value)
			if err != nil {
				return nil
			}
			return enc
		}
	}
	return nil
}

// Read all response body data, and write to target writer.
func (r *ResponseContext) WriteToWriter(w io.Writer) (*ResponseHeader, int64, error) {
	if r.Err != nil {
		return nil, 0, r.Err
	}
	defer iox.Close(r.Response.Body)
	written, err := io.Copy(w, r.Response.Body)
	return r.toResponseHeader(), written, err
}

// Read all response body, write to target writer.
func (r *ResponseContext) WriteToFile(path string) (*ResponseHeader, int64, error) {
	if r.Err != nil {
		return nil, 0, r.Err
	}
	defer iox.Close(r.Response.Body)
	f, err := os.Create(path)
	if err != nil {
		return r.toResponseHeader(), 0, err
	}
	defer iox.Close(f)
	written, err := io.Copy(f, r.Response.Body)
	return r.toResponseHeader(), written, err
}

// DecodeJSON decode http body as json, into a value.
func (r *ResponseContext) DecodeJSON(v interface{}) (*ResponseHeader, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	enc := r.GetEncoding()
	var reader io.Reader = r.Response.Body
	defer iox.Close(r.Response.Body)

	if enc != nil && enc != unicode.UTF8 {
		reader = enc.NewDecoder().Reader(reader)
	}

	err := json.NewDecoder(reader).Decode(v)
	return r.toResponseHeader(), err
}

func (r *ResponseContext) toResponseHeader() *ResponseHeader {
	if r.Response == nil {
		return nil
	}
	return &ResponseHeader{
		Status:     r.Response.Status,
		StatusCode: r.Response.StatusCode,
		Proto:      r.Response.Proto,
		ProtoMajor: r.Response.ProtoMajor,
		ProtoMinor: r.Response.ProtoMinor,
		Header:     r.Response.Header,
		Trailer:    r.Response.Trailer,
	}
}
