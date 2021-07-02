package reflectx

import (
	"fmt"
	"reflect"
)

// StructToMap convert struct to map, using field name as key, field value as map value.
// The Field name can be override using 'name' struct tag.
//
// Note: only exported fields are set into map.
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
			sf := v.Type().Field(i)
			name := sf.Tag.Get("name")
			if name == "" {
				name = sf.Name
			}
			m[name] = f.Interface()
		}
	}
	return m
}
