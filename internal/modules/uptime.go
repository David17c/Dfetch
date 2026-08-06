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

	fields := strings.Fields(string(content))
	if len(fields) == 0 {
		return "unknown"
	}

	sec := fields[0]
	if dot := strings.IndexByte(sec, '.'); dot != -1 {
		sec = sec[:dot]
	}

	total, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		return "unknown"
	}

	type unit struct {
		value int64
		long  string
		short string
	}

	units := []unit{
		{total / Century, "century", "c"},
		{0, "year", "y"},
		{0, "month", "mo"},
		{0, "week", "w"},
		{0, "day", "d"},
		{0, "hour", "h"},
		{0, "minute", "m"},
	}

	total %= Century
	units[1].value = total / Year
	total %= Year

	units[2].value = total / Month
	total %= Month

	units[3].value = total / Week
	total %= Week

	units[4].value = total / Day
	total %= Day

	units[5].value = total / Hour
	total %= Hour

	units[6].value = total / Minute

	count := 0
	for _, u := range units {
		if u.value > 0 {
			count++
		}
	}

	if count == 0 {
		return "0 minutes"
	}

	short := count > 3

	var parts []string
	for _, u := range units {
		if u.value == 0 {
			continue
		}

		if short {
			parts = append(parts, strconv.FormatInt(u.value, 10)+u.short)
			continue
		}

		name := u.long
		if u.value != 1 {
			if name == "minute" {
				name = "mins"
			} else {
				name += "s"
			}
		}

		parts = append(parts, strconv.FormatInt(u.value, 10)+" "+name)
	}

	return strings.Join(parts, " ")
}
