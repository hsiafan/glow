package jsonx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	json, err := Marshal(1)
	assert.NoError(t, err)
	assert.Equal(t, "1", json)

	json, err = Marshal("test")
	assert.NoError(t, err)
	assert.Equal(t, "\"test\"", json)

	json, err = Marshal([]int{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, "[1,2]", json)
}

func TestMarshalIndent(t *testing.T) {
	json, err := MarshalIndent(1, "\t")
	assert.NoError(t, err)
	assert.Equal(t, "1", json)

	json, err = MarshalIndent("test", "\t")
	assert.NoError(t, err)
	assert.Equal(t, "\"test\"", json)

	json, err = MarshalIndent([]int{1, 2}, "\t")
	assert.NoError(t, err)
	assert.Equal(t, "[\n\t1,\n\t2\n]", json)
}

func TestUnMarshal(t *testing.T) {

}
