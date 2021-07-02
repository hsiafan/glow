package floatx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatFloat(t *testing.T) {
	assert.Equal(t, "1.20", Format64(1.201, 2))
	assert.Equal(t, "1.20", Format64(1.198, 2))
	assert.Equal(t, "1.20", Format32(1.201, 2))
	assert.Equal(t, "1.20", Format32(1.198, 2))
}

func TestParseFloat(t *testing.T) {
	assert.Equal(t, 1.20, SafeParse64("1.20", 0))
	assert.Equal(t, 0.0, SafeParse64("1.20x", 0))
}
