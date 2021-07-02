package httpx

import (
	"context"
	"net/http"
	"strings"

	"golang.org/x/text/encoding"
)

// RequestOption is defined for custom Http Request
type RequestOption func(r *http.Request) (error, *http.Request)

// SetQueries add query params to url. The params key/value will be encoded.
func SetQueries(params ...*Param) RequestOption {
	return SetQueriesWithEncoding(nil, params...)
}

// SetQueriesWithEncoding add query params to url. The params key/value will be encoded using specified encoding.
func SetQueriesWithEncoding(enc encoding.Encoding, params ...*Param) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		if len(params) == 0 {
			return nil, r
		}
		url := r.URL
		var sb strings.Builder
		if url.RawQuery != "" {
			sb.WriteString(url.RawQuery)
			sb.WriteByte('&')
		}
		err := EncodeParamsTo(sb, enc, params...)
		if err != nil {
			return err, nil
		}
		url.RawQuery = sb.String()
		return nil, r
	}
}

// AddHeader add a header to request.For host header, use SetHost.
func AddHeader(name, value string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		r.Header.Add(name, value)
		return nil, r
	}
}

// AddHeaders add headers to request.For host header, use SetHost.
func AddHeaders(headers map[string]string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		for name, value := range headers {
			r.Header.Add(name, value)
		}
		return nil, r
	}

}

// AddHeaders2 add headers to request.For host header, use SetHost.
func AddHeaders2(headers ...*Header) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		for _, h := range headers {
			r.Header.Add(h.Name, h.Value)
		}
		return nil, r
	}
}

// SetHeader set a header. For host header, use SetHost.
func SetHeader(name, value string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		r.Header.Set(name, value)
		return nil, r
	}
}

// SetHeaders set headers. For host header, use SetHost.
func SetHeaders(headers map[string]string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		for name, value := range headers {
			r.Header.Add(name, value)
		}
		return nil, r
	}
}

// SetHeaders2 set headers. For host header, use SetHost.
func SetHeaders2(headers ...*Header) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		for _, h := range headers {
			r.Header.Add(h.Name, h.Value)
		}
		return nil, r
	}
}

// AddCookie add a cookie to request
func AddCookie(name, value string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		r.AddCookie(&http.Cookie{Name: name, Value: value})
		return nil, r
	}
}

// AddCookies add cookies to request
func AddCookies(cookies map[string]string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		for name, value := range cookies {
			r.AddCookie(&http.Cookie{Name: name, Value: value})
		}
		return nil, r
	}
}

// AddCookies2 add cookies to request
func AddCookies2(firstCookie *Cookie, cookies ...*Cookie) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		r.AddCookie(firstCookie)
		for _, cookie := range cookies {
			r.AddCookie(cookie)
		}
		return nil, r
	}
}

// BasicAuth set request HTTP Basic Authentication with the provided username and password
func BasicAuth(user, password string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		r.SetBasicAuth(user, password)
		return nil, r
	}
}

// SetHost set Host header for request. The host header cannot set by SetHeader
func SetHost(host string) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		r.Host = host
		return nil, r
	}
}

// WithContext set the context for request
func WithContext(ctx context.Context) RequestOption {
	return func(r *http.Request) (error, *http.Request) {
		return nil, r.WithContext(ctx)
	}
}
