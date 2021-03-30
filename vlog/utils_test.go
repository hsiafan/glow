package vlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSize(t *testing.T) {
	v, _ := parseSize("1000")
	assert.Equal(t, int64(1000), v)
	v, _ = parseSize("100k")
	assert.Equal(t, int64(102400), v)
	v, _ = parseSize("100m")
	assert.Equal(t, int64(104857600), v)
	v, _ = parseSize("100G")
	assert.Equal(t, int64(107374182400), v)
	v, _ = parseSize("100T")
	assert.Equal(t, int64(109951162777600), v)
}
