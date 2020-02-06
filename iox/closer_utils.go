package iox

import "io"

// close ignore error
func Close(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}
