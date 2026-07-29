package modules

import (
	"os"
	"strconv"
	"strings"
)

const (
	Minute  int64 = 60
	Hour          = 60 * Minute
	Day           = 24 * Hour
	Week          = 7 * Day
	Month         = 30 * Day
	Year          = 365 * Day
	Century       = 100 * Year
)

func Uptime(format string) string {
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

	total := int64(secondsFloat)

	centuries := total / Century
	total %= Century

	years := total / Year
	total %= Year

	months := total / Month
	total %= Month

	weeks := total / Week
	total %= Week

	days := total / Day
	total %= Day

	hours := total / Hour
	total %= Hour

	minutes := total / Minute

	if centuries == 0 && years == 0 && months == 0 && weeks == 0 {
		var parts []string

		if days > 0 {
			parts = append(parts, strconv.FormatInt(days, 10)+" days")
		}

		parts = append(parts,
			strconv.FormatInt(hours, 10)+" hours",
			strconv.FormatInt(minutes, 10)+" minutes",
		)

		return strings.Join(parts, " ")
	}

	var result strings.Builder

	if centuries > 0 {
		result.WriteString(strconv.FormatInt(centuries, 10) + "c ")
	}

	if years > 0 {
		result.WriteString(strconv.FormatInt(years, 10) + "y ")
	}

	if months > 0 {
		result.WriteString(strconv.FormatInt(months, 10) + "mo ")
	}

	if weeks > 0 {
		result.WriteString(strconv.FormatInt(weeks, 10) + "w ")
	}

	if days > 0 {
		result.WriteString(strconv.FormatInt(days, 10) + "d ")
	}

	result.WriteString(
		strconv.FormatInt(hours, 10) + "h " +
			strconv.FormatInt(minutes, 10) + "m",
	)

	return strings.TrimSpace(result.String())
}
