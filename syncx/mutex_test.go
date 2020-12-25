package syncx

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestTryLock(t *testing.T) {
	var m sync.Mutex
	m.Lock()
	assert.False(t, TryLock(&m))
	m.Unlock()
	assert.True(t, TryLock(&m))
	m.Unlock()
}

func TestWithLock(t *testing.T) {
	var m sync.Mutex
	WithLock(&m, func() {

	})

}
