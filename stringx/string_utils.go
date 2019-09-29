package stringx

import (
	"strconv"
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

// SafeParseInt convert str to int. if str is not a illegal int value representation, return defaultValue
func SafeParseInt(str string, defaultValue int) int {
	if value, err := strconv.Atoi(str); err != nil {
		return value
	}
	return defaultValue
}

// SafeParseInt64 convert str to int64. if str is not a illegal int value representation, return defaultValue
func SafeParseInt64(str string, defaultValue int64) int64 {
	if value, err := strconv.ParseInt(str, 10, 64); err != nil {
		return value
	}
	return defaultValue
}

// SafeParseInt32 convert str to int32. if str is not a illegal int value representation, return defaultValue
func SafeParseInt32(str string, defaultValue int32) int32 {
	if value, err := strconv.ParseInt(str, 10, 32); err != nil {
		return int32(value)
	}
	return defaultValue
}

// SafeParseInt16 convert str to int16. if str is not a illegal int value representation, return defaultValue
func SafeParseInt16(str string, defaultValue int16) int16 {
	if value, err := strconv.ParseInt(str, 10, 16); err != nil {
		return int16(value)
	}
	return defaultValue
}

// SafeParseInt8 convert str to int8. if str is not a illegal int value representation, return defaultValue
func SafeParseInt8(str string, defaultValue int8) int8 {
	if value, err := strconv.ParseInt(str, 10, 8); err != nil {
		return int8(value)
	}
	return defaultValue
}

// SafeParseInt convert str to uint. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint(str string, defaultValue uint) uint {
	if value, err := strconv.ParseUint(str, 10, 0); err != nil {
		return uint(value)
	}
	return defaultValue
}

// SafeParseUint64 convert str to uint64. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint64(str string, defaultValue uint64) uint64 {
	if value, err := strconv.ParseUint(str, 10, 64); err != nil {
		return value
	}
	return defaultValue
}

// SafeParseUint32 convert str to uint32. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint32(str string, defaultValue uint32) uint32 {
	if value, err := strconv.ParseUint(str, 10, 32); err != nil {
		return uint32(value)
	}
	return defaultValue
}

// SafeParseUint16 convert str to uint16. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint16(str string, defaultValue uint16) uint16 {
	if value, err := strconv.ParseUint(str, 10, 16); err != nil {
		return uint16(value)
	}
	return defaultValue
}

// SafeParseUint8 convert str to uint8. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint8(str string, defaultValue uint8) uint8 {
	if value, err := strconv.ParseUint(str, 10, 8); err != nil {
		return uint8(value)
	}
	return defaultValue
}

// SafeParseFloat64 convert str to float64. if str is not a illegal uint value representation, return defaultValue
func SafeParseFloat64(str string, defaultValue float64) float64 {
	if value, err := strconv.ParseFloat(str, 64); err != nil {
		return value
	}
	return defaultValue
}

// SafeParseFloat32 convert str to float32. if str is not a illegal uint value representation, return defaultValue
func SafeParseFloat32(str string, defaultValue float32) float32 {
	if value, err := strconv.ParseFloat(str, 32); err != nil {
		return float32(value)
	}
	return defaultValue
}
