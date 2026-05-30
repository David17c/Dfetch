package sysinfo

import (
	"os"
	"strings"
)

func Kernel() string {
	b, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "unknown"
	}

	s := strings.TrimSpace(string(b))
	if s == "" {
		return "unknown"
	}

	return s
}
