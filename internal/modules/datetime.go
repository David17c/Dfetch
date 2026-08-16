package modules

import (
	"dfetch/internal/config"
	"time"
)

func DateTime(format string) string {
	fields := config.Fields(format)

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
		return config.Format(format, config.Values{})
	}

	now := time.Now()

	values := config.Values{}

	if needsTime {
		values["time"] = now.Format("15:04:05")
	}

	if needsDate {
		values["date"] = now.Format("2006-01-02")
	}

	return config.Format(format, values)
}
