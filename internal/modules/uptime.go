package modules

import (
	"dfetch/internal/config"
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

type uptimeValues struct {
	centuries int64
	years     int64
	months    int64
	weeks     int64
	days      int64
	hours     int64
	minutes   int64
}

func Uptime(format string) string {
	fields := config.Fields(format)

	needsUptime := false
	needsUnits := false

	for _, field := range fields {
		switch field {
		case "uptime":
			needsUptime = true

		case "centuries",
			"years",
			"months",
			"weeks",
			"days",
			"hours",
			"minutes":
			needsUnits = true
		}
	}

	if !needsUptime && !needsUnits {
		return config.Format(format, config.Values{})
	}

	values := readUptime()

	if values == nil {
		return config.Format(format, config.Values{
			"uptime": "unknown",
		})
	}

	result := config.Values{}

	for _, field := range fields {
		switch field {
		case "uptime":
			result["uptime"] = formatUptime(*values)

		case "centuries":
			result["centuries"] = strconv.FormatInt(values.centuries, 10)

		case "years":
			result["years"] = strconv.FormatInt(values.years, 10)

		case "months":
			result["months"] = strconv.FormatInt(values.months, 10)

		case "weeks":
			result["weeks"] = strconv.FormatInt(values.weeks, 10)

		case "days":
			result["days"] = strconv.FormatInt(values.days, 10)

		case "hours":
			result["hours"] = strconv.FormatInt(values.hours, 10)

		case "minutes":
			result["minutes"] = strconv.FormatInt(values.minutes, 10)
		}
	}

	return config.Format(format, result)
}

func readUptime() *uptimeValues {
	content, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return nil
	}

	fields := strings.Fields(string(content))
	if len(fields) == 0 {
		return nil
	}

	sec := fields[0]

	if dot := strings.IndexByte(sec, '.'); dot != -1 {
		sec = sec[:dot]
	}

	total, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		return nil
	}

	result := &uptimeValues{}

	result.centuries = total / Century
	total %= Century

	result.years = total / Year
	total %= Year

	result.months = total / Month
	total %= Month

	result.weeks = total / Week
	total %= Week

	result.days = total / Day
	total %= Day

	result.hours = total / Hour
	total %= Hour

	result.minutes = total / Minute

	return result
}

func formatUptime(u uptimeValues) string {
	type unit struct {
		value int64
		long  string
		short string
	}

	units := []unit{
		{u.centuries, "century", "c"},
		{u.years, "year", "y"},
		{u.months, "month", "mo"},
		{u.weeks, "week", "w"},
		{u.days, "day", "d"},
		{u.hours, "hour", "h"},
		{u.minutes, "minute", "m"},
	}

	count := 0

	for _, unit := range units {
		if unit.value > 0 {
			count++
		}
	}

	if count == 0 {
		return "0 minutes"
	}

	short := count > 3

	var parts []string

	for _, unit := range units {
		if unit.value == 0 {
			continue
		}

		value := strconv.FormatInt(unit.value, 10)

		if short {
			parts = append(parts, value+unit.short)
			continue
		}

		name := unit.long

		if unit.value != 1 {
			if name == "minute" {
				name = "mins"
			} else {
				name += "s"
			}
		}

		parts = append(parts, value+" "+name)
	}

	return strings.Join(parts, ", ")
}
