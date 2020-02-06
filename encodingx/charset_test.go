package encodingx

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func TestEncodeString(t *testing.T) {
	data, err := Encode("测试", simplifiedchinese.GB18030)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0xb2, 0xe2, 0xca, 0xd4}, data)
}

func TestDecodeString(t *testing.T) {
	str, err := Decode([]byte{0xb2, 0xe2, 0xca, 0xd4}, simplifiedchinese.GB18030)
	assert.NoError(t, err)
	assert.Equal(t, "测试", str)
}
