package modules

import (
	"dfetch/internal/format"
	"os"
	"strings"
)

func Host(formatstring string) string {
	info, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_family")
	if err != nil {
		return format.Format(formatstring, format.Values{
			"host": "unknown",
		})
	}

	return format.Format(formatstring, format.Values{
		"host": strings.TrimSpace(string(info)),
	})
}
