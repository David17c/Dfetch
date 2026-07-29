package modules

import (
	"fmt"
	"os"
	"strings"
)

func Kernel(format string) string {
	kernelRelease := readProcFile("/proc/sys/kernel/osrelease")

	if format == "short" {
		if kernelRelease != "" {
			return kernelRelease
		}
		return "unknown"
	}

	kernelType := readProcFile("/proc/sys/kernel/ostype")

	switch {
	case kernelType != "" && kernelRelease != "":
		return fmt.Sprintf("%s %s", kernelType, kernelRelease)
	case kernelRelease != "":
		return kernelRelease
	case kernelType != "":
		return kernelType
	default:
		return "unknown"
	}
}

func readProcFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}
