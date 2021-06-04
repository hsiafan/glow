package ascii

// IsUpper return if v is a upper case ascii char.
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

// IsDigit return if is a digit char([0, 9])
func IsDigit(v byte) bool {
	return '0' <= v && v <= '9'
}

// IsAlphaBet return if is a ascii alphabet char([A, Z], [a, z])
func IsAlphaBet(v byte) bool {
	return 'A' <= v && v <= 'Z' || 'a' <= v && v <= 'z'
}

// IsHex return if is a hex digit char([0, 9], [A, F], [a, f])
func IsHex(v byte) bool {
	return '0' <= v && v <= '9' || 'a' <= v && v <= 'f' || 'A' <= v && v <= 'F'
}
