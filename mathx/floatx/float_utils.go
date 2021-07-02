package floatx

import (
	"math"
	"strconv"
)

// Format64 convert float64 to string, with decimals numbers.
func Format64(value float64, decimals int) string {
	return strconv.FormatFloat(value, 'f', decimals, 64)
}

// Format32 convert float32 to string, with decimals numbers.
func Format32(value float32, decimals int) string {
	return strconv.FormatFloat(float64(value), 'f', decimals, 32)
}

// SimpleFormat64 convert float64 to string
func SimpleFormat64(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// SimpleFormat32 convert float32 to string
func SimpleFormat32(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', -1, 32)
}

// Parse64 convert str to float64.
func Parse64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Parse32 convert str to float32.
func Parse32(str string) (float32, error) {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(value), err
}

// SafeParse64 convert str to float64. if str is not a illegal uint value representation, return defaultValue
func SafeParse64(str string, defaultValue float64) float64 {
	if value, err := Parse64(str); err == nil {
		return value
	}
	return defaultValue
}

// SafeParse32 convert str to float32. if str is not a illegal uint value representation, return defaultValue
func SafeParse32(str string, defaultValue float32) float32 {
	if value, err := Parse32(str); err == nil {
		return value
	}
	return defaultValue
}

// Max64 returns the larger one of float64 values
func Max64(value1, value2 float64) float64 {
	return math.Max(value1, value2)
}

// Min64 returns the smaller one of float64 values
func Min64(value1, value2 float64) float64 {
	return math.Min(value1, value2)
}

// Max32 return the larger one of float values
func Max32(value1, value2 float32) float32 {
	return float32(math.Max(float64(value1), float64(value2)))
}

// Min32 return the smaller one of float values
func Min32(value1, value2 float32) float32 {
	return float32(math.Min(float64(value1), float64(value2)))
}
