package httpx

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func Test_parseParam(t *testing.T) {
	assert.Equal(t, &Param{"test", ""}, ParseParam("test"))
	assert.Equal(t, &Param{"", "value"}, ParseParam("=value"))
	assert.Equal(t, &Param{"test", "value"}, ParseParam("test=value"))
}

func TestEncodeQuery(t *testing.T) {
	encoded, err := EncodeQuery("åäö", charmap.ISO8859_1)
	assert.NoError(t, err)
	assert.Equal(t, "%E5%E4%F6", encoded)

	encoded, err = EncodeQuery("测试", simplifiedchinese.GBK)
	assert.NoError(t, err)
	assert.Equal(t, "%B2%E2%CA%D4", encoded)
}
