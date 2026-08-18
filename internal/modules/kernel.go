package modules

import (
	"dfetch/internal/format"
	"os"
	"strings"
)

func Kernel(formatstring string) string {
	fields := format.Fields(formatstring)

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
		return format.Format(formatstring, format.Values{})
	}

	values := format.Values{}

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

	return format.Format(formatstring, values)
}
