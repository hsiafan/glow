package netx

import (
	"strings"
)

// Address contains a host part, and an optional port part.
// This struct is designed to be immutable
type Address struct {
	host    string
	port    uint16
	hasPort bool
}

// ParseAddress parse address from string
func ParseAddress(addressStr string) (*Address, error) {
	if strings.LastIndex(addressStr, ":") == -1 {
		return &Address{host: addressStr}, nil
	}
	host, port, err := SplitHostPort(addressStr)
	if err != nil {
		return nil, err
	}
	return &Address{
		host:    host,
		port:    port,
		hasPort: true,
	}, nil
}

// NewAddress create new address with port
func NewAddress(host string, port uint16) *Address {
	return &Address{host: host, port: port, hasPort: true}
}

// NewAddressWithoutPort create new address without port
func NewAddressWithoutPort(host string) *Address {
	return &Address{host: host}
}

// Host return the host of address
func (a *Address) Host() string {
	return a.host
}

// Port return the port of address, if does not has a port, return false as second value
func (a *Address) Port() (uint16, bool) {
	return a.port, a.hasPort
}

// HasPort return if address has port part
func (a *Address) HasPort() bool {
	return a.hasPort
}

// PortOrDefault return the port of address, if does not has a port, return default port
func (a *Address) PortOrDefault(defaultPort uint16) uint16 {
	if a.hasPort {
		return a.port
	}
	return defaultPort
}

// WithPortIfMissing If port is missing, return a new Address with port specified; else return original address.
func (a *Address) WithPortIfMissing(port uint16) *Address {
	if a.hasPort {
		return a
	}
	return &Address{
		host:    a.host,
		port:    port,
		hasPort: true,
	}
}

// String convert to address string
func (a *Address) String() string {
	if a.hasPort {
		return JoinHostPort(a.host, a.port)
	}
	return a.host
}
