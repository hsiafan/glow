package uintx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "100", Format(100))
}

func TestParse(t *testing.T) {
	assert.Equal(t, uint(100), SafeParse("100", 0))
	assert.Equal(t, uint(0), SafeParse("100x", 0))
}
