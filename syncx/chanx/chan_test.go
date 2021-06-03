package chanx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsClosed(t *testing.T) {
	c := make(chan int)
	assert.False(t, IsClosed(c))
	close(c)
	assert.True(t, IsClosed(c))

	cc := make(chan struct{})
	assert.False(t, IsClosed(cc))
	close(cc)
	assert.True(t, IsClosed(c))
}
