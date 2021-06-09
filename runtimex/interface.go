package runtimex

import (
	"github.com/hsiafan/glow/internal"
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
