package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isLower(t *testing.T) {
	assert.True(t, IsLowerASCII('a'))
	assert.False(t, IsLowerASCII('-'))
	assert.False(t, IsLowerASCII('A'))
}

func Test_isUpper(t *testing.T) {
	assert.False(t, IsUpperASCII('a'))
	assert.False(t, IsUpperASCII('-'))
	assert.True(t, IsUpperASCII('A'))
}

func Test_toLower(t *testing.T) {
	assert.Equal(t, byte('a'), ToLowerASCII('a'))
	assert.Equal(t, byte('-'), ToLowerASCII('-'))
	assert.Equal(t, byte('a'), ToLowerASCII('A'))
}

func Test_toUpper(t *testing.T) {
	assert.Equal(t, byte('A'), ToUpperASCII('a'))
	assert.Equal(t, byte('-'), ToUpperASCII('-'))
	assert.Equal(t, byte('A'), ToUpperASCII('A'))
}
