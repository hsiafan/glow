package stringx

import (
	"fmt"
	"github.com/hsiafan/glow/floatx"
	"github.com/hsiafan/glow/intx"
	"github.com/hsiafan/glow/stringx/ascii"
	"github.com/hsiafan/glow/unsafex"
	"strconv"
	"strings"
	"unicode"
)

const (
	nilString = "<nil>"
)

// ValueOf return string representation for value
func ValueOf(v interface{}) string {
	if v == nil {
		return nilString
	}

	switch v.(type) {
	case string:
		return v.(string)
	case bool:
		return strconv.FormatBool(v.(bool))
	case int:
		return intx.FormatInt(v.(int))
	case int8:
		return intx.FormatInt8(v.(int8))
	case int16:
		return intx.FormatInt16(v.(int16))
	case int32:
		return intx.FormatInt32(v.(int32))
	case int64:
		return intx.FormatInt64(v.(int64))
	case uint:
		return intx.FormatUint(v.(uint))
	case uint8:
		return intx.FormatUint8(v.(uint8))
	case uint16:
		return intx.FormatUint16(v.(uint16))
	case uint32:
		return intx.FormatUint32(v.(uint32))
	case uint64:
		return intx.FormatUint64(v.(uint64))
	case float32:
		return floatx.SimpleFormat32(v.(float32))
	case float64:
		return floatx.SimpleFormat64(v.(float64))
	case complex64:
		c := v.(complex64)
		return strconv.FormatComplex(complex(float64(real(c)), float64(imag(c))), 'f', -1, 64)
	case complex128:
		return strconv.FormatComplex(v.(complex128), 'f', -1, 128)
	case fmt.Stringer:
		return v.(fmt.Stringer).String()
	case error:
		return v.(error).Error()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// AppendIfMissing return a str end with suffix appended if not has the suffix; otherwise return str it's self
func AppendIfMissing(str string, suffix string) string {
	if !strings.HasSuffix(str, suffix) {
		return str + suffix
	}
	return str
}

// PrependIfMissing return a str start with suffix appended if not has the prefix; otherwise return str it's self
func PrependIfMissing(str string, prefix string) string {
	if !strings.HasSuffix(str, prefix) {
		return prefix + str
	}
	return str
}

// AppendIfNotEmpty return str with suffix if str is not empty; return the origin str otherwise.
func AppendIfNotEmpty(str string, suffix string) string {
	if str == "" {
		return str
	}
	return str + suffix
}

// PrependIfNotEmpty return str with prefix if str is not empty; return the origin str otherwise.
func PrependIfNotEmpty(str string, prefix string) string {
	if str == "" {
		return str
	}
	return prefix + str
}

// SubstringAfter return sub string after the sep. If str does not contains sep, return empty str.
func SubstringAfter(str string, sep string) string {
	index := strings.Index(str, sep)
	if index == -1 {
		return ""
	}
	return str[index+len(sep):]
}

// SubstringAfterLast return sub string after the last sep. If str does not contains sep, return empty str.
func SubstringAfterLast(str string, sep string) string {
	index := strings.LastIndex(str, sep)
	if index == -1 {
		return ""
	}
	return str[index+len(sep):]
}

// SubstringBefore return sub string after the sep. If str does not contains sep, return the original str.
func SubstringBefore(str string, sep string) string {
	index := strings.Index(str, sep)
	if index == -1 {
		return str
	}
	return str[:index]
}

// SubstringBeforeLast return sub string after the last sep. If str does not contains sep, return the original str.
func SubstringBeforeLast(str string, sep string) string {
	index := strings.LastIndex(str, sep)
	if index == -1 {
		return str
	}
	return str[:index]
}

// PadLeft pad str to width, with padding rune at left.
// If str len already equals with or larger than width, return original str.
func PadLeft(str string, width int, r rune) string {
	if len(str) >= width {
		return str
	}
	var builder strings.Builder
	builder.Grow(width)
	padded := width - len(str)
	for i := 0; i < padded; i++ {
		builder.WriteRune(r)
	}
	builder.WriteString(str)
	return builder.String()
}

// PadLeft pad str to width, with padding rune at right.
// If str len already equals with or larger than width, return original str.
func PadRight(str string, width int, r rune) string {
	if len(str) >= width {
		return str
	}
	var builder strings.Builder
	builder.Grow(width)
	padded := width - len(str)
	builder.WriteString(str)
	for i := 0; i < padded; i++ {
		builder.WriteRune(r)
	}
	return builder.String()
}

// PadToCenter pad str to width, with padding rune at left and right.
// If str len already equals with or larger than width, return original str.
func PadToCenter(str string, width int, r rune) string {
	if len(str) >= width {
		return str
	}
	var builder strings.Builder
	builder.Grow(width)
	padded := width - len(str)
	for i := 0; i < padded/2; i++ {
		builder.WriteRune(r)
	}
	builder.WriteString(str)
	for i := 0; i < padded-padded/2; i++ {
		builder.WriteRune(r)
	}
	return builder.String()
}

// Capitalize return str with first char of ascii str upper case.
func Capitalize(str string) string {
	if str == "" {
		return str
	}
	if ascii.IsUpper(str[0]) {
		return str
	}
	bytes := []byte(str)
	bytes[0] = ascii.ToUpper(str[0])
	return unsafex.BytesToString(bytes)
}

// DeCapitalize return str with first char of ascii str lower case.
func DeCapitalize(str string) string {
	if str == "" {
		return str
	}
	if ascii.IsLower(str[0]) {
		return str
	}
	bytes := []byte(str)
	bytes[0] = ascii.ToLower(str[0])
	return unsafex.BytesToString(bytes)
}

// SnakeToCamel convert underscore style str to Camel.
// The param capitalized determine if first char is uppercase.
func SnakeToCamel(s string, capitalized bool) string {
	var sb strings.Builder
	var beginNewWord = false
	var firstChar = true
	for _, c := range s {
		if c == '_' {
			beginNewWord = true
			continue
		}
		if firstChar {
			if capitalized {
				sb.WriteRune(unicode.ToUpper(c))
			} else {
				sb.WriteRune(c)
			}
			firstChar = false
			beginNewWord = false
			continue
		}
		if !beginNewWord {
			sb.WriteRune(c)
		} else {
			beginNewWord = false
			sb.WriteRune(unicode.ToUpper(c))
		}
	}
	return sb.String()
}

// CamelToSnake convert camel style str to underscore snake style.
func CamelToSnake(s string) string {
	runes := []rune(s)
	var sb strings.Builder
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if unicode.IsUpper(c) {
			if i > 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(c))
			j := i + 1
			for ; j < len(runes); j++ {
				nc := runes[j]
				if !unicode.IsUpper(nc) {
					j--
					break
				}
			}
			j--
			if j > i {
				for k := i + 1; k <= j; k++ {
					sb.WriteRune(unicode.ToLower(runes[k]))
				}
				i = j
			}
		} else {
			sb.WriteRune(c)
		}
	}

	return sb.String()
}

// Copy copy a string content, for reducing large string content memory usage when do substring.
// This method allocate a new string content byte array, thereby allow the larger string to be released by the garbage collector once it is no longer referenced
func Copy(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	sb.WriteString(s)
	return sb.String()
}
