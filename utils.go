package glow

// Panic if err is not nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
