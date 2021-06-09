package runtimex

import (
	"github.com/hsiafan/glow/internal"
	"unsafe"
)

// WrapAsByteSlice wrap a memory region as byte slice
func WrapAsByteSlice(p unsafe.Pointer, length int) []byte {
	sp := &internal.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}
	s := *(*[]byte)(unsafe.Pointer(sp))
	return s
}
