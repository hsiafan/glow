package intx

import "strconv"

const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// FormatInt convert int to str
func FormatInt(v int) string {
	return strconv.FormatInt(int64(v), 10)
}

// FormatInt64 convert int64 to str
func FormatInt64(v int64) string {
	return strconv.FormatInt(v, 10)
}

// FormatInt32 convert int32 to str
func FormatInt32(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

// FormatInt16 convert int16 to str
func FormatInt16(v int16) string {
	return strconv.FormatInt(int64(v), 10)
}

// FormatInt8 convert int8 to str
func FormatInt8(v int8) string {
	return strconv.FormatInt(int64(v), 10)
}

// ParseInt convert str to int
func ParseInt(str string) (int, error) {
	value, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(value), err
}

// ParseInt64 convert str to int64
func ParseInt64(str string) (int64, error) {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(value), err
}

// ParseInt32 convert str to int32
func ParseInt32(str string) (int32, error) {
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), err
}

// ParseInt16 convert str to int16
func ParseInt16(str string) (int16, error) {
	value, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(value), err
}

// ParseInt8 convert str to int8
func ParseInt8(str string) (int8, error) {
	value, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(value), err
}

// SafeParseInt convert str to int. if str is not a illegal int value representation, return defaultValue
func SafeParseInt(str string, defaultValue int) int {
	if value, err := ParseInt(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseInt64 convert str to int64. if str is not a illegal int value representation, return defaultValue
func SafeParseInt64(str string, defaultValue int64) int64 {
	if value, err := ParseInt64(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseInt32 convert str to int32. if str is not a illegal int value representation, return defaultValue
func SafeParseInt32(str string, defaultValue int32) int32 {
	if value, err := ParseInt32(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseInt16 convert str to int16. if str is not a illegal int value representation, return defaultValue
func SafeParseInt16(str string, defaultValue int16) int16 {
	if value, err := ParseInt16(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseInt8 convert str to int8. if str is not a illegal int value representation, return defaultValue
func SafeParseInt8(str string, defaultValue int8) int8 {
	if value, err := ParseInt8(str); err == nil {
		return value
	}
	return defaultValue
}
