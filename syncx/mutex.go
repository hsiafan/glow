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

// TryRLock try to add a read lock. If cannot lock due to a writer hold the lock, return false immediately.
func TryRLock(m *sync.RWMutex) bool {
	rw := (*internal.RWMutex)(unsafe.Pointer(m))
	if atomic.AddInt32(&rw.ReaderCount, 1) < 0 {
		// A writer is pending, wait for it.
		return false
	}
	return true
}

const rwmutexMaxReaders = 1 << 30

// TryWLock try to add a write lock. If cannot lock, return false immediately.
func TryWLock(m *sync.RWMutex) bool {
	rw := (*internal.RWMutex)(unsafe.Pointer(m))
	// First, resolve competition with other writers.
	if !TryLock(&rw.WLock) {
		return false
	}
	// Announce to readers there is a pending writer.
	r := atomic.AddInt32(&rw.ReaderCount, -rwmutexMaxReaders) + rwmutexMaxReaders
	// Wait for active readers.
	if r != 0 && atomic.AddInt32(&rw.ReaderWait, r) != 0 {
		return false
	}
	return true
}
