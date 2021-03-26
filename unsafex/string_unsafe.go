package unsafex

import (
	"unsafe"
)

// BytesToString unsafe convert byte array content to string
// This function do not allocate new buffer, reuse the bytes data.
func BytesToString(bytes []byte) string {
	sliceHeader := (*SliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	sp := (*string)(unsafe.Pointer(&stringHeader))
	return *sp
}

// StringToBytes unsafe convert string to byte array
// This function do not allocate new buffer, reuse the string buffer.
func StringToBytes(str string) []byte {
	stringHeader := (*StringHeader)(unsafe.Pointer(&str))
	sliceHeader := SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len, // just same with len
	}
	sp := (*[]byte)(unsafe.Pointer(&sliceHeader))
	return *sp
}

// StringHeader copy from reflect.StringHeader, using Data type unsafe.Pointer instead of unitptr, to prevent stack allocation, or free by gc
type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// SliceHeader copy from reflect.SliceHeader, using Data type unsafe.Pointer instead of unitptr, to prevent stack allocation, or free by gc
type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
