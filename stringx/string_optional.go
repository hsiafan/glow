package stringx

// OptionalString represent a string value, or non-exists.
type OptionalString struct {
	Value  string
	Exists bool
}

// Or return the value of optional if exists, return passed in candidate str if not exists.
func (os *OptionalString) Or(str string) string {
	if os.Exists {
		return os.Value
	}
	return str
}
