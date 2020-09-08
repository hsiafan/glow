package hashx

import (
	"crypto/md5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5Bytes(t *testing.T) {
	assert.Equal(t, "202cb962ac59075b964b07152d234b70", HashBytes([]byte("123"), md5.New()).ToHex())
}
