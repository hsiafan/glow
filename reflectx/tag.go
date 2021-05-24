package reflectx

import (
	"github.com/hsiafan/glow/floatx"
	"github.com/hsiafan/glow/intx"
	"reflect"
	"strconv"
)

// Get tag value with name. if not present, return defaultValue
func GetTagValue(tag reflect.StructTag, name string, defaultValue string) string {
	if value, ok := tag.Lookup(name); ok {
		return value
	}
	return defaultValue
}

// Get tag value as int with name. if not present, return defaultValue.
// If tag present and parse error, return error.
func GetIntTagValue(tag reflect.StructTag, name string, defaultValue int) (int, error) {
	if value, ok := tag.Lookup(name); ok {
		return strconv.Atoi(value)
	}
	return defaultValue, nil
}

// Get tag value as int64 with name. if not present, return defaultValue.
// If tag present and parse error, return error.
func GetInt64TagValue(tag reflect.StructTag, name string, defaultValue int64) (int64, error) {
	if value, ok := tag.Lookup(name); ok {
		return strconv.ParseInt(value, 10, 64)
	}
	return defaultValue, nil
}

// Get tag value as uint with name. if not present, return defaultValue.
// If tag present and parse error, return error.
func GetUIntTagValue(tag reflect.StructTag, name string, defaultValue uint) (uint, error) {
	if value, ok := tag.Lookup(name); ok {
		return intx.ParseUnsigned(value)
	}
	return defaultValue, nil
}

// Get tag value as uint64 with name. if not present, return defaultValue.
// If tag present and parse error, return error.
func GetUInt64TagValue(tag reflect.StructTag, name string, defaultValue uint64) (uint64, error) {
	if value, ok := tag.Lookup(name); ok {
		return strconv.ParseUint(value, 10, 64)
	}
	return defaultValue, nil
}

// Get tag value as boolean with name. if not present, return defaultValue.
// If tag present and parse error, return error.
func GetBoolTagValue(tag reflect.StructTag, name string, defaultValue bool) (bool, error) {
	if value, ok := tag.Lookup(name); ok {
		return strconv.ParseBool(value)
	}
	return defaultValue, nil
}

// Get tag value as float64 with name. if not present, return defaultValue.
// If tag present and parse error, return error.
func GetFloat64TagValue(tag reflect.StructTag, name string, defaultValue float64) (float64, error) {
	if value, ok := tag.Lookup(name); ok {
		return floatx.Parse64(value)
	}
	return defaultValue, nil
}
