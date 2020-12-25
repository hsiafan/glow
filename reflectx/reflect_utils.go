package reflectx

import (
	"fmt"
	"reflect"
)

// IsInt return if is int value
func IsInt(v interface{}) bool {
	switch v.(type) {
	case int:
		return true
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	default:
		return false
	}
}

// IsFloat return if is float value
func IsFloat(v interface{}) bool {
	switch v.(type) {
	case float32:
		return true
	case float64:
		return true
	default:
		return false
	}
}

// StructToMap convert struct to map, using field name as key, field value as map value.
// Only exported fields are set into map.
func StructToMap(value interface{}) map[string]interface{} {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return toMap(v.Elem())
	case reflect.Struct:
		return toMap(v)
	default:
		panic(fmt.Sprintf("%v is not struct", v.Kind()))
	}
}

func toMap(v reflect.Value) map[string]interface{} {
	if v.Kind() != reflect.Struct {
		panic(fmt.Sprintf("%v is not struct", v.Kind()))
	}
	m := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanInterface() {
			m[v.Type().Field(i).Name] = f.Interface()
		}
	}
	return m
}
