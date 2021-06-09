package runtimex

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestWrapAsByteSlice(t *testing.T) {
	length := 10000
	var s = make([]byte, length)
	s2 := WrapAsByteSlice(unsafe.Pointer(&s[0]), length)
	s2[0] = 1
	s2[length-1] = 2
	assert.Equal(t, uint8(1), s[0])
	assert.Equal(t, uint8(2), s[length-1])
}
