package glow

import (
	"sync/atomic"
	"unsafe"
)

// CopyChecker holds back pointer to itself to detect object copying.
// To use it, embed a CopyChecker field into your struct, and call Check before your operations.
type CopyChecker uintptr

// Check check if is copied, and panic if is copied.
func (c *CopyChecker) Check() {
	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		panic("copy detected")
	}
}

// NoCopy may be embedded into structs which must not be copied after the first use.
// It behaves as a sync.Locker, so can be detected by go vet.
type NoCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}
