package randx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand_Int32Within(t *testing.T) {
	r := New()
	v := r.Int32Within(1)
	assert.Equal(t, int32(0), v)
	v = r.Int32Within(2)
	assert.True(t, 0 <= v && v < 2)
	v = r.Int32Within(3)
	assert.True(t, 0 <= v && v < 3)
}

func TestRand_Int64Within(t *testing.T) {
	r := New()
	v := r.Int64Within(1)
	assert.Equal(t, int64(0), v)
	v = r.Int64Within(2)
	assert.True(t, 0 <= v && v < 2)
	v = r.Int64Within(3)
	assert.True(t, 0 <= v && v < 3)
}

func TestRand_IntWithin(t *testing.T) {
	r := New()
	v := r.IntWithin(1)
	assert.Equal(t, 0, v)
	v = r.IntWithin(2)
	assert.True(t, 0 <= v && v < 2)
	v = r.IntWithin(3)
	assert.True(t, 0 <= v && v < 3)
}

func TestRand_IntBetween(t *testing.T) {
	r := New()
	v := r.IntBetween(1, 2)
	assert.Equal(t, 1, v)
	v = r.IntBetween(2, 5)
	assert.True(t, 2 <= v && v < 5)
}

func Benchmark_Int32Within(b *testing.B) {
	r := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Int32Within(1 << 30)
	}
}

func Benchmark_Int64Within(b *testing.B) {
	r := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Int64Within(1 << 30)
	}
}
