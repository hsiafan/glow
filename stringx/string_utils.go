package stringx

import (
	"github.com/hsiafan/glow/stringx/ascii"
	"github.com/hsiafan/glow/unsafex"
	"strings"
)

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
	var builder Builder
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
	var builder Builder
	builder.Grow(width)
	padded := width - len(str)
	builder.WriteString(str)
	for i := 0; i < padded; i++ {
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
