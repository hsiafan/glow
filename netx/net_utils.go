package netx

import (
	"net"
	"strconv"
)

// JoinHostPort combines host and port into a network address of the
// form "host:port". If host contains a colon, as found in literal
// IPv6 addresses, then JoinHostPort returns "[host]:port".
func JoinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}

// SplitHostPort splits a network address of the form "host:port",
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
// host%zone and port.
func SplitHostPort(str string) (string, int, error) {
	host, portStr, err := net.SplitHostPort(str)
	if err != nil {
		return host, 0, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return host, 0, err
	}
	return host, port, nil
}

type HostType int8

const (
	DOMAIN HostType = 0
	IPv4   HostType = 1
	IPv6   HostType = 2
)

func GetHostType(host string) HostType {
	ip := net.ParseIP(host)
	if ip == nil {
		return DOMAIN
	}
	if ip.To4() != nil {
		return IPv4
	}
	if ip.To16() != nil {
		return IPv6
	}
	//should not happen
	panic("unknown ip type: " + ip.String())
}
