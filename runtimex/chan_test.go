package runtimex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsClosed(t *testing.T) {
	c := make(chan int)
	assert.False(t, ChanClosed(c))
	close(c)
	assert.True(t, ChanClosed(c))

	cc := make(chan struct{})
	assert.False(t, ChanClosed(cc))
	close(cc)
	assert.True(t, ChanClosed(c))
}
