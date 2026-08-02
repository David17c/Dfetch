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

func Uptime() string {
	content, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}

	parts := strings.Fields(string(content))
	if len(parts) == 0 {
		return "unknown"
	}

	sec := parts[0]
	if dot := strings.IndexByte(sec, '.'); dot != -1 {
		sec = sec[:dot]
	}

	total, err := strconv.ParseInt(sec, 10, 64)

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

	// Just in case
	if centuries > 0 {
		result.WriteString(strconv.FormatInt(centuries, 10))
		result.WriteString("c ")
	}

	if years > 0 {
		result.WriteString(strconv.FormatInt(years, 10))
		result.WriteString("y ")
	}

	if months > 0 {
		result.WriteString(strconv.FormatInt(months, 10))
		result.WriteString("mo ")
	}

	if weeks > 0 {
		result.WriteString(strconv.FormatInt(weeks, 10))
		result.WriteString("w ")
	}

	if days > 0 {
		result.WriteString(strconv.FormatInt(days, 10))
		result.WriteString("d ")
	}

	result.WriteString(strconv.FormatInt(hours, 10))
	result.WriteString("h ")
	result.WriteString(strconv.FormatInt(minutes, 10))
	result.WriteString("m")

	return strings.TrimSpace(result.String())

}
