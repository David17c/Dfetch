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

	total := int(totalSeconds)

	days := total / 86400
	hours := (total % 86400) / 3600
	minutes := (total % 3600) / 60
	seconds := total % 60

	result := ""

	if days >= 1 {
		result += strconv.Itoa(days) + "d "
	}

	result += strconv.Itoa(hours) + "h " +
		strconv.Itoa(minutes) + "m " +
		strconv.Itoa(seconds) + "s"

	return result
}
