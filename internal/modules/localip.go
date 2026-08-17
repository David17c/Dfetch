package modules

import (
	"dfetch/internal/config"
	"net"
	"strconv"
)

func LocalIP(format string) string {
	fields := config.Fields(format)

	needsIP := false
	needsPrefix := false
	needsAddress := false

	for _, field := range fields {
		switch field {
		case "ip":
			needsIP = true

		case "prefix":
			needsPrefix = true

		case "address":
			needsAddress = true
		}
	}

	if !needsIP && !needsPrefix && !needsAddress {
		return config.Format(format, config.Values{})
	}

	ip, prefix := detectLocalIP()

	if ip == "" {
		return config.Format(format, config.Values{
			"ip":      "unknown",
			"prefix":  "unknown",
			"address": "unknown",
		})
	}

	values := config.Values{}

	for _, field := range fields {
		switch field {
		case "ip":
			values["ip"] = ip

		case "prefix":
			values["prefix"] = prefix

		case "address":
			if prefix != "" {
				values["address"] = ip + "/" + prefix
			} else {
				values["address"] = ip
			}
		}
	}

	return config.Format(format, values)
}

func detectLocalIP() (string, string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", ""
	}

	var ipv6IP, ipv6Prefix string

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
				prefix = strconv.Itoa(ones)
			}

			if ipv4 := ip.To4(); ipv4 != nil {
				return ipv4.String(), prefix
			}

			if ip.To16() != nil && ipv6IP == "" {
				ipv6IP = ip.String()
				ipv6Prefix = prefix
			}
		}
	}

	if ipv6IP != "" {
		return ipv6IP, ipv6Prefix
	}

	return "", ""
}
