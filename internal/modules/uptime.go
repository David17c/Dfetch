package modules

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

	parts := strings.Fields(string(content))
	if len(parts) < 1 {
		return "unknown"
	}

	secondsFloat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return "unknown"
	}

	total := int(secondsFloat)

	days := total / 86400
	hours := (total % 86400) / 3600
	minutes := (total % 3600) / 60
	seconds := total % 60

	var result string

	if days > 0 {
		result += strconv.Itoa(days) + "d "
	}

	result += strconv.Itoa(hours) + "h " +
		strconv.Itoa(minutes) + "m " +
		strconv.Itoa(seconds) + "s"

	return result
}
