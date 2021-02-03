package intx

import "strconv"

const MaxUint = ^uint(0)
const MinUint = 0

// FormatUint convert uint to str
func FormatUint(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

// FormatUint64 convert uint64 to str
func FormatUint64(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// FormatUint32 convert uint32 to str
func FormatUint32(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

// FormatUint16 convert uint16 to str
func FormatUint16(v uint16) string {
	return strconv.FormatUint(uint64(v), 10)
}

// FormatUint8 convert uint8 to str
func FormatUint8(v uint8) string {
	return strconv.FormatUint(uint64(v), 10)
}

// ParseUint64 convert str to uint64. if str is not a illegal uint value representation, return defaultValue
func ParseUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// ParseUint32 convert str to uint32. if str is not a illegal uint value representation, return defaultValue
func ParseUint32(str string) (uint32, error) {
	v, err := strconv.ParseUint(str, 10, 32)
	return uint32(v), err
}

// ParseUint16 convert str to uint16. if str is not a illegal uint value representation, return defaultValue
func ParseUint16(str string) (uint16, error) {
	v, err := strconv.ParseUint(str, 10, 16)
	return uint16(v), err
}

// ParseUint8 convert str to uint8. if str is not a illegal uint value representation, return defaultValue
func ParseUint8(str string) (uint8, error) {
	v, err := strconv.ParseUint(str, 10, 8)
	return uint8(v), err
}

// ParseUint convert str to uint. if str is not a illegal uint value representation, return defaultValue
func ParseUint(str string) (uint, error) {
	v, err := strconv.ParseUint(str, 10, 0)
	return uint(v), err
}

// SafeParseUint64 convert str to uint64. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint64(str string, defaultValue uint64) uint64 {
	if value, err := ParseUint64(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUint32 convert str to uint32. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint32(str string, defaultValue uint32) uint32 {
	if value, err := ParseUint32(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUint16 convert str to uint16. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint16(str string, defaultValue uint16) uint16 {
	if value, err := ParseUint16(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUint8 convert str to uint8. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint8(str string, defaultValue uint8) uint8 {
	if value, err := ParseUint8(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUint convert str to uint. if str is not a illegal uint value representation, return defaultValue
func SafeParseUint(str string, defaultValue uint) uint {
	if value, err := ParseUint(str); err == nil {
		return value
	}
	return defaultValue
}

// MinU64 return smaller one of uint64 values
func MinU64(value1, value2 uint64) uint64 {
	if value1 < value2 {
		return value1
	}
	return value2
}

// MaxU64 return bigger one of uint64 values
func MaxU64(value1, value2 uint64) uint64 {
	if value1 > value2 {
		return value1
	}
	return value2
}

// MinU32 return smaller one of uint32 values
func MinU32(value1, value2 uint32) uint32 {
	if value1 < value2 {
		return value1
	}
	return value2
}

// MaxU32 return bigger one of uint32 values
func MaxU32(value1, value2 uint32) uint32 {
	if value1 > value2 {
		return value1
	}
	return value2
}

// MinU return smaller one of int values
func MinU(value1, value2 uint) uint {
	if value1 < value2 {
		return value1
	}
	return value2
}

// MaxU return bigger one of uint values
func MaxU(value1, value2 uint) uint {
	if value1 > value2 {
		return value1
	}
	return value2
}
