package modules

import (
	"dfetch/internal/format"
	"time"
)

func DateTime(formatstring string) string {
	fields := format.Fields(formatstring)

	needsTime := false
	needsDate := false

	for _, field := range fields {
		switch field {
		case "time":
			needsTime = true

		case "date":
			needsDate = true
		}
	}

	if !needsTime && !needsDate {
		return format.Format(formatstring, format.Values{})
	}

	now := time.Now()

	values := format.Values{}

	if needsTime {
		values["time"] = now.Format("15:04:05")
	}

	if needsDate {
		values["date"] = now.Format("2006-01-02")
	}

	return format.Format(formatstring, values)
}
