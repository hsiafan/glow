package reflectx

// IsInt return if is int value
func IsInt(v interface{}) bool {
	if _, ok := v.(int); ok {
		return true
	}
	if _, ok := v.(int8); ok {
		return true
	}
	if _, ok := v.(int16); ok {
		return true
	}
	if _, ok := v.(int32); ok {
		return true
	}
	if _, ok := v.(int64); ok {
		return true
	}
	if _, ok := v.(uint); ok {
		return true
	}
	if _, ok := v.(uint8); ok {
		return true
	}
	if _, ok := v.(uint16); ok {
		return true
	}
	if _, ok := v.(uint32); ok {
		return true
	}
	if _, ok := v.(uint64); ok {
		return true
	}
	return false
}

// IsFloat return if is float value
func IsFloat(v interface{}) bool {
	if _, ok := v.(float32); ok {
		return true
	}
	if _, ok := v.(float64); ok {
		return true
	}
	return false
}
