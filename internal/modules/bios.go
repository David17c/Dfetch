package modules

import (
	"os"
	"strings"
)

func Bios() string {
	n, err := os.ReadFile("/sys/class/dmi/id/bios_version")
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(n))
}
