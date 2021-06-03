package stringx

import (
	"unsafe"
)

// FromBytes unsafe convert byte array content to string
// This function do not allocate new buffer, reuse the bytes data.
func FromBytes(bytes []byte) string {
	sliceHeader := (*sliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := stringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	sp := (*string)(unsafe.Pointer(&stringHeader))
	return *sp
}

// ToBytes unsafe convert string to byte array
// This function do not allocate new buffer, reuse the string buffer.
func ToBytes(str string) []byte {
	stringHeader := (*stringHeader)(unsafe.Pointer(&str))
	sliceHeader := sliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len, // just same with len
	}
	sp := (*[]byte)(unsafe.Pointer(&sliceHeader))
	return *sp
}

// stringHeader copy from reflect.StringHeader, using Data type unsafe.Pointer instead of unitptr, to prevent stack allocation, or free by gc
type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// sliceHeader copy from reflect.SliceHeader, using Data type unsafe.Pointer instead of unitptr, to prevent stack allocation, or free by gc
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
