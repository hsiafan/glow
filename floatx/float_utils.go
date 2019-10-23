package floatx

import (
	"strconv"
)

// Format64 convert float64 to string
func Format64(value float64, decimals int) string {
	return strconv.FormatFloat(value, 'f', decimals, 64)
}

// Format32 convert float32 to string
func Format32(value float32, decimals int) string {
	return strconv.FormatFloat(float64(value), 'f', decimals, 32)
}

// Parse64 convert str to float64.
func Parse64(str string, defaultValue float64) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Parse32 convert str to float32.
func Parse32(str string, defaultValue float32) (float32, error) {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(value), err
}

// SafeParse64 convert str to float64. if str is not a illegal uint value representation, return defaultValue
func SafeParse64(str string, defaultValue float64) float64 {
	if value, err := Parse64(str, 64); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse32 convert str to float32. if str is not a illegal uint value representation, return defaultValue
func SafeParse32(str string, defaultValue float32) float32 {
	if value, err := Parse32(str, 32); err == nil {
		return value
	}
	return defaultValue
}
