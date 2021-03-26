package stringx

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type testFace interface {
}

type testStruct struct {
	TestField string
}

func TestValueOf(t *testing.T) {
	assert.Equal(t, "1", ValueOf(1))
	var i *int
	assert.Equal(t, "<nil>", ValueOf(i))
	var ii = 1
	i = &ii
	assert.True(t, strings.HasPrefix(ValueOf(i), "0x"))
	assert.Equal(t, "<nil>", ValueOf(nil))

	assert.Equal(t, "10", ValueOf(float32(10)))

	var m map[string]string = nil
	var f testFace = m
	assert.Equal(t, "map[]", ValueOf(f))

	var ts *testStruct = nil
	f = ts
	assert.Equal(t, "<nil>", ValueOf(f))
	assert.Equal(t, "<nil>", ValueOf(ts))
	ts = &testStruct{TestField: "1"}
	assert.Equal(t, "&{1}", ValueOf(ts))

	c := complex(10, 20)
	assert.Equal(t, "(10+20i)", ValueOf(c))

	err := errors.New("test error")
	assert.Equal(t, "test error", ValueOf(err))
}

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
	assert.Equal(t, "est", SubstringAfter("123test", "t").Value)
	assert.Equal(t, "", SubstringAfter("test", "st").Value)
	assert.Equal(t, "test", SubstringAfter("test", "").Value)
	assert.Equal(t, "", SubstringAfter("test", "xxst").Value)
	assert.Equal(t, "", SubstringAfter("", "t").Value)
}

func TestSubStringAfterLast(t *testing.T) {
	assert.Equal(t, "", SubstringAfterLast("123test", "t").Value)
	assert.Equal(t, "", SubstringAfterLast("test", "").Value)
	assert.Equal(t, "", SubstringAfterLast("test", "xxst").Value)
	assert.Equal(t, "x", SubstringAfterLast("testx", "t").Value)
	assert.Equal(t, "", SubstringAfterLast("", "t").Value)
}

func TestSubstringBefore(t *testing.T) {
	assert.Equal(t, "123", SubstringBefore("123test", "t").Value)
	assert.Equal(t, "", SubstringBefore("test", "te").Value)
	assert.Equal(t, "", SubstringBefore("test", "").Value)
	assert.Equal(t, false, SubstringBefore("test", "xxst").Exists)
	assert.Equal(t, "", SubstringBefore("", "t").Value)
}

func TestSubstringBeforeLast(t *testing.T) {
	assert.Equal(t, "123tes", SubstringBeforeLast("123test", "t").Value)
	assert.Equal(t, "", SubstringBeforeLast("test", "te").Value)
	assert.Equal(t, "test", SubstringBeforeLast("test", "").Value)
	assert.Equal(t, false, SubstringBeforeLast("test", "xxst").Exists)
	assert.Equal(t, "", SubstringBeforeLast("", "t").Value)
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

func TestAppendIfNotEmpty(t *testing.T) {
	assert.Equal(t, "123 ", AppendIfNotEmpty("123", " "))
	assert.Equal(t, "123", AppendIfNotEmpty("123", ""))
	assert.Equal(t, "", AppendIfNotEmpty("", " "))
}

func TestPrependIfNotEmpty(t *testing.T) {
	assert.Equal(t, " 123", PrependIfNotEmpty("123", " "))
	assert.Equal(t, "123", PrependIfNotEmpty("123", ""))
	assert.Equal(t, "", PrependIfNotEmpty("", " "))
}

func TestCapitalize(t *testing.T) {
	assert.Equal(t, "", Capitalize(""))
	assert.Equal(t, "1", Capitalize("1"))
	assert.Equal(t, "A", Capitalize("a"))
	assert.Equal(t, "Aa", Capitalize("Aa"))
	assert.Equal(t, "Aa", Capitalize("aa"))
}

func TestDeCapitalize(t *testing.T) {
	assert.Equal(t, "", DeCapitalize(""))
	assert.Equal(t, "1", DeCapitalize("1"))
	assert.Equal(t, "a", DeCapitalize("A"))
	assert.Equal(t, "aa", DeCapitalize("Aa"))
	assert.Equal(t, "aa", DeCapitalize("aa"))
}

func TestSnakeToCamel(t *testing.T) {
	assert.Equal(t, "", SnakeToCamel("", true))
	assert.Equal(t, "", SnakeToCamel("_", true))
	assert.Equal(t, "Test", SnakeToCamel("_test", true))
	assert.Equal(t, "Test", SnakeToCamel("test", true))
	assert.Equal(t, "test", SnakeToCamel("_test", false))
	assert.Equal(t, "test", SnakeToCamel("test", false))
	assert.Equal(t, "testLive", SnakeToCamel("test_live", false))
	assert.Equal(t, "testLive", SnakeToCamel("test__live", false))
	assert.Equal(t, "TestLive", SnakeToCamel("test_live", true))
}

func TestCamelToSnake(t *testing.T) {
	assert.Equal(t, "", CamelToSnake(""))
	assert.Equal(t, "test", CamelToSnake("test"))
	assert.Equal(t, "test_live", CamelToSnake("testLive"))
	assert.Equal(t, "test_live", CamelToSnake("TestLive"))
	assert.Equal(t, "test_http", CamelToSnake("TestHTTP"))
	assert.Equal(t, "test_http_port", CamelToSnake("TestHTTPPort"))
	assert.Equal(t, "http_port", CamelToSnake("HTTPPort"))
	assert.Equal(t, "http", CamelToSnake("HTTP"))
}

func TestCopy(t *testing.T) {
	var s = "test123"
	s2 := s[:4]
	s3 := Copy(s2)
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s3))
	assert.Equal(t, 4, sh.Len)
}

func TestSlice(t *testing.T) {
	assert.Equal(t, "test", Slice("test123", 0, -3))
	assert.Equal(t, "123", SliceToEnd("test123", -3))
}

func TestFirstNonEmpty(t *testing.T) {
	assert.Equal(t, "", FirstNonEmpty("", ""))
	assert.Equal(t, "1", FirstNonEmpty("1", "2"))
	assert.Equal(t, "2", FirstNonEmpty("", "2"))
}
