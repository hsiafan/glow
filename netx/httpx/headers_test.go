package httpx

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {

	d, err := ParseDateHeader("Wed, 21 Oct 2015 07:28:00 GMT")
	assert.NoError(t, err)
	_ = d
	//ed, _ := time.Parse(time.RFC1123, "2015-10-21 15:28:00 +0800")
	//assert.Equal(t, ed, d)
}

func TestParseDate(t *testing.T) {
	s := FormatDateHeader(time.Now())
	_ = s
	//fmt.Println(s)
}
