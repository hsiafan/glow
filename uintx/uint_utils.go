package uintx

import "strconv"

// Format convert uint to str
func Format(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Format64 convert uint64 to str
func Format64(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// Format32 convert uint32 to str
func Format32(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Format16 convert uint16 to str
func Format16(v uint16) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Format8 convert uint8 to str
func Format8(v uint8) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Parse64 convert str to uint64. if str is not a illegal uint value representation, return defaultValue
func Parse64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// Parse32 convert str to uint32. if str is not a illegal uint value representation, return defaultValue
func Parse32(str string) (uint32, error) {
	v, err := strconv.ParseUint(str, 10, 32)
	return uint32(v), err
}

// Parse16 convert str to uint16. if str is not a illegal uint value representation, return defaultValue
func Parse16(str string) (uint16, error) {
	v, err := strconv.ParseUint(str, 10, 16)
	return uint16(v), err
}

// Parse8 convert str to uint8. if str is not a illegal uint value representation, return defaultValue
func Parse8(str string) (uint8, error) {
	v, err := strconv.ParseUint(str, 10, 8)
	return uint8(v), err
}

// Parse convert str to uint. if str is not a illegal uint value representation, return defaultValue
func Parse(str string) (uint, error) {
	v, err := strconv.ParseUint(str, 10, 0)
	return uint(v), err
}

// SafeParse64 convert str to uint64. if str is not a illegal uint value representation, return defaultValue
func SafeParse64(str string, defaultValue uint64) uint64 {
	if value, err := Parse64(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse32 convert str to uint32. if str is not a illegal uint value representation, return defaultValue
func SafeParse32(str string, defaultValue uint32) uint32 {
	if value, err := Parse32(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse16 convert str to uint16. if str is not a illegal uint value representation, return defaultValue
func SafeParse16(str string, defaultValue uint16) uint16 {
	if value, err := Parse16(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse8 convert str to uint8. if str is not a illegal uint value representation, return defaultValue
func SafeParse8(str string, defaultValue uint8) uint8 {
	if value, err := Parse8(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse convert str to uint. if str is not a illegal uint value representation, return defaultValue
func SafeParse(str string, defaultValue uint) uint {
	if value, err := Parse(str); err == nil {
		return value
	}
	return defaultValue
}
