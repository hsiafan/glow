package iox

import "io"

// close ignore error
func CloseQuite(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}
