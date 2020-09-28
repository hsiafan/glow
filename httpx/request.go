package httpx

import (
	"net/http"
	"strings"

	"golang.org/x/text/encoding"
)

// RequestOption is defined for custom Http Request
type RequestOption func(r *http.Request) error

// SetQueries add query params to url. The params key/value will be encoded.
func SetQueries(params ...*Param) RequestOption {
	return SetQueriesWithEncoding(nil, params...)
}

// SetQueriesWithEncoding add query params to url. The params key/value will be encoded using specified encoding.
func SetQueriesWithEncoding(enc encoding.Encoding, params ...*Param) RequestOption {
	return func(r *http.Request) error {
		if len(params) == 0 {
			return nil
		}
		url := r.URL
		var sb strings.Builder
		if url.RawQuery != "" {
			sb.WriteString(url.RawQuery)
			sb.WriteByte('&')
		}
		err := EncodeParamsTo(sb, enc, params...)
		if err != nil {
			return err
		}
		url.RawQuery = sb.String()
		return nil
	}
}

// AddHeader add a header to request.For host header, use SetHost.
func AddHeader(name, value string) RequestOption {
	return func(r *http.Request) error {
		r.Header.Add(name, value)
		return nil
	}
}

// AddHeaders add headers to request.For host header, use SetHost.
func AddHeaders(headers map[string]string) RequestOption {
	return func(r *http.Request) error {
		for name, value := range headers {
			r.Header.Add(name, value)
		}
		return nil
	}

}

// AddHeaders2 add headers to request.For host header, use SetHost.
func AddHeaders2(headers ...*Header) RequestOption {
	return func(r *http.Request) error {
		for _, h := range headers {
			r.Header.Add(h.Name, h.Value)
		}
		return nil
	}
}

// SetHeader set a header. For host header, use SetHost.
func SetHeader(name, value string) RequestOption {
	return func(r *http.Request) error {
		r.Header.Set(name, value)
		return nil
	}
}

// SetHeaders set headers. For host header, use SetHost.
func SetHeaders(headers map[string]string) RequestOption {
	return func(r *http.Request) error {
		for name, value := range headers {
			r.Header.Add(name, value)
		}
		return nil
	}
}

// SetHeaders2 set headers. For host header, use SetHost.
func SetHeaders2(headers ...*Header) RequestOption {
	return func(r *http.Request) error {
		for _, h := range headers {
			r.Header.Add(h.Name, h.Value)
		}
		return nil
	}
}

// AddCookie add a cookie to request
func AddCookie(name, value string) RequestOption {
	return func(r *http.Request) error {
		r.AddCookie(&http.Cookie{Name: name, Value: value})
		return nil
	}
}

// AddCookies add cookies to request
func AddCookies(cookies map[string]string) RequestOption {
	return func(r *http.Request) error {
		for name, value := range cookies {
			r.AddCookie(&http.Cookie{Name: name, Value: value})
		}
		return nil
	}
}

// AddCookies2 add cookies to request
func AddCookies2(firstCookie *Cookie, cookies ...*Cookie) RequestOption {
	return func(r *http.Request) error {
		r.AddCookie(firstCookie)
		for _, cookie := range cookies {
			r.AddCookie(cookie)
		}
		return nil
	}
}

// BasicAuth set request HTTP Basic Authentication with the provided username and password
func BasicAuth(user, password string) RequestOption {
	return func(r *http.Request) error {
		r.SetBasicAuth(user, password)
		return nil
	}
}

// SetHost set Host header for request. The host header cannot set by SetHeader
func SetHost(host string) RequestOption {
	return func(r *http.Request) error {
		r.Host = host
		return nil
	}
}
