package getsysinfo

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Battery() (int, string) {
	batPath := "/sys/class/power_supply/BAT0"

	capacityBytes, err := os.ReadFile(filepath.Join(batPath, "capacity"))
	if err != nil {
		return 0, "unknown"
	}

	capacity, err := strconv.Atoi(strings.TrimSpace(string(capacityBytes)))
	if err != nil {
		return 0, "unknown"
	}

	statusBytes, err := os.ReadFile(filepath.Join(batPath, "status"))
	if err != nil {
		return capacity, "unknown"
	}

	status := strings.TrimSpace(string(statusBytes))

	return capacity, status
}
