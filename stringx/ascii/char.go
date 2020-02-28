package ascii

// If ascii char is upper case
func IsUpper(v byte) bool {
	return 'A' <= v && v <= 'Z'
}

// If ascii char is lower case
func IsLower(v byte) bool {
	return 'a' <= v && v <= 'z'
}

// convert ascii char to upper case
func ToUpper(v byte) byte {
	if IsLower(v) {
		return v - 'a' + 'A'
	}
	return v
}

// convert ascii char to upper case
func ToLower(v byte) byte {
	if IsUpper(v) {
		return v - 'A' + 'a'
	}
	return v
}
