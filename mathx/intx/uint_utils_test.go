package intx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "100", FormatUnsigned(100))
}

func TestParse(t *testing.T) {
	assert.Equal(t, uint(100), SafeParseUnsigned("100", 0))
	assert.Equal(t, uint(0), SafeParseUnsigned("100x", 0))
}
