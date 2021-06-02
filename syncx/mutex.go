package syncx

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

// TryLock try to lock a mutex, if cannot acquire, return false immediately.
// If acquired, return true.
func TryLock(m *sync.Mutex) bool {
	// state is first field
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(m)), 0, mutexLocked) {
		return true
	}
	return false
}

// IsLocked return weather lock is locked
func IsLocked(m *sync.Mutex) bool {
	return atomic.LoadInt32((*int32)(unsafe.Pointer(m))) == mutexLocked
}
