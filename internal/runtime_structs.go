package internal

import "unsafe"

// SliceHeader copy from reflect.SliceHeader, using Data type unsafe.Pointer instead of unitptr, to prevent stack allocation, or free by gc
type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// StringHeader copy from reflect.StringHeader, using Data type unsafe.Pointer instead of unitptr, to prevent stack allocation, or free by gc
type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// Interface having same structure with go interface
type Interface struct {
	Type  unsafe.Pointer
	Value unsafe.Pointer
}

// Mutex copy from sync.Mutex
type Mutex struct {
	State int32
	Sema  uint32
}

// Chan for go channel
type Chan struct {
	QueueCount    uint           // total data in the queue
	DataQueueSize uint           // size of the circular queue
	Buf           unsafe.Pointer // points to an array of dataqsiz elements
	ElemSize      uint16
	Closed        uint32
}
