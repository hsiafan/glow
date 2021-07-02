package httpx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	tp, st := ParseMimeType(MimetypeJson).Unpack()
	assert.Equal(t, "application", tp)
	assert.Equal(t, "json", st)
}
