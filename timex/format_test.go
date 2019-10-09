package timex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	tm := EpochMills(1570612959000)
	assert.Equal(t, "2019-10-09", tm.Format(SimpleDate))
	assert.Equal(t, "2019-10-09 17:22:39", tm.Format(SimpleTime))
	assert.Equal(t, "2019-10-09 17:22:39.000", tm.Format(SimpleTimeMills))
}
