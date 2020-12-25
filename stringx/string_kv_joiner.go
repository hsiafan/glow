package stringx

import "strings"

// KVJoiner is a tool to join string with prefix, suffix, and delimiter
// Usage:
//  joiner := &KVJoiner{
//		Prefix:      "[",
//		Suffix:      "]",
//		Separator:   ", ",
//		KVSeparator: "=",
//	}
//	joiner.Add("a", "1")
//	joiner.AddAll(map[string]string{
//		"b": "2",
//	})
//	joiner.AddAny("c", "3")
type KVJoiner struct {
	Prefix      string // the prefix of joined string result
	Suffix      string // the suffix of joined string result
	Separator   string // the str to join kv items
	KVSeparator string // the str to join key and value
	builder     strings.Builder
	written     bool
}

// Reset resets the KVJoiner to be empty, can be reused.
func (j *KVJoiner) Reset() {
	j.builder.Reset()
	j.written = false
}

// Add add a new string key-value entry to KVJoiner
func (j *KVJoiner) Add(key string, value string) {
	j.prepend()
	j.builder.WriteString(key)
	j.builder.WriteString(j.KVSeparator)
	j.builder.WriteString(value)
}

// AddAny add a new key-value entry to KVJoiner
func (j *KVJoiner) AddAny(key interface{}, value interface{}) {
	j.prepend()
	j.builder.WriteString(ValueOf(key))
	j.builder.WriteString(j.KVSeparator)
	j.builder.WriteString(ValueOf(value))
}

// AddAny add all key-value items in a map to
func (j *KVJoiner) AddAll(m map[string]string) {
	for k, v := range m {
		j.Add(k, v)
	}
}

// String join all values as string
func (j *KVJoiner) String() string {
	if !j.written {
		j.builder.WriteString(j.Prefix)
		j.written = true
	}
	j.builder.WriteString(j.Suffix)
	return j.builder.String()
}

func (j *KVJoiner) prepend() {
	if !j.written {
		j.builder.WriteString(j.Prefix)
		j.written = true
	} else {
		j.builder.WriteString(j.Separator)
	}
}
