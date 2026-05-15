package getsysinfo

import (
	"os"
	"strconv"
	"strings"
)

func Uptime() string {
	content, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}

	parts := strings.SplitN(string(content), " ", 2)
	if len(parts) == 0 {
		return "unknown"
	}

	totalSeconds, _ := strconv.ParseFloat(parts[0], 64)

	hours := int(totalSeconds) / 3600
	minutes := (int(totalSeconds) % 3600) / 60
	seconds := int(totalSeconds) % 60

	return strconv.Itoa(hours) + "h " +
		strconv.Itoa(minutes) + "m " +
		strconv.Itoa(seconds) + "s"
}
