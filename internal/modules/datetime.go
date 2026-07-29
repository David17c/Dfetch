package modules

import "time"

func DateTime(format string) string {
	if format == "time" {
		return time.Now().Format("15:04:05")
	} else if format == "date" {
		return time.Now().Format("2006-01-02")
	} else {
		return time.Now().Format("2006-01-02 15:04:05")
	}
}

// Time and date are in the same file since there so simple
