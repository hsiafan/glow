package timex

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEpochMills(t *testing.T) {
	tm := EpochMills(1570612959000)
	y, m, d := tm.Date()
	assert.Equal(t, 2019, y)
	assert.Equal(t, time.Month(10), m)
	assert.Equal(t, 9, d)
}

func TestEpoch(t *testing.T) {
	tm := Epoch(1570612959)
	y, m, d := tm.Date()
	assert.Equal(t, 2019, y)
	assert.Equal(t, time.Month(10), m)
	assert.Equal(t, 9, d)
}
