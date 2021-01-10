package unsafex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceIsNil(t *testing.T) {
	assert.True(t, InterfaceIsNil(nil))
	var face interface{}
	var i *int = nil
	face = i
	assert.True(t, InterfaceIsNil(face))
	assert.False(t, InterfaceIsNil(1))
}
