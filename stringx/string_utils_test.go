package stringx

import "testing"
import "github.com/stretchr/testify/assert"

func TestAppendIfMissing(t *testing.T) {
	assert.Equal(t, "test123", AppendIfMissing("test", "123"))
	assert.Equal(t, "test", AppendIfMissing("test", "st"))
	assert.Equal(t, "test", AppendIfMissing("test", ""))
}

func TestPrependIfMissing(t *testing.T) {
	assert.Equal(t, "123test", PrependIfMissing("test", "123"))
	assert.Equal(t, "test", PrependIfMissing("test", "t"))
	assert.Equal(t, "test", PrependIfMissing("test", ""))
}

func TestSubStringAfter(t *testing.T) {
	assert.Equal(t, "est", SubstringAfter("123test", "t"))
	assert.Equal(t, "", SubstringAfter("test", "st"))
	assert.Equal(t, "test", SubstringAfter("test", ""))
	assert.Equal(t, "", SubstringAfter("test", "xxst"))
	assert.Equal(t, "", SubstringAfter("", "t"))
}

func TestSubStringAfterLast(t *testing.T) {
	assert.Equal(t, "", SubstringAfterLast("123test", "t"))
	assert.Equal(t, "", SubstringAfterLast("test", ""))
	assert.Equal(t, "", SubstringAfterLast("test", "xxst"))
	assert.Equal(t, "x", SubstringAfterLast("testx", "t"))
	assert.Equal(t, "", SubstringAfterLast("", "t"))
}

func TestSubstringBefore(t *testing.T) {
	assert.Equal(t, "123", SubstringBefore("123test", "t"))
	assert.Equal(t, "", SubstringBefore("test", "te"))
	assert.Equal(t, "", SubstringBefore("test", ""))
	assert.Equal(t, "test", SubstringBefore("test", "xxst"))
	assert.Equal(t, "", SubstringBefore("", "t"))
}

func TestSubstringBeforeLast(t *testing.T) {
	assert.Equal(t, "123tes", SubstringBeforeLast("123test", "t"))
	assert.Equal(t, "", SubstringBeforeLast("test", "te"))
	assert.Equal(t, "test", SubstringBeforeLast("test", ""))
	assert.Equal(t, "test", SubstringBeforeLast("test", "xxst"))
	assert.Equal(t, "", SubstringBeforeLast("", "t"))
}

func TestSplitPair(t *testing.T) {
	first, second := SplitPair("123=est", "=")
	assert.Equal(t, "123", first)
	assert.Equal(t, "est", second)
	first, second = SplitPair("123test", "=")
	assert.Equal(t, "123test", first)
	assert.Equal(t, "", second)
	first, second = SplitPair("123test", "")
	assert.Equal(t, "123test", first)
	assert.Equal(t, "", second)
	first, second = SplitPair("", "")
	assert.Equal(t, "", first)
	assert.Equal(t, "", second)
}

func TestPadLeft(t *testing.T) {
	assert.Equal(t, "   123", PadLeft("123", 6, ' '))
	assert.Equal(t, "123", PadLeft("123", 3, ' '))
	assert.Equal(t, "123", PadLeft("123", 0, ' '))
}

func TestPadRight(t *testing.T) {
	assert.Equal(t, "123   ", PadRight("123", 6, ' '))
	assert.Equal(t, "123", PadRight("123", 3, ' '))
	assert.Equal(t, "123", PadRight("123", 0, ' '))
}
