package modules

import "time"

func DateTime(format string) string {
	now := time.Now()

	switch format {
	case "time":
		return now.Format("15:04:05")
	case "date":
		return now.Format("2006-01-02")
	default:
		return now.Format("2006-01-02 15:04:05")
	}
}
