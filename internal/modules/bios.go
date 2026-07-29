package modules

import (
	"fmt"
	"os"
	"strings"
)

func Bios(format string) string {
	n, err := os.ReadFile("/sys/class/dmi/id/bios_version")
	if err != nil {
		return "unknown"
	}

	bios_version := strings.TrimSpace(string(n))

	if format == "short" {
		return bios_version
	}

	d, err := os.ReadFile("/sys/class/dmi/id/bios_date")
	if err != nil {
		return bios_version
	}

	bios_release_date := strings.TrimSpace(string(d))

	return fmt.Sprintf("%s %s", bios_version, bios_release_date)

}
