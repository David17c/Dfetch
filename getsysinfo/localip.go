// getsysinfo/getlocalip.go
package getsysinfo

import (
	"net"
)

func LocalIP() (string, string) {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		return "unknown", "unknown"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP.String()

	var version string
	if localAddr.IP.To4() != nil {
		version = "IPv4"
	} else if localAddr.IP.To16() != nil {
		version = "IPv6"
	} else {
		version = "unknown"
	}

	return ip, version
}
