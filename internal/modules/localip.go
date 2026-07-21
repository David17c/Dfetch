package modules

import (
	"net"
	"strconv"
)

func Local_IP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "unknown"
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
			var (
				ip   net.IP
				mask net.IPMask
			)

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				mask = v.Mask
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil ||
				ip.IsLoopback() ||
				ip.IsLinkLocalUnicast() {
				continue
			}

			prefix := ""
			if mask != nil {
				ones, _ := mask.Size()
				prefix = "/" + strconv.Itoa(ones)
			}

			if ipv4 := ip.To4(); ipv4 != nil {
				return ipv4.String() + prefix
			}

			if ip.To16() != nil {
				return ip.String() + prefix
			}
		}
	}

	return "unknown"
}
