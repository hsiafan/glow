package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "", Format(""))
	assert.Equal(t, "1", Format("1"))
	assert.Equal(t, "1,2", Format("{},{}", 1, 2))
	assert.Equal(t, "1,2", Format("{},{}", 1, 2, 3))
	assert.Equal(t, "1,2", Format("{0},{1}", 1, 2))
	assert.Equal(t, "1,1", Format("{0},{0}", 1, 2))
	assert.Panics(t, func() {
		Format("{-0},{0}", 1, 2)
	})
	assert.Panics(t, func() {
		Format("{},{0}", 1, 2)
	})
	assert.Panics(t, func() {
		Format("{0", 1, 2)
	})
	assert.Panics(t, func() {
		Format("0}", 1, 2)
	})
	assert.Panics(t, func() {
		Format("{}}", 1, 2)
	})
	assert.Equal(t, "{1", Format("{{{0}", 1, 2))
	assert.Equal(t, "}}", Format("}}}}", 1, 2))

	assert.Equal(t, "1.00", Format("{:.2f}", 1.0))
	assert.Equal(t, "1", Format("{:<}", 1))
	assert.Equal(t, "1         ", Format("{:<10}", 1))
	assert.Equal(t, "    1     ", Format("{:^10}", 1))
	assert.Equal(t, "0000000001", Format("{:0>10}", 1))
	assert.Equal(t, "0000000144", Format("{:0>10o}", 100))
	assert.Equal(t, "00000000a0", Format("{:0>10x}", 160))
	assert.Equal(t, "00000000A0", Format("{:0>10X}", 160))

	assert.Equal(t, "A0", Format("{:X}", 160))
}
