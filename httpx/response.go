package httpx

import (
	"encoding/json"
	"github.com/hsiafan/glow/iox"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/encoding/unicode"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Http response, all info.
type ResponseContext struct {
	Response *http.Response
	Err      error
}

// Response without body...
type ResponseInfo struct {
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

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	Request *http.Request
}

// Read and discard all response body
func (r *ResponseContext) DiscardBody() (*ResponseInfo, int64, error) {
	return r.WriteTo(ioutil.Discard)
}

// Read all response body, to bytes
func (r *ResponseContext) ReadAll() (*ResponseInfo, []byte, error) {
	if r.Err != nil {
		return nil, nil, r.Err
	}
	defer iox.Close(r.Response.Body)
	data, err := ioutil.ReadAll(r.Response.Body)
	return r.toResponseInfo(), data, err
}

// Read all response body, to string
func (r *ResponseContext) ReadAllString() (*ResponseInfo, string, error) {
	if r.Err != nil {
		return nil, "", r.Err
	}
	enc := r.getEncoding()
	var reader io.Reader = r.Response.Body
	defer iox.Close(r.Response.Body)

	if enc != nil && enc != unicode.UTF8 {
		reader = enc.NewDecoder().Reader(reader)
	}

	content, err := iox.ReadAllToString(reader)
	return r.toResponseInfo(), content, err
}

// get encoding from response header
func (r *ResponseContext) getEncoding() encoding.Encoding {
	contentType := r.Response.Header.Get("Content-Type")
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
			} else {
				return enc
			}
		}
	}
	return nil
}

// Read all response body, write to target writer.
func (r *ResponseContext) WriteTo(w io.Writer) (*ResponseInfo, int64, error) {
	if r.Err != nil {
		return nil, 0, r.Err
	}
	defer iox.Close(r.Response.Body)
	written, err := io.Copy(w, r.Response.Body)
	return r.toResponseInfo(), written, err
}

// Read all response body, write to target writer.
func (r *ResponseContext) WriteToFile(path string) (*ResponseInfo, int64, error) {
	if r.Err != nil {
		return nil, 0, r.Err
	}
	defer iox.Close(r.Response.Body)
	f, err := os.Create(path)
	if err != nil {
		return r.toResponseInfo(), 0, err
	}
	defer iox.Close(f)
	written, err := io.Copy(f, r.Response.Body)
	return r.toResponseInfo(), written, err
}

func (r *ResponseContext) DecodeJson(v interface{}) (*ResponseInfo, error) {
	if r.Err != nil {
		return nil, r.Err
	}
	enc := r.getEncoding()
	var reader io.Reader = r.Response.Body
	defer iox.Close(r.Response.Body)

	if enc != nil && enc != unicode.UTF8 {
		reader = enc.NewDecoder().Reader(reader)
	}

	err := json.NewDecoder(reader).Decode(v)
	return r.toResponseInfo(), err
}

func (r *ResponseContext) toResponseInfo() *ResponseInfo {
	if r.Response == nil {
		return nil
	}
	return &ResponseInfo{
		Status:     r.Response.Status,
		StatusCode: r.Response.StatusCode,
		Proto:      r.Response.Proto,
		ProtoMajor: r.Response.ProtoMajor,
		ProtoMinor: r.Response.ProtoMinor,
		Header:     r.Response.Header,
		Trailer:    r.Response.Trailer,
		Request:    r.Response.Request,
	}
}
