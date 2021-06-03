package runtimex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsClosed(t *testing.T) {
	c := make(chan int)
	assert.False(t, ChanIsClosed(c))
	close(c)
	assert.True(t, ChanIsClosed(c))

	cc := make(chan struct{})
	assert.False(t, ChanIsClosed(cc))
	close(cc)
	assert.True(t, ChanIsClosed(c))
}
