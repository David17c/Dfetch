package modules

import (
	"dfetch/internal/config"
	"os"
	"strings"
)

func Host(format string) string {
	info, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_family")
	if err != nil {
		return config.Format(format, config.Values{
			"host": "unknown",
		})
	}

	return config.Format(format, config.Values{
		"host": strings.TrimSpace(string(info)),
	})
}
