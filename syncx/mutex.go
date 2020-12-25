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

// WithLock run code within protection of lock.
func WithLock(m *sync.Mutex, f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}
