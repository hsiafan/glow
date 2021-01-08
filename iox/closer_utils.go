package iox

import "io"

// Close a resource, and ignore error. Useful for avoiding warning when use with defer:
// defer iox.Close(r)
func Close(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

// WithCloser run function, and then close the closer
func WithCloser(closer io.Closer, f func(io.Closer)) {
	defer Close(closer)
	f(closer)
}
