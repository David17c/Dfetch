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
	if len(parts) == 0 {
		return "unknown"
	}

	secondsFloat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return "unknown"
	}

	total := int(secondsFloat)

	weeks := total / 604800
	days := (total % 604800) / 86400
	hours := (total % 86400) / 3600
	minutes := (total % 3600) / 60
	seconds := total % 60

	var result strings.Builder

	if weeks > 0 {
		result.WriteString(strconv.Itoa(weeks) + "w ")
	}

	if days > 0 {
		result.WriteString(strconv.Itoa(days) + "d ")
	}

	result.WriteString(
		strconv.Itoa(hours) + "h " +
			strconv.Itoa(minutes) + "m " +
			strconv.Itoa(seconds) + "s",
	)

	return result.String()
}
