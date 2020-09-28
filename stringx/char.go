package stringx

// IsUpperASCII return if v is a upper case ascii char.
func IsUpperASCII(v byte) bool {
	return 'A' <= v && v <= 'Z'
}

// IsLowerASCII return if v is a lower case ascii char.
func IsLowerASCII(v byte) bool {
	return 'a' <= v && v <= 'z'
}

// ToUpperASCII convert ascii char to upper case.
// Return v if v is not a lower case ascii character.
func ToUpperASCII(v byte) byte {
	if IsLowerASCII(v) {
		return v - 'a' + 'A'
	}
	return v
}

// ToLowerASCII convert ascii char to upper case
// Return v if v is not a upper case ascii character.
func ToLowerASCII(v byte) byte {
	if IsUpperASCII(v) {
		return v - 'A' + 'a'
	}
	return v
}
