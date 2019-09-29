package iox

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReadAllLines(t *testing.T) {
	str := `123
test
456`
	reader := strings.NewReader(str)
	lines, err := ReadAllLines(reader)
	assert.NoError(t, err)
	assert.Equal(t, strings.Split(str, "\n"), lines)
}
