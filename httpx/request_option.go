package httpx

import (
	"net/http"
	"strings"

	"golang.org/x/text/encoding"
)

// For custom Http Request
type RequestOption func(r *http.Request) error

func Queries(params ...*Param) RequestOption {
	return QueriesWithEncoding(params, nil)
}

func QueriesWithEncoding(params []*Param, enc encoding.Encoding) RequestOption {
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
		err := EncodeParamsTo(params, enc, sb)
		if err != nil {
			return err
		}
		url.RawQuery = sb.String()
		return nil
	}
}

func AddHeader(name, value string) RequestOption {
	return func(r *http.Request) error {
		r.Header.Add(name, value)
		return nil
	}
}

func AddHeaders(headers ...Param) RequestOption {
	return func(r *http.Request) error {
		for _, param := range headers {
			r.Header.Add(param.Name, param.Value)
		}
		return nil
	}

}

func SetHeader(name, value string) RequestOption {
	return func(r *http.Request) error {
		r.Header.Set(name, value)
		return nil
	}
}

func SetHeaders(headers map[string]string) RequestOption {
	return func(r *http.Request) error {
		for name, value := range headers {
			r.Header.Add(name, value)
		}
		return nil
	}
}

func AddCookie(name, value string) RequestOption {
	return func(r *http.Request) error {
		r.AddCookie(&http.Cookie{Name: name, Value: value})
		return nil
	}
}

func AddCookies(cookies ...Param) RequestOption {
	return func(r *http.Request) error {
		for _, cookie := range cookies {
			r.AddCookie(&http.Cookie{Name: cookie.Name, Value: cookie.Value})
		}
		return nil
	}
}

// Set request HTTP Basic Authentication with the provided username and password
func BasicAuth(user, password string) RequestOption {
	return func(r *http.Request) error {
		r.SetBasicAuth(user, password)
		return nil
	}
}

// set Host header for reqeust. The host header cannot set by SetHeader
func SetHost(host string) RequestOption {
	return func(r *http.Request) error {
		r.Host = host
		return nil
	}
}
