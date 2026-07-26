package modules

import (
	"fmt"
	"os"
	"strings"
)

func Kernel() string {
	var kernelType, kernelRelease string

	if b, err := os.ReadFile("/proc/sys/kernel/ostype"); err == nil {
		kernelType = strings.TrimSpace(string(b))
	}

	if b, err := os.ReadFile("/proc/sys/kernel/osrelease"); err == nil {
		kernelRelease = strings.TrimSpace(string(b))
	}

	switch {
	case kernelType != "" && kernelRelease != "":
		return fmt.Sprintf("%s %s", kernelType, kernelRelease)
	case kernelType != "":
		return kernelType
	case kernelRelease != "":
		return kernelRelease
	default:
		return "unknown"
	}
}
