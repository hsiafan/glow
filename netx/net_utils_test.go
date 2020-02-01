package netx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHostType(t *testing.T) {
	assert.Equal(t, DOMAIN, GetHostType("a.com"))
	assert.Equal(t, IPv4, GetHostType("127.0.0.1"))
	assert.Equal(t, IPv6, GetHostType("::127.0.0.1"))
}

func TestJoinHostPort(t *testing.T) {
	assert.Equal(t, "a.com:80", JoinHostPort("a.com", 80))
}

func TestSplitHostPort(t *testing.T) {
	host, port, err := SplitHostPort("a.com:80")
	assert.NoError(t, err)
	assert.Equal(t, "a.com", host)
	assert.Equal(t, uint16(80), port)
	host, port, err = SplitHostPort("a.com")
	assert.Error(t, err)
	host, port, err = SplitHostPort("a.com:xx")
	assert.Error(t, err)
	host, port, err = SplitHostPort("a.com:-1")
	assert.Error(t, err)
	host, port, err = SplitHostPort("a.com:1000000")
	assert.Error(t, err)
}
