package intx

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestParseInt(t *testing.T) {
	assert.Equal(t, 100000000, SafeParse("100000000", -1))
}

func TestFormatInt(t *testing.T) {
	assert.Equal(t, "100", Format(100))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 0, Abs(0))
	assert.Equal(t, 1, Abs(-1))
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, MaxInt, Abs(MaxInt))
	assert.Equal(t, MinInt, Abs(MinInt))
}

func TestAbs64(t *testing.T) {
	assert.Equal(t, int64(0), Abs64(0))
	assert.Equal(t, int64(1), Abs64(-1))
	assert.Equal(t, int64(1), Abs64(1))
	assert.Equal(t, int64(math.MaxInt64), Abs64(math.MaxInt64))
	assert.Equal(t, int64(math.MinInt64), Abs64(math.MinInt64))
}

func TestAbs32(t *testing.T) {
	assert.Equal(t, int32(0), Abs32(0))
	assert.Equal(t, int32(1), Abs32(-1))
	assert.Equal(t, int32(1), Abs32(1))
	assert.Equal(t, int32(math.MaxInt32), Abs32(math.MaxInt32))
	assert.Equal(t, int32(math.MinInt32), Abs32(math.MinInt32))
}
