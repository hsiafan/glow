package stringx

import "strings"

// Splitter is a tool to split string with prefix, suffix, and separator into pieces.
//
// Usage:
//  splitter := &Splitter{Separator:",", Prefix:"[", Suffix:"]"}
//  ss := splitter.Split(str)
type Splitter struct {
	Prefix    string // the prefix of string to remove
	Suffix    string // the suffix of string to remove
	Separator string // the delimiter to split str
}

// Split spit a string into pieces, using prefix, suffix, and separator of this Splitter.
func (sp *Splitter) Split(s string) []string {
	if sp.Prefix != "" && strings.HasPrefix(s, sp.Prefix) {
		s = s[len(sp.Prefix):]
	}
	if sp.Suffix != "" && strings.HasSuffix(s, sp.Suffix) {
		s = s[:len(s)-len(sp.Suffix)]
	}
	return strings.Split(s, sp.Separator)
}

// StringEntry is a key-value string entry.
type StringEntry struct {
	Key   string // the key
	Value string // the value
}

// KVSplitter is a tool to split string with prefix, suffix, and separator into key-value entries.
//
// Usage:
//  splitter := &KVSplitter{Separator:",", Prefix:"[", Suffix:"]", KVSeparator: "="}
//  entries := splitter.Split(str)
type KVSplitter struct {
	Prefix      string // the prefix of string to remove
	Suffix      string // the suffix of string to remove
	Separator   string // the delimiter to split str
	KVSeparator string // the delimiter to split key and value
}

// Split spit a string into pieces, using prefix, suffix, and separator of this KVSplitter.
func (sp *KVSplitter) Split(s string) []StringEntry {
	if sp.Prefix != "" && strings.HasPrefix(s, sp.Prefix) {
		s = s[len(sp.Prefix):]
	}
	if sp.Suffix != "" && strings.HasSuffix(s, sp.Suffix) {
		s = s[:len(s)-len(sp.Suffix)]
	}
	items := strings.Split(s, sp.Separator)
	entries := make([]StringEntry, len(items))
	for i, item := range items {
		key, value := Split2(item, sp.KVSeparator)
		entries[i] = StringEntry{key, value}
	}
	return entries
}

// SplitToMap spit a string into string map, using prefix, suffix, and separator of this KVSplitter.
// If has multi values for one same key, the last value is stored.
func (sp *KVSplitter) SplitToMap(s string) map[string]string {
	if sp.Prefix != "" && strings.HasPrefix(s, sp.Prefix) {
		s = s[len(sp.Prefix):]
	}
	if sp.Suffix != "" && strings.HasSuffix(s, sp.Suffix) {
		s = s[:len(s)-len(sp.Suffix)]
	}
	items := strings.Split(s, sp.Separator)
	m := map[string]string{}
	for _, item := range items {
		key, value := Split2(item, sp.KVSeparator)
		m[key] = value
	}
	return m
}
