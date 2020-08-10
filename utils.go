package glow

// Panic if err is not nil
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
