package stringx

import "fmt"

// Joiner is a tool to join string with prefix, suffix, and delimiter
type Joiner struct {
	Prefix    string // the prefix of joined string result
	Suffix    string // the suffix of joined string result
	Delimiter string // the delimiter to join str
	builder   Builder
	written   bool
}

// Reset resets the Joiner to be empty, can be reused.
func (j *Joiner) Reset() *Joiner {
	j.builder.Reset()
	j.written = false
	return j
}

// Add add new data item to Joiner
func (j *Joiner) Add(data []byte) *Joiner {
	j.prepend()
	j.builder.Write(data)
	return j
}

// Add add a new string item to Joiner
func (j *Joiner) AddString(str string) *Joiner {
	j.prepend()
	j.builder.WriteString(str)
	return j
}

// Add add a new int item to Joiner
func (j *Joiner) AddInt(value int) *Joiner {
	j.prepend()
	j.builder.WriteInt(value)
	return j
}

// Add add a new uint item to Joiner
func (j *Joiner) AddUint(value uint) *Joiner {
	j.prepend()
	j.builder.WriteUint(value)
	return j
}

// Add add a new int64 item to Joiner
func (j *Joiner) AddInt64(value int64) *Joiner {
	j.prepend()
	j.builder.WriteInt64(value)
	return j
}

// Add add a new uint64 item to Joiner
func (j *Joiner) AddUint64(value uint64) *Joiner {
	j.prepend()
	j.builder.WriteUint64(value)
	return j
}

// Add add a new stringer item to Joiner
func (j *Joiner) AddStringer(value fmt.Stringer) *Joiner {
	j.prepend()
	j.builder.WriteStringer(value)
	return j
}

// Add add a new value of any type item to Joiner
func (j *Joiner) AddAny(value interface{}) *Joiner {
	j.prepend()
	j.builder.WriteAny(value)
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
		j.builder.WriteString(j.Delimiter)
	}
}
