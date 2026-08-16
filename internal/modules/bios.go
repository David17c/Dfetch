package modules

import (
	"dfetch/internal/config"
	"os"
	"strings"
)

func Bios(format string) string {
	info, err := os.ReadFile("/sys/class/dmi/id/bios_version")
	if err != nil {
		return config.Format(format, config.Values{
			"bios": "unknown",
		})
	}

	return config.Format(format, config.Values{
		"bios": strings.TrimSpace(string(info)),
	})
}
