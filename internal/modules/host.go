package modules

import (
	"os"
	"strings"
)

func Host() string {
	host, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_family")
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(host))

}
