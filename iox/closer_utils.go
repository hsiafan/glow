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

// TryClose close a resource, if the resource is an io.Closer. If close return an error, the error is ignored.
func TryClose(resource interface{}) {
	if resource == nil {
		return
	}
	if closer, ok := resource.(io.Closer); ok {
		if closer != nil {
			_ = closer.Close()
		}
	}
}

// CloseMulti close multi closers and ignore errors
func CloseMulti(closers ...io.Closer) {
	for _, closer := range closers {
		_ = closer.Close()
	}
}

// TryCloseMulti try close multi resources, if it is a closer
func TryCloseMulti(resources ...interface{}) {
	for _, r := range resources {
		TryClose(r)
	}
}

// TryCloseMultiReader try close multi readers, if it is a closer
func TryCloseMultiReader(readers ...io.Reader) {
	for _, r := range readers {
		TryClose(r)
	}
}

// TryCloseMultiWriter try close multi resources, if it is a closer
func TryCloseMultiWriter(writers ...io.Writer) {
	for _, w := range writers {
		TryClose(w)
	}
}
