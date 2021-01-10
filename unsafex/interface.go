package unsafex

import "unsafe"

// A struct having same structure with interface
type eface struct {
	_type unsafe.Pointer
	value unsafe.Pointer
}

// InterfaceValuePtr return the value pointer of interface
func InterfaceValuePtr(iface interface{}) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&iface)).value
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
