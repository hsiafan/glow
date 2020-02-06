package unsafex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesToString(t *testing.T) {
	assert.Equal(t, "test", BytesToString([]byte("test")))
}

func TestStringToBytes(t *testing.T) {
	assert.Equal(t, []byte("test"), StringToBytes("test"))
}
