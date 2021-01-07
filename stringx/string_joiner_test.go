package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStringer int

func (s testStringer) String() string {
	return "test"
}

func TestJoiner_Add(t *testing.T) {
	var joiner = Joiner{
		Prefix:    "[",
		Suffix:    "]",
		Separator: ",",
	}
	var testStringer testStringer
	str := joiner.AddBytes([]byte("testAdd")).Add("testString").AddInt(1).AddUint(2).
		AddInt64(3).AddUint64(4).AddStringer(testStringer).AddAny(1234).
		String()
	assert.Equal(t, "[testAdd,testString,1,2,3,4,test,1234]", str)

	joiner.Reset()
	assert.Equal(t, "[]", joiner.String())

	joiner.Reset()
	assert.Equal(t, "[123]", joiner.Add("123").String())

	joiner = Joiner{
		Prefix:    "[",
		Suffix:    "]",
		Separator: ",",
		OmitNil:   true,
	}
	str = joiner.Add("1").AddAny(nil).AddStringer(nil).String()
	assert.Equal(t, "[1]", str)

	joiner = Joiner{
		Prefix:    "[",
		Suffix:    "]",
		Separator: ",",
		OmitEmpty: true,
	}
	str = joiner.Add("1").Add("").String()
	assert.Equal(t, "[1]", str)

	joiner = Joiner{
		Prefix:     "[",
		Suffix:     "]",
		Separator:  ",",
		NilToEmpty: true,
	}
	str = joiner.Add("1").AddAny(nil).AddStringer(nil).String()
	assert.Equal(t, "[1,,]", str)

	joiner = Joiner{
		Prefix:     "[",
		Suffix:     "]",
		Separator:  ",",
		OmitEmpty:  true,
		NilToEmpty: true,
	}
	str = joiner.Add("1").AddAny(nil).AddStringer(nil).String()
	assert.Equal(t, "[1]", str)
}
