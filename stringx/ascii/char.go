package ascii

// IsLower return if v is a upper case ascii char.
func IsUpper(v byte) bool {
	return 'A' <= v && v <= 'Z'
}

// IsLower return if v is a lower case ascii char.
func IsLower(v byte) bool {
	return 'a' <= v && v <= 'z'
}

// ToUpper convert ascii char to upper case.
// Return v if v is not a lower case ascii character.
func ToUpper(v byte) byte {
	if IsLower(v) {
		return v - 'a' + 'A'
	}
	return v
}

// ToLower convert ascii char to upper case
// Return v if v is not a upper case ascii character.
func ToLower(v byte) byte {
	if IsUpper(v) {
		return v - 'A' + 'a'
	}
	return v
}
