package stringx

import (
	"fmt"
	"github.com/hsiafan/glow/intx"
	"strings"
)

// Joiner contains join config for joining string.
// This struct do not have internal state, hence can be reused.
type Joiner struct {
	Prefix     string // the prefix of joined string result
	Suffix     string // the suffix of joined string result
	Separator  string // the delimiter to join str
	OmitNil    bool   // if skip nil value
	OmitEmpty  bool   // if skip empty string
	NilToEmpty bool   // if trans nil value to empty string
}

// NewBuffer create and return one new JoinBuffer for join string using the Joiner
func (j *Joiner) NewBuffer() *JoinBuffer {
	return &JoinBuffer{
		Joiner: *j,
	}
}

// Join join string items to one string
func (j *Joiner) Join(strings []string) string {
	joiner := j.NewBuffer()
	joiner.AddAll(strings...)
	return joiner.String()
}

// JoinStringer join fmt.Stringer items to one string
func (j *Joiner) JoinStringer(stringers []fmt.Stringer) string {
	joiner := j.NewBuffer()
	joiner.AddAllStringer(stringers...)
	return joiner.String()
}

// JoinStringer join fmt.Stringer items to one string
func (j *Joiner) JoinAny(values []interface{}) string {
	joiner := j.NewBuffer()
	joiner.AddAllAny(values...)
	return joiner.String()
}

// JoinBuffer is a tool to join string with prefix, suffix, and delimiter.
//
// Usage:
//  joiner := &JoinBuffer{Separator:",", Prefix:"[", Suffix:"]"}
//  joiner.Add(str)
//  s := joiner.String()
type JoinBuffer struct {
	Joiner
	builder strings.Builder
	written bool
}

// Reset resets the JoinBuffer to be empty, can be reused.
func (j *JoinBuffer) Reset() *JoinBuffer {
	j.builder.Reset()
	j.written = false
	return j
}

// AddBytes add new data item to JoinBuffer. The binary data is treated as utf-8 encoded string.
func (j *JoinBuffer) AddBytes(data []byte) *JoinBuffer {
	if len(data) == 0 && j.OmitEmpty {
		return j
	}
	j.prepend()
	j.builder.Write(data)
	return j
}

// Add add a new string item to JoinBuffer
func (j *JoinBuffer) Add(str string) *JoinBuffer {
	if len(str) == 0 && j.OmitEmpty {
		return j
	}
	j.prepend()
	j.builder.WriteString(str)
	return j
}

// AddInt add a new int item to JoinBuffer
func (j *JoinBuffer) AddInt(value int) *JoinBuffer {
	j.prepend()
	j.builder.WriteString(intx.FormatInt(value))
	return j
}

// AddUint add a new uint item to JoinBuffer
func (j *JoinBuffer) AddUint(value uint) *JoinBuffer {
	j.prepend()
	j.builder.WriteString(intx.FormatUint(value))
	return j
}

// AddInt64 add a new int64 item to JoinBuffer
func (j *JoinBuffer) AddInt64(value int64) *JoinBuffer {
	j.prepend()
	j.builder.WriteString(intx.FormatInt64(value))
	return j
}

// AddUint64 add a new uint64 item to JoinBuffer
func (j *JoinBuffer) AddUint64(value uint64) *JoinBuffer {
	j.prepend()
	j.builder.WriteString(intx.FormatUint64(value))
	return j
}

// AddStringer add a new stringer item to JoinBuffer
func (j *JoinBuffer) AddStringer(value fmt.Stringer) *JoinBuffer {
	if value == nil && j.OmitNil {
		return j
	}
	if value == nil {
		if j.NilToEmpty {
			j.Add("")
		} else {
			j.Add(ValueOf(value))
		}
	} else {
		j.Add(value.String())
	}
	return j
}

// AddAny add a new value of any type item to JoinBuffer
func (j *JoinBuffer) AddAny(value interface{}) *JoinBuffer {
	if value == nil && j.OmitNil {
		return j
	}
	if value == nil {
		if j.NilToEmpty {
			j.Add("")
		} else {
			j.Add(ValueOf(value))
		}
	} else {
		j.Add(ValueOf(value))
	}
	return j
}

// AddAll add all strings to JoinBuffer
func (j *JoinBuffer) AddAll(ss ...string) *JoinBuffer {
	for _, s := range ss {
		j.Add(s)
	}
	return j
}

// AddAllStringer add all Stringer's string value to JoinBuffer
func (j *JoinBuffer) AddAllStringer(ss ...fmt.Stringer) *JoinBuffer {
	for _, s := range ss {
		j.AddStringer(s)
	}
	return j
}

// AddAllAny add all values as string to JoinBuffer
func (j *JoinBuffer) AddAllAny(ss ...interface{}) *JoinBuffer {
	for _, s := range ss {
		j.AddAny(s)
	}
	return j
}

// String join all values as string
func (j *JoinBuffer) String() string {
	if !j.written {
		j.builder.WriteString(j.Prefix)
		j.written = true
	}
	j.builder.WriteString(j.Suffix)
	return j.builder.String()
}

func (j *JoinBuffer) prepend() {
	if !j.written {
		j.builder.WriteString(j.Prefix)
		j.written = true
	} else {
		j.builder.WriteString(j.Separator)
	}
}
