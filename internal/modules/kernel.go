package modules

import (
	"os"
	"strings"
)

func Kernel() string {
	release, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(release))
}
