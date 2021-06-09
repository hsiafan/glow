package stringx

import (
	"github.com/hsiafan/glow/internal"
	"unsafe"
)

// FromBytes unsafe convert byte array content to string
// This function do not allocate new buffer, reuse the bytes data.
func FromBytes(bytes []byte) string {
	sliceHeader := (*internal.SliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := internal.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	sp := (*string)(unsafe.Pointer(&stringHeader))
	return *sp
}

// ToBytes unsafe convert string to byte array
// This function do not allocate new buffer, reuse the string buffer.
func ToBytes(str string) []byte {
	stringHeader := (*internal.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := internal.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len, // just same with len
	}
	sp := (*[]byte)(unsafe.Pointer(&sliceHeader))
	return *sp
}
