package runtimex

import (
	"github.com/stretchr/testify/assert"
	"reflect"
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

func TestInterfaceIsIndirect(t *testing.T) {
	var a interface{} = 1
	assert.False(t, InterfaceIsIndirect(a))

	var i = 1
	var b interface{} = &i
	assert.True(t, InterfaceIsIndirect(b))

	var c interface{} = nil
	assert.False(t, InterfaceIsIndirect(c))

	var ip *int = nil
	var d interface{} = ip
	assert.True(t, InterfaceIsIndirect(d))
}

func TestInterfaceKind(t *testing.T) {
	assert.Equal(t, reflect.Invalid, InterfaceKind(nil))
	assert.Equal(t, reflect.Int, InterfaceKind(1))
	var i = 1
	assert.Equal(t, reflect.Ptr, InterfaceKind(&i))
	assert.Equal(t, reflect.Func, InterfaceKind(func() {}))
	assert.Equal(t, reflect.Slice, InterfaceKind([]int{}))
	assert.Equal(t, reflect.Map, InterfaceKind(map[int]int{}))
	assert.Equal(t, reflect.String, InterfaceKind(""))
	assert.Equal(t, reflect.Struct, InterfaceKind(struct{}{}))
}
