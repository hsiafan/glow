package stringx

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestBuilder_WriteInt(t *testing.T) {
	var builder Builder
	assert.Equal(t, "10", builder.WriteInt(10).String())
}

func TestBuilder_WriteInt64(t *testing.T) {
	var builder Builder
	var value int64 = 1000
	assert.Equal(t, "1000", builder.WriteInt64(value).String())
}

type testStringer struct {
}

func (t testStringer) String() string {
	return "test"
}

func TestBuilder_WriteStringer(t *testing.T) {
	var builder Builder
	var value testStringer
	assert.Equal(t, "test", builder.WriteStringer(value).String())
}

func TestBuilder_WriteAny(t *testing.T) {
	var builder Builder
	var value testStringer
	assert.Equal(t, "test1", builder.WriteAny(value).WriteAny(1).String())
}

func TestBuilder_WriteUInt(t *testing.T) {
	var builder Builder
	assert.Equal(t, "10", builder.WriteUint(10).String())
}

func TestBuilder_WriteUInt64(t *testing.T) {
	var builder Builder
	var value uint64 = 1000
	assert.Equal(t, "1000", builder.WriteUint64(value).String())
}
