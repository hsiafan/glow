package netx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress(t *testing.T) {
	address, err := ParseAddress("a.com:8080")
	assert.NoError(t, err)
	assert.Equal(t, "a.com", address.Host())
	assert.Equal(t, uint16(8080), address.PortOrDefault(1))
	assert.Equal(t, "a.com:8080", address.String())
	address = address.WithPortIfMissing(8081)
	assert.Equal(t, uint16(8080), address.PortOrDefault(1))
}

func TestAddressWithoutPort(t *testing.T) {
	address, err := ParseAddress("a.com")
	assert.NoError(t, err)
	assert.Equal(t, "a.com", address.Host())
	assert.Equal(t, uint16(1), address.PortOrDefault(1))
	assert.Equal(t, "a.com", address.String())
	address = address.WithPortIfMissing(8080)
	assert.Equal(t, uint16(8080), address.PortOrDefault(1))
	assert.Equal(t, "a.com:8080", address.String())
}
