package modules

import (
	"dfetch/internal/format"
	"os"
	"strings"
)

func Bios(formatstring string) string {
	info, err := os.ReadFile("/sys/class/dmi/id/bios_version")
	if err != nil {
		return format.Format(formatstring, format.Values{
			"bios": "unknown",
		})
	}

	return format.Format(formatstring, format.Values{
		"bios": strings.TrimSpace(string(info)),
	})
}
