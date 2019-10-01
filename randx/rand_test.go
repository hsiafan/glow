package randx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRand_Int32Within(t *testing.T) {
	r := New()
	assert.Equal(t, int32(0), r.Int32Within(1))
	v := r.Int32Within(2)
	assert.True(t, 0 <= v && v < 2)
}

func TestRand_Int64Within(t *testing.T) {
	r := New()
	assert.Equal(t, int64(0), r.Int64Within(1))
	v := r.Int64Within(2)
	assert.True(t, 0 <= v && v < 2)
	v = r.Int64Within(3)
	assert.True(t, 0 <= v && v < 3)
}

func TestRand_IntWithin(t *testing.T) {
	r := New()
	assert.Equal(t, 0, r.IntWithin(1))
	v := r.IntWithin(2)
	assert.True(t, 0 <= v && v < 2)
	v = r.IntWithin(3)
	assert.True(t, 0 <= v && v < 3)
}

func TestRand_IntBetween(t *testing.T) {
	r := New()
	assert.Equal(t, 1, r.IntBetween(1, 2))
	v := r.IntBetween(2, 5)
	assert.True(t, 2 <= v && v < 5)
}
