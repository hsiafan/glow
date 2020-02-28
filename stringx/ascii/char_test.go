package ascii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isLower(t *testing.T) {
	assert.True(t, IsLower('a'))
	assert.False(t, IsLower('-'))
	assert.False(t, IsLower('A'))
}

func Test_isUpper(t *testing.T) {
	assert.False(t, IsUpper('a'))
	assert.False(t, IsUpper('-'))
	assert.True(t, IsUpper('A'))
}

func Test_toLower(t *testing.T) {
	assert.Equal(t, byte('a'), ToLower('a'))
	assert.Equal(t, byte('-'), ToLower('-'))
	assert.Equal(t, byte('a'), ToLower('A'))
}

func Test_toUpper(t *testing.T) {
	assert.Equal(t, byte('A'), ToUpper('a'))
	assert.Equal(t, byte('-'), ToUpper('-'))
	assert.Equal(t, byte('A'), ToUpper('A'))
}
