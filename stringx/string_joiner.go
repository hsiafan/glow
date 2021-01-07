package stringx

import (
	"fmt"
	"github.com/hsiafan/glow/intx"
	"strings"
)

// JoinConfig contains join config for joining string
type JoinConfig struct {
	Prefix     string // the prefix of joined string result
	Suffix     string // the suffix of joined string result
	Separator  string // the delimiter to join str
	OmitNil    bool   // if skip nil value
	OmitEmpty  bool   // if skip empty string
	NilToEmpty bool   // if trans nil value to empty string
}

// NewJoiner create and return on new Joiner using the JoinConfig
func (j *JoinConfig) NewJoiner() *Joiner {
	return &Joiner{
		Prefix:     j.Prefix,
		Suffix:     j.Suffix,
		Separator:  j.Separator,
		OmitNil:    j.OmitNil,
		OmitEmpty:  j.OmitEmpty,
		NilToEmpty: j.NilToEmpty,
	}
}

// Join join string items to one string
func (j *JoinConfig) Join(strings []string) string {
	joiner := j.NewJoiner()
	joiner.AddAll(strings...)
	return joiner.String()
}

// JoinStringer join fmt.Stringer items to one string
func (j *JoinConfig) JoinStringer(stringers []fmt.Stringer) string {
	joiner := j.NewJoiner()
	joiner.AddAllStringer(stringers...)
	return joiner.String()
}

// JoinStringer join fmt.Stringer items to one string
func (j *JoinConfig) JoinAny(values []interface{}) string {
	joiner := j.NewJoiner()
	joiner.AddAllAny(values...)
	return joiner.String()
}

// Joiner is a tool to join string with prefix, suffix, and delimiter.
//
// Usage:
//  joiner := &Joiner{Separator:",", Prefix:"[", Suffix:"]"}
//  joiner.Add(str)
//  s := joiner.String()
type Joiner struct {
	Prefix     string // the prefix of joined string result
	Suffix     string // the suffix of joined string result
	Separator  string // the delimiter to join str
	OmitNil    bool   // if skip nil value
	OmitEmpty  bool   // if skip empty string
	NilToEmpty bool   // if trans nil value to empty string
	builder    strings.Builder
	written    bool
}

// Reset resets the Joiner to be empty, can be reused.
func (j *Joiner) Reset() *Joiner {
	j.builder.Reset()
	j.written = false
	return j
}

// AddBytes add new data item to Joiner. The binary data is treated as utf-8 encoded string.
func (j *Joiner) AddBytes(data []byte) *Joiner {
	if len(data) == 0 && j.OmitEmpty {
		return j
	}
	j.prepend()
	j.builder.Write(data)
	return j
}

// Add add a new string item to Joiner
func (j *Joiner) Add(str string) *Joiner {
	if len(str) == 0 && j.OmitEmpty {
		return j
	}
	j.prepend()
	j.builder.WriteString(str)
	return j
}

// AddInt add a new int item to Joiner
func (j *Joiner) AddInt(value int) *Joiner {
	j.prepend()
	j.builder.WriteString(intx.FormatInt(value))
	return j
}

// AddUint add a new uint item to Joiner
func (j *Joiner) AddUint(value uint) *Joiner {
	j.prepend()
	j.builder.WriteString(intx.FormatUint(value))
	return j
}

// AddInt64 add a new int64 item to Joiner
func (j *Joiner) AddInt64(value int64) *Joiner {
	j.prepend()
	j.builder.WriteString(intx.FormatInt64(value))
	return j
}

// AddUint64 add a new uint64 item to Joiner
func (j *Joiner) AddUint64(value uint64) *Joiner {
	j.prepend()
	j.builder.WriteString(intx.FormatUint64(value))
	return j
}

// AddStringer add a new stringer item to Joiner
func (j *Joiner) AddStringer(value fmt.Stringer) *Joiner {
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

// AddAny add a new value of any type item to Joiner
func (j *Joiner) AddAny(value interface{}) *Joiner {
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

// AddAll add all strings to Joiner
func (j *Joiner) AddAll(ss ...string) *Joiner {
	for _, s := range ss {
		j.Add(s)
	}
	return j
}

// AddAllStringer add all Stringer's string value to Joiner
func (j *Joiner) AddAllStringer(ss ...fmt.Stringer) *Joiner {
	for _, s := range ss {
		j.AddStringer(s)
	}
	return j
}

// AddAllAny add all values as string to Joiner
func (j *Joiner) AddAllAny(ss ...interface{}) *Joiner {
	for _, s := range ss {
		j.AddAny(s)
	}
	return j
}

// String join all values as string
func (j *Joiner) String() string {
	if !j.written {
		j.builder.WriteString(j.Prefix)
		j.written = true
	}
	j.builder.WriteString(j.Suffix)
	return j.builder.String()
}

func (j *Joiner) prepend() {
	if !j.written {
		j.builder.WriteString(j.Prefix)
		j.written = true
	} else {
		j.builder.WriteString(j.Separator)
	}
}
