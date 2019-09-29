package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoiner_Add(t *testing.T) {
	var joiner = Joiner{
		Prefix:    "[",
		Suffix:    "]",
		Delimiter: ",",
	}
	var testStringer testStringer
	str := joiner.Add([]byte("testAdd")).AddString("testString").AddInt(1).AddUint(2).
		AddInt64(3).AddUint64(4).AddStringer(testStringer).AddAny(1234).
		String()
	assert.Equal(t, "[testAdd,testString,1,2,3,4,test,1234]", str)

	joiner.Reset()
	assert.Equal(t, "[]", joiner.String())

	joiner.Reset()
	assert.Equal(t, "[123]", joiner.AddString("123").String())
}
