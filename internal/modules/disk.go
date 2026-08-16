package modules

import (
	"dfetch/internal/config"
	"fmt"
	"syscall"
)

func formatBytesWithUnit(bytes uint64, divisor float64) string {
	return fmt.Sprintf("%.1f", float64(bytes)/divisor)
}

func Disk(format, mount string) string {
	fields := config.Fields(format)

	var stat syscall.Statfs_t

	if mount == "" {
		mount = "/"
	}

	if err := syscall.Statfs(mount, &stat); err != nil {
		return config.Format(format, config.Values{
			"disk": "unknown",
		})
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)

	if total == 0 {
		return config.Format(format, config.Values{
			"disk": "unknown",
		})
	}

	used := total - free
	percent := float64(used) / float64(total) * 100

	const (
		KB = 1000
		MB = KB * 1000
		GB = MB * 1000
		TB = GB * 1000
	)

	var (
		unit    string
		divisor float64
	)

	switch {
	case total >= TB:
		unit = "TB"
		divisor = TB

	case total >= GB:
		unit = "GB"
		divisor = GB

	case total >= MB:
		unit = "MB"
		divisor = MB

	case total >= KB:
		unit = "KB"
		divisor = KB

	default:
		values := config.Values{}

		for _, field := range fields {
			switch field {
			case "disk":
				values["disk"] = fmt.Sprintf("%d / %d B", used, total)

			case "used":
				values["used"] = fmt.Sprintf("%d", used)

			case "total":
				values["total"] = fmt.Sprintf("%d", total)

			case "unit":
				values["unit"] = "B"

			case "percent":
				values["percent"] = fmt.Sprintf("%.0f", percent)
			}
		}

		return config.Format(format, values)
	}

	usedValue := formatBytesWithUnit(used, divisor)
	totalValue := formatBytesWithUnit(total, divisor)

	values := config.Values{}

	for _, field := range fields {
		switch field {
		case "disk":
			values["disk"] = fmt.Sprintf(
				"%s / %s %s",
				usedValue,
				totalValue,
				unit,
			)

		case "used":
			values["used"] = usedValue

		case "total":
			values["total"] = totalValue

		case "unit":
			values["unit"] = unit

		case "percent":
			values["percent"] = fmt.Sprintf("%.0f", percent)
		}
	}

	return config.Format(format, values)
}
