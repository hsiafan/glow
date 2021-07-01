package runtimex

import (
	"github.com/hsiafan/glow/internal"
	"reflect"
	"unsafe"
)

// InterfaceValuePtr return the value pointer of interface
func InterfaceValuePtr(iface interface{}) unsafe.Pointer {
	return (*internal.Interface)(unsafe.Pointer(&iface)).Value
}

// InterfaceIsNil return true if interface is nil or it's value is a nil pointer
func InterfaceIsNil(iface interface{}) bool {
	if iface == nil {
		return true
	}
	if InterfaceValuePtr(iface) == nil {
		return true
	}
	return false
}

// InterfaceIsIndirect return true, if iface is not nil, and contains a pointer(which maybe nil).
// Otherwise, return false.
func InterfaceIsIndirect(iface interface{}) bool {
	if iface == nil {
		return false
	}
	_type := (*internal.Interface)(unsafe.Pointer(&iface)).Type
	return _type.Kind&internal.KindDirectIface != 0
}

// InterfaceKind the kind of value iface contains. Return reflect.Invalid if interface is nil.
func InterfaceKind(iface interface{}) reflect.Kind {
	if iface == nil {
		return reflect.Invalid
	}
	_type := (*internal.Interface)(unsafe.Pointer(&iface)).Type
	return reflect.Kind(_type.Kind & internal.KindMask)
}
