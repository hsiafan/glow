package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesToString(t *testing.T) {
	assert.Equal(t, "test", FromBytes([]byte("test")))
}

func TestStringToBytes(t *testing.T) {
	assert.Equal(t, []byte("test"), ToBytes("test"))
}
