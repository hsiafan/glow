package iox

import (
	"io"
)

// Close a resource, and ignore error. Useful for avoiding warning when use with defer:
// defer iox.Close(r)
func Close(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

// CloseIfIsCloser close a resource, if the resource is an io.Closer. If close return an error, the error is ignored.
func CloseIfIsCloser(resource interface{}) {
	if resource == nil {
		return
	}
	if closer, ok := resource.(io.Closer); ok {
		if closer != nil {
			_ = closer.Close()
		}
	}
}

// WithCloser run function, and then close the closer
func WithCloser(closer io.Closer, f func(io.Closer)) {
	defer Close(closer)
	f(closer)
}
