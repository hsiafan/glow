package reflectx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaller(t *testing.T) {
	caller := GetCaller(1)
	assert.Equal(t, "github.com/hsiafan/glow/reflectx", caller.Package)
	assert.Equal(t, "caller_test.go", caller.File)
	assert.Equal(t, "TestCaller", caller.Function)
	assert.Equal(t, 9, caller.LineNo)
}
