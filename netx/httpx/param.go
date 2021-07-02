package httpx

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"net/url"
	"strings"
)

// Param is a http get param key/value, or www-form-encoded post param.
type Param struct {
	Name  string
	Value string
}

// NewParam create new param
func NewParam(name string, value string) *Param {
	return &Param{Name: name, Value: value}
}

// ParseParam parse xx=xx to param. The white space char with '=' will be trimmed.
// If str do not contains '=', the str is treat as key. If contains multi '=', use the first one.
func ParseParam(str string) *Param {
	str = strings.TrimSpace(str)
	idx := strings.IndexByte(str, '=')
	if idx < 0 {
		return &Param{Name: str}
	}
	return &Param{Name: strings.TrimSpace(str[:idx]), Value: strings.TrimSpace(str[idx+1:])}
}

// EncodeParams encode params to encoded str.
func EncodeParams(enc encoding.Encoding, params ...*Param) (string, error) {
	var sb strings.Builder
	if err := EncodeParamsTo(sb, enc, params...); err != nil {
		return "", err
	}
	return sb.String(), nil
}

// EncodeParamsTo encode params, and write to buf.
func EncodeParamsTo(buf strings.Builder, enc encoding.Encoding, params ...*Param) error {
	for _, param := range params {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		name, err := EncodeQuery(param.Name, enc)
		if err != nil {
			return err
		}
		buf.WriteString(url.QueryEscape(name))
		buf.WriteByte('=')
		value, err := EncodeQuery(param.Value, enc)
		if err != nil {
			return err
		}
		buf.WriteString(url.QueryEscape(value))
	}
	return nil
}

// EncodeQuery encode a single str token as http param key or value.
func EncodeQuery(str string, enc encoding.Encoding) (string, error) {
	if enc == nil || enc == unicode.UTF8 {
		return url.QueryEscape(str), nil
	}
	str, err := enc.NewEncoder().String(str)
	if err != nil {
		return "", err
	}
	return url.QueryEscape(str), nil
}

// DecodeQuery decode a single str token to http param key or value.
func DecodeQuery(str string, enc encoding.Encoding) (string, error) {
	s, err := url.QueryUnescape(str)
	if err != nil {
		return "", err
	}
	if enc == nil || enc == unicode.UTF8 {
		return s, nil
	}
	return enc.NewDecoder().String(s)
}
