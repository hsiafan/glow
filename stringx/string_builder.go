package stringx

import (
	"fmt"
	"strconv"
	"strings"
)

// Builder is a string builder, which can append values of well known types.
// The zero value is ready to use.
type Builder struct {
	builder strings.Builder
}

// String returns the accumulated string.
func (b *Builder) String() string {
	return b.builder.String()
}

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *Builder) Len() int {
	return b.builder.Len()
}

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (b *Builder) Cap() int {
	return b.builder.Cap()
}

// Reset resets the Builder to be empty.
func (b *Builder) Reset() *Builder {
	b.builder.Reset()
	return b
}

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
func (b *Builder) Grow(n int) *Builder {
	b.builder.Grow(n)
	return b
}

// Write appends the contents of p to b's buffer.
func (b *Builder) Write(p []byte) *Builder {
	if _, err := b.builder.Write(p); err != nil {
		panic(err)
	}
	return b
}

// WriteByte appends the byte c to b's buffer.
func (b *Builder) WriteByte(c byte) *Builder {
	if err := b.builder.WriteByte(c); err != nil {
		panic(err)
	}
	return b
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
func (b *Builder) WriteRune(r rune) *Builder {
	if _, err := b.builder.WriteRune(r); err != nil {
		panic(err)
	}
	return b
}

// WriteString appends the contents of s to b's buffer.
func (b *Builder) WriteString(s string) *Builder {
	if _, err := b.builder.WriteString(s); err != nil {
		panic(err)
	}
	return b
}

// WriteStringer appends the stringer value to b's buffer.
func (b *Builder) WriteStringer(stringer fmt.Stringer) *Builder {
	return b.WriteString(stringer.String())
}

// WriteAny appends a value of any type to b's buffer.
func (b *Builder) WriteAny(value interface{}) *Builder {
	if _, err := fmt.Fprint(&b.builder, value); err != nil {
		panic(err)
	}
	return b
}

// WriteInt appends the int value to b's buffer.
func (b *Builder) WriteInt(value int) *Builder {
	return b.WriteString(strconv.Itoa(value))
}

// WriteUint appends the uint value to b's buffer.
func (b *Builder) WriteUint(value uint) *Builder {
	return b.WriteString(strconv.FormatUint(uint64(value), 10))
}

// WriteInt64 appends the int64 value to b's buffer.
func (b *Builder) WriteInt64(value int64) *Builder {
	return b.WriteString(strconv.FormatInt(value, 10))
}

// WriteUint64 appends the uint64 to b's buffer.
func (b *Builder) WriteUint64(value uint64) *Builder {
	return b.WriteString(strconv.FormatUint(value, 10))
}
