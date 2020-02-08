package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestFormat(t *testing.T) {
	tm := EpochMills(1570612959000)
	assert.Equal(t, "2019-10-09", tm.Format(SimpleDate))
	assert.Equal(t, "2019-10-09 17:22:39", tm.Format(SimpleTime))
	assert.Equal(t, "2019-10-09 17:22:39.000", tm.Format(SimpleTimeMills))
}
