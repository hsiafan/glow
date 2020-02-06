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

// Mutex with TryLock, etc...
type Mutex struct {
	sync.Mutex
}

// Try lock
func (m *Mutex) TryLock() bool {
	// state is first field
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
	return false
}
