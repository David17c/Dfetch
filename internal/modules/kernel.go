package modules

import (
	"dfetch/internal/config"
	"os"
	"strings"
)

func Kernel(format string) string {
	fields := config.Fields(format)

	needsVersion := false
	needsType := false

	for _, field := range fields {
		switch field {
		case "version":
			needsVersion = true

		case "type":
			needsType = true
		}
	}

	if !needsVersion && !needsType {
		return config.Format(format, config.Values{})
	}

	values := config.Values{}

	if needsVersion {
		version, err := os.ReadFile("/proc/sys/kernel/osrelease")
		if err != nil {
			values["version"] = "unknown"
		} else {
			values["version"] = strings.TrimSpace(string(version))
		}
	}

	if needsType {
		values["type"] = "Linux"
	}

	return config.Format(format, values)
}
