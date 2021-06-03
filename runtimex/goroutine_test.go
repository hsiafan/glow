package runtimex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentGoroutineId(t *testing.T) {
	id := CurrentGoroutineId()
	assert.NotEmpty(t, id)
}
