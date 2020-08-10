package iox

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLineReader_ReadLine(t *testing.T) {
	inputs := []string{
		"123\n456\n789\n",
		"123\n456\n78",
		"123\n456\n789\n\n",
		"0123\r\n456\r\n789\r\n",
		"0123\r\n456\r\n78",
		"0123\r\n456\r\n789\r\n\r\n",
		"\n",
		"\r\n",
	}
	expects := [][]string{
		{"123", "456", "789"},
		{"123", "456", "78"},
		{"123", "456", "789", ""},
		{"0123", "456", "789"},
		{"0123", "456", "78"},
		{"0123", "456", "789", ""},
		{""},
		{""},
	}

	for i, input := range inputs {
		expect := expects[i]
		r := strings.NewReader(input)
		lr := NewLineReader(r)
		lines, err := lr.ReadAllLines()
		assert.NoError(t, err)
		assert.Equal(t, expect, lines)
	}

	for i, input := range inputs {
		expect := expects[i]
		r := strings.NewReader(input)
		lr := NewLineReaderSize(r, 1)
		lines, err := lr.ReadAllLines()
		assert.NoError(t, err)
		assert.Equal(t, expect, lines)
	}

}
