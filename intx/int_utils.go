package intx

import "strconv"

// Format convert int to str
func Format(v int) string {
	return strconv.FormatInt(int64(v), 10)
}

// Format64 convert int64 to str
func Format64(v int64) string {
	return strconv.FormatInt(v, 10)
}

// Format32 convert int32 to str
func Format32(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

// Format16 convert int16 to str
func Format16(v int16) string {
	return strconv.FormatInt(int64(v), 10)
}

// Format8 convert int8 to str
func Format8(v int8) string {
	return strconv.FormatInt(int64(v), 10)
}

// Parse convert str to int
func Parse(str string) (int, error) {
	value, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(value), err
}

// Parse64 convert str to int64
func Parse64(str string) (int64, error) {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(value), err
}

// Parse32 convert str to int32
func Parse32(str string) (int32, error) {
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), err
}

// Parse16 convert str to int16
func Parse16(str string) (int16, error) {
	value, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(value), err
}

// Parse8 convert str to int8
func Parse8(str string) (int8, error) {
	value, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(value), err
}

// SafeParse convert str to int. if str is not a illegal int value representation, return defaultValue
func SafeParse(str string, defaultValue int) int {
	if value, err := Parse(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse64 convert str to int64. if str is not a illegal int value representation, return defaultValue
func SafeParse64(str string, defaultValue int64) int64 {
	if value, err := Parse64(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse32 convert str to int32. if str is not a illegal int value representation, return defaultValue
func SafeParse32(str string, defaultValue int32) int32 {
	if value, err := Parse32(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse16 convert str to int16. if str is not a illegal int value representation, return defaultValue
func SafeParse16(str string, defaultValue int16) int16 {
	if value, err := Parse16(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse8 convert str to int8. if str is not a illegal int value representation, return defaultValue
func SafeParse8(str string, defaultValue int8) int8 {
	if value, err := Parse8(str); err == nil {
		return value
	}
	return defaultValue
}
