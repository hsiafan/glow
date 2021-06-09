package syncx

import (
	"github.com/hsiafan/glow/internal"
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
	mm := (*internal.Mutex)(unsafe.Pointer(m))
	// state is first field
	if atomic.CompareAndSwapInt32(&mm.State, 0, mutexLocked) {
		return true
	}
	return false
}

// IsLocked return weather lock is locked
func IsLocked(m *sync.Mutex) bool {
	mm := (*internal.Mutex)(unsafe.Pointer(m))
	return atomic.LoadInt32(&mm.State) == mutexLocked
}
