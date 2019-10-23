package stringx

import (
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

// SplitPair split str into two part, by delimiter.
// If str dost not contains delimiter, the first returned str is original str, the second is empty str.
// If delimiter is empty, the first returned str is original str, the second is empty str.
func SplitPair(str string, delimiter string) (string, string) {
	if len(delimiter) == 0 {
		return str, ""
	}
	index := strings.Index(str, delimiter)
	if index == -1 {
		return str, ""
	}
	return str[:index], str[index+len(delimiter):]
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
