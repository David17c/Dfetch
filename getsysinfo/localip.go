package getsysinfo

import (
	"net"
)

func LocalIP() (string, string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "unknown", "unknown"
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			if ipv4 := ip.To4(); ipv4 != nil {
				return ipv4.String(), "IPv4"
			}

			if ip.To16() != nil {
				return ip.String(), "IPv6"
			}
		}
	}
	return "unknown", "unknown"
}
