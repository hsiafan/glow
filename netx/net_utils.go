package netx

import (
	"fmt"
	"net"
	"strconv"
)

// JoinHostPort combines host and port into a network address of the
// form "host:port". If host contains a colon, as found in literal
// IPv6 addresses, then JoinHostPort returns "[host]:port".
func JoinHostPort(host string, port uint16) string {
	return net.JoinHostPort(host, strconv.Itoa(int(port)))
}

// SplitHostPort splits a network address of the form "host:port",
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or host%zone and port.
// If port part not exists, or is not a valid port, return an error.
func SplitHostPort(str string) (string, uint16, error) {
	host, portStr, err := net.SplitHostPort(str)
	if err != nil {
		return host, 0, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return host, 0, err
	}
	if port < 0 || port >= 65536 {
		return "", 0, fmt.Errorf("invaild port: %v", port)
	}
	return host, uint16(port), nil
}

// HostType is an enum to distinct host type, can be domain, ipv4, or ipv6
type HostType int8

const (
	// Not a host
	ILLEGAL HostType = -1
	// domain host
	DOMAIN HostType = 0
	// ipv4 address host
	IPv4 HostType = 1
	// ipv6 address host
	IPv6 HostType = 2
)

// GetHostType return type for this host
func GetHostType(host string) HostType {
	if len(host) == 0 {
		return ILLEGAL
	}
	ip := net.ParseIP(host)
	if ip == nil {
		//TODO: validate host
		return DOMAIN
	}
	if ip.To4() != nil {
		return IPv4
	}
	if ip.To16() != nil {
		return IPv6
	}
	return ILLEGAL
}
