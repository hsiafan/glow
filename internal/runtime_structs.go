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
	Type  *RType
	Value unsafe.Pointer
}

const (
	KindDirectIface = 1 << 5
	KindMask        = (1 << 5) - 1
)

// RType is the common implementation of most values.
type RType struct {
	Size       uintptr
	PtrData    uintptr // number of bytes in the type that can contain pointers
	Hash       uint32  // hash of type; avoids computation in hash tables
	Flag       uint8   // extra type information flags
	Align      uint8   // alignment of variable with this type
	FieldAlign uint8   // alignment of struct field with this type
	Kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal     func(unsafe.Pointer, unsafe.Pointer) bool
	GCData    *byte // garbage collection data
	Str       int32 // string form
	PtrToThis int32 // type for pointer to this type, may be zero
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
