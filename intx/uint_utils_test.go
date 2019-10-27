package intx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "100", FormatUint(100))
}

func TestParse(t *testing.T) {
	assert.Equal(t, uint(100), SafeParseUint("100", 0))
	assert.Equal(t, uint(0), SafeParseUint("100x", 0))
}
