package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKVJoiner(t *testing.T) {
	joiner := &KVJoiner{
		Prefix:      "[",
		Suffix:      "]",
		Separator:   ", ",
		KVSeparator: "=",
	}
	joiner.Add("a", "1")
	joiner.AddAll(map[string]string{
		"b": "2",
	})
	joiner.AddAny("c", "3")
	assert.Equal(t, "[a=1, b=2, c=3]", joiner.String())
}
