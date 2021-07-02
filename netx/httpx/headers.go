package httpx

import (
	"net/http"
	"time"
)

// header names
const (
	HeaderAccept                        string = "Accept"
	HeaderAcceptCharset                 string = "Accept-Charset"
	HeaderAcceptFeatures                string = "Accept-Features"
	HeaderAcceptEncoding                string = "Accept-Encoding"
	HeaderAcceptLanguage                string = "Accept-Language"
	HeaderAcceptRanges                  string = "Accept-Ranges"
	HeaderAccessControlAllowCredentials string = "Access-Control-Allow-Credentials"
	HeaderAccessControlAllowOrigin      string = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     string = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     string = "Access-Control-Allow-Headers"
	HeaderAccessControlMaxAge           string = "Access-Control-Max-Age"
	HeaderAccessControlExposeHeaders    string = "Access-Control-Expose-Headers"
	HeaderAccessControlRequestMethod    string = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   string = "Access-Control-Request-Headers"
	HeaderAge                           string = "Age"
	HeaderAllow                         string = "Allow"
	HeaderAlternates                    string = "Alternates"
	HeaderAuthorization                 string = "Authorization"
	HeaderCacheControl                  string = "Cache-Control"
	HeaderConnection                    string = "Connection"
	HeaderContentEncoding               string = "Content-Encoding"
	HeaderContentLanguage               string = "Content-Language"
	HeaderContentLength                 string = "Content-Length"
	HeaderContentLocation               string = "Content-Location"
	HeaderContentMD5                    string = "Content-MD5"
	HeaderContentRange                  string = "Content-Range"
	HeaderContentSecurityPolicy         string = "Content-Security-Policy"
	HeaderContenttype                   string = "Content-Type"
	HeaderCookie                        string = "Cookie"
	HeaderDNT                           string = "DNT"
	HeaderDate                          string = "Date"
	HeaderETag                          string = "ETag"
	HeaderExpect                        string = "Expect"
	HeaderExpires                       string = "Expires"
	HeaderFrom                          string = "From"
	HeaderHost                          string = "Host"
	HeaderIfMatch                       string = "If-Match"
	HeaderIfModifiedSince               string = "If-Modified-Since"
	HeaderIfNoneMatch                   string = "If-None-Match"
	HeaderIfRange                       string = "If-Range"
	HeaderIfUnmodifiedSince             string = "If-Unmodified-Since"
	HeaderLastEventID                   string = "Last-Event-ID"
	HeaderLastModified                  string = "Last-Modified"
	HeaderLink                          string = "Link"
	HeaderLocation                      string = "Location"
	HeaderMaxForwards                   string = "Max-Forwards"
	HeaderNegotiate                     string = "Negotiate"
	HeaderOrigin                        string = "Origin"
	HeaderPragma                        string = "Pragma"
	HeaderProxyAuthenticate             string = "Proxy-Authenticate"
	HeaderProxyAuthorization            string = "Proxy-Authorization"
	HeaderRange                         string = "Range"
	HeaderReferer                       string = "Referer"
	HeaderRetryAfter                    string = "Retry-After"
	HeaderSecWebsocketExtensions        string = "Sec-Websocket-Extensions"
	HeaderSecWebsocketKey               string = "Sec-Websocket-Key"
	HeaderSecWebsocketOrigin            string = "Sec-Websocket-Origin"
	HeaderSecWebsocketProtocol          string = "Sec-Websocket-Protocol"
	HeaderSecWebsocketVersion           string = "Sec-Websocket-Version"
	HeaderServer                        string = "Server"
	HeaderSetCookie                     string = "Set-Cookie"
	HeaderSetCookie2                    string = "Set-Cookie2"
	HeaderStrictTransportSecurity       string = "Strict-Transport-Security"
	HeaderTransferEncoding              string = "Transfer-Encoding"
	HeaderUpgrade                       string = "Upgrade"
	HeaderUserAgent                     string = "User-Agent"
	HeaderVariantVary                   string = "Variant-Vary"
	HeaderVary                          string = "Vary"
	HeaderVia                           string = "Via"
	HeaderWarning                       string = "Warning"
	HeaderWWWAuthenticate               string = "WWW-Authenticate"
	HeaderXContentDuration              string = "X-Content-Duration"
	HeaderXContentSecurityPolicy        string = "X-Content-Security-Policy"
	HeaderXDNSPrefetchControl           string = "X-DNSPrefetch-Control"
	HeaderXFrameOptions                 string = "X-Frame-Options"
	HeaderXRequestedWith                string = "X-Requested-With"
)

// Header is one http header. including name and value.
type Header struct {
	Name  string
	Value string
}

// Unpack return contents for convenient assign to multi variables.
func (h *Header) Unpack() (name, value string) {
	return h.Name, h.Value
}

// NewHeader create one new header
func NewHeader(name, value string) *Header {
	return &Header{name, value}
}

// NewDateHeader create one new header, with date time value.
func NewDateHeader(name string, value time.Time) *Header {
	return &Header{name, FormatDateHeader(value)}
}

// ParseDateHeader parse date header value, and return time in current Location.
func ParseDateHeader(dateStr string) (time.Time, error) {
	d, err := time.Parse(http.TimeFormat, dateStr)
	if err != nil {
		return d, err
	}
	return d.In(time.Local), err
}

// FormatDateHeader convert date to http header value.
func FormatDateHeader(date time.Time) string {
	return date.In(time.UTC).Format(http.TimeFormat)
}
