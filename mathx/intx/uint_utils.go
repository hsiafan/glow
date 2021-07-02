package intx

import "strconv"

const MaxUint = ^uint(0)
const MinUint = 0

// FormatUnsigned convert uint to str
func FormatUnsigned(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

// FormatUnsigned64 convert uint64 to str
func FormatUnsigned64(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// FormatUnsigned32 convert uint32 to str
func FormatUnsigned32(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

// FormatUnsigned16 convert uint16 to str
func FormatUnsigned16(v uint16) string {
	return strconv.FormatUint(uint64(v), 10)
}

// FormatUnsigned8 convert uint8 to str
func FormatUnsigned8(v uint8) string {
	return strconv.FormatUint(uint64(v), 10)
}

// ParseUnsigned64 convert str to uint64. if str is not a illegal uint value representation, return defaultValue
func ParseUnsigned64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// ParseUnsigned32 convert str to uint32. if str is not a illegal uint value representation, return defaultValue
func ParseUnsigned32(str string) (uint32, error) {
	v, err := strconv.ParseUint(str, 10, 32)
	return uint32(v), err
}

// ParseUnsigned16 convert str to uint16. if str is not a illegal uint value representation, return defaultValue
func ParseUnsigned16(str string) (uint16, error) {
	v, err := strconv.ParseUint(str, 10, 16)
	return uint16(v), err
}

// ParseUnsigned8 convert str to uint8. if str is not a illegal uint value representation, return defaultValue
func ParseUnsigned8(str string) (uint8, error) {
	v, err := strconv.ParseUint(str, 10, 8)
	return uint8(v), err
}

// ParseUnsigned convert str to uint. if str is not a illegal uint value representation, return defaultValue
func ParseUnsigned(str string) (uint, error) {
	v, err := strconv.ParseUint(str, 10, 0)
	return uint(v), err
}

// SafeParseUnsigned64 convert str to uint64. if str is not a illegal uint value representation, return defaultValue
func SafeParseUnsigned64(str string, defaultValue uint64) uint64 {
	if value, err := ParseUnsigned64(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUnsigned32 convert str to uint32. if str is not a illegal uint value representation, return defaultValue
func SafeParseUnsigned32(str string, defaultValue uint32) uint32 {
	if value, err := ParseUnsigned32(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUnsigned16 convert str to uint16. if str is not a illegal uint value representation, return defaultValue
func SafeParseUnsigned16(str string, defaultValue uint16) uint16 {
	if value, err := ParseUnsigned16(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUnsigned8 convert str to uint8. if str is not a illegal uint value representation, return defaultValue
func SafeParseUnsigned8(str string, defaultValue uint8) uint8 {
	if value, err := ParseUnsigned8(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParseUnsigned convert str to uint. if str is not a illegal uint value representation, return defaultValue
func SafeParseUnsigned(str string, defaultValue uint) uint {
	if value, err := ParseUnsigned(str); err == nil {
		return value
	}
	return defaultValue
}

// MinUnsigned64 return smaller one of uint64 values
func MinUnsigned64(value1, value2 uint64) uint64 {
	if value1 < value2 {
		return value1
	}
	return value2
}

// MaxUnsigned64 return bigger one of uint64 values
func MaxUnsigned64(value1, value2 uint64) uint64 {
	if value1 > value2 {
		return value1
	}
	return value2
}

// MinUnsigned32 return smaller one of uint32 values
func MinUnsigned32(value1, value2 uint32) uint32 {
	if value1 < value2 {
		return value1
	}
	return value2
}

// MaxUnsigned32 return bigger one of uint32 values
func MaxUnsigned32(value1, value2 uint32) uint32 {
	if value1 > value2 {
		return value1
	}
	return value2
}

// MinUnsigned16 return smaller one of uint16 values
func MinUnsigned16(value1, value2 uint16) uint16 {
	if value1 < value2 {
		return value1
	}
	return value2
}

// MaxUnsigned16 return bigger one of uint16 values
func MaxUnsigned16(value1, value2 uint16) uint16 {
	if value1 > value2 {
		return value1
	}
	return value2
}

// MinUnsigned return smaller one of int values
func MinUnsigned(value1, value2 uint) uint {
	if value1 < value2 {
		return value1
	}
	return value2
}

// MaxUnsigned return bigger one of uint values
func MaxUnsigned(value1, value2 uint) uint {
	if value1 > value2 {
		return value1
	}
	return value2
}
