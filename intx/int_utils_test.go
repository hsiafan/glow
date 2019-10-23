package intx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInt(t *testing.T) {
	assert.Equal(t, 100000000, SafeParse("100000000", -1))
}

func TestFormatInt(t *testing.T) {
	assert.Equal(t, "100", Format(100))
}
