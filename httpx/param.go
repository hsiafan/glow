package httpx

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"net/url"
	"strings"
)

// Param
type Param struct {
	Name  string
	Value string
}

// Create new param
func NewParam(name string, value string) *Param {
	return &Param{Name: name, Value: value}
}

// parse xx=xx to param
func ParseParam(str string) *Param {
	str = strings.TrimSpace(str)
	idx := strings.IndexByte(str, '=')
	if idx < 0 {
		return &Param{Name: str}
	}
	return &Param{Name: strings.TrimSpace(str[:idx]), Value: strings.TrimSpace(str[idx+1:])}
}

func EncodeParams(params []*Param, enc encoding.Encoding) (string, error) {
	var sb strings.Builder
	if err := EncodeParamsTo(params, enc, sb); err != nil {
		return "", err
	}
	return sb.String(), nil
}

func EncodeParamsTo(params []*Param, enc encoding.Encoding, buf strings.Builder) error {
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
