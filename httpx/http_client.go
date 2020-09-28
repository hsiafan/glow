package httpx

import (
	"crypto/tls"
	"errors"
	"github.com/hsiafan/glow/timex"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
)

// Client means Http Client
type Client struct {
	client    *http.Client
	dialer    *net.Dialer
	transport *http.Transport
	tlsConfig *tls.Config
	userAgent string
}

// NewClient create new http client
func NewClient(options ...ClientOption) *Client {
	dialer := &net.Dialer{
		Timeout:   timex.SecondsDuration(10),
		KeepAlive: timex.SecondsDuration(15), // tcp keep alive
	}
	tlsConfig := &tls.Config{}
	transport := &http.Transport{
		DialContext:       dialer.DialContext,
		DisableKeepAlives: false,
		ForceAttemptHTTP2: true,
		TLSClientConfig:   tlsConfig,
	}
	client := &Client{
		dialer:    dialer,
		transport: transport,
		tlsConfig: tlsConfig,
		client: &http.Client{
			Transport: transport,
			Timeout:   timex.MinutesDuration(2),
		},
	}
	for _, option := range options {
		option(client)
	}

	return client
}

// Head send a head request
func (c *Client) Head(url string, options ...RequestOption) *ResponseContext {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return &ResponseContext{Err: err}
	}
	return c.Send(req, options...)
}

// Get send a get request
func (c *Client) Get(url string, options ...RequestOption) *ResponseContext {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &ResponseContext{Err: err}
	}
	return c.Send(req, options...)
}

// Delete send a delete request
func (c *Client) Delete(url string, options ...RequestOption) *ResponseContext {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return &ResponseContext{Err: err}
	}
	return c.Send(req, options...)
}

// Put put send a put request with body
func (c *Client) Put(url string, body Body, options ...RequestOption) *ResponseContext {
	if body == nil {
		body = EmptyBody()
	}
	reader, err := body.GetReader()
	if err != nil {
		return &ResponseContext{Err: err}
	}
	req, err := http.NewRequest("PUT", url, reader)
	if err != nil {
		return &ResponseContext{Err: err}
	}

	contentType, err := c.makeContentType(body.MimeType(), body.Encoding())
	if err != nil {
		return &ResponseContext{Err: err}
	}
	if contentType != "" {
		req.Header.Set(HeaderContenttype, contentType)
	}
	return c.Send(req, options...)
}

// Post send a post request with body
func (c *Client) Post(url string, body Body, options ...RequestOption) *ResponseContext {
	if body == nil {
		body = EmptyBody()
	}
	reader, err := body.GetReader()
	if err != nil {
		return &ResponseContext{Err: err}
	}
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return &ResponseContext{Err: err}
	}
	contentType, err := c.makeContentType(body.MimeType(), body.Encoding())
	if err != nil {
		return &ResponseContext{Err: err}
	}
	if contentType != "" {
		req.Header.Set(HeaderContenttype, contentType)
	}
	return c.Send(req, options...)
}

func (c *Client) makeContentType(contentType string, encoding encoding.Encoding) (string, error) {
	if contentType == "" {
		return "", nil
	}
	if encoding == nil {
		return contentType, nil
	}
	encName, err := htmlindex.Name(encoding)
	if err != nil {
		return "", err
	}
	return contentType + "; charset=" + encName, nil
}

// Send send a http request with options.
func (c *Client) Send(r *http.Request, options ...RequestOption) *ResponseContext {
	if c.userAgent != "" {
		r.Header.Set(HeaderUseragent, c.userAgent)
	}
	for _, option := range options {
		if err := option(r); err != nil {
			return &ResponseContext{Err: err}
		}
	}
	resp, err := c.client.Do(r)
	return &ResponseContext{resp, err}
}

// ClientOption for setting http client option.
type ClientOption func(client *Client)

// RequestTimeout set http client timeout for one request
func RequestTimeout(timeout time.Duration) ClientOption {
	return func(client *Client) {
		client.client.Timeout = timeout
	}
}

// DialTimeout set tcp dial timeout
func DialTimeout(timeout time.Duration) ClientOption {
	return func(client *Client) {
		client.dialer.Timeout = timeout
	}
}

// EnableCookie enable auto cookie handle (receive, store, send)
func EnableCookie() ClientOption {
	return func(client *Client) {
		// New cookie jar should always succeed
		cookieJar, _ := cookiejar.New(nil)
		client.client.Jar = cookieJar
	}
}

// DisableCompression prevents the client from requesting compression with an "Accept-Encoding: gzip"
func DisableCompression() ClientOption {
	return func(client *Client) {
		client.transport.DisableCompression = true
	}
}

// DisableFollowRedirects disable redirect handle
func DisableFollowRedirects() ClientOption {
	return func(client *Client) {
		client.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return errors.New("redirect disabled")
		}
	}
}

// MaxIdleConns set max count for all hosts of this client. Default is no limit
func MaxIdleConns(count int) ClientOption {
	return func(client *Client) {
		client.transport.MaxIdleConns = count
	}
}

// MaxIdleConnsPerHost set max idle connection count for one host. Default is no limit
func MaxIdleConnsPerHost(count int) ClientOption {
	return func(client *Client) {
		client.transport.MaxIdleConnsPerHost = count
	}
}

// MaxConnsPerHost set max connection count for one host. Default is no limit
func MaxConnsPerHost(count int) ClientOption {
	return func(client *Client) {
		client.transport.MaxConnsPerHost = count
	}
}

// IdleConnTimeout set max duration for close idle connections. Default is not timeout
func IdleConnTimeout(timeout time.Duration) ClientOption {
	return func(client *Client) {
		client.transport.IdleConnTimeout = timeout
	}
}

// ResponseHeaderTimeout set max duration for wait response header after write all request. Default is no timeout.
func ResponseHeaderTimeout(timeout time.Duration) ClientOption {
	return func(client *Client) {
		client.transport.ResponseHeaderTimeout = timeout
	}
}

// ExpectContinueTimeout set max duration for waiting a server's first response headers after fully
// writing the request headers if the request has an "Expect: 100-continue" header.
// Default is not timeout
func ExpectContinueTimeout(timeout time.Duration) ClientOption {
	return func(client *Client) {
		client.transport.ExpectContinueTimeout = timeout
	}
}

// DisableHTTPKeepAlive disable Http keep alive(which means disable http connection reuse)
func DisableHTTPKeepAlive() ClientOption {
	return func(client *Client) {
		client.transport.DisableKeepAlives = true
	}
}

// DisableTCPKeepAlive disable tcp connection keep alive
func DisableTCPKeepAlive() ClientOption {
	return func(client *Client) {
		client.dialer.KeepAlive = -1
	}
}

// DisableTLSVerify disable tls certificate chain and server name verifying.
func DisableTLSVerify() ClientOption {
	return func(client *Client) {
		client.tlsConfig.InsecureSkipVerify = true
	}
}

// UserAgent set use-agent for this client
func UserAgent(userAgent string) ClientOption {
	return func(client *Client) {
		client.userAgent = userAgent
	}
}

// UseProxy set proxy by proxy url.
// The proxy type is determined by the URL scheme. "http", "https", and "socks5" are supported.
// If the scheme is empty, "http" is assumed.
// If Proxy url parse error, no proxy is used.
func UseProxy(proxy string) ClientOption {
	return func(client *Client) {
		_url, err := url.Parse(proxy)
		client.transport.Proxy = func(request *http.Request) (*url.URL, error) {
			return _url, err
		}
	}
}
