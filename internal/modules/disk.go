package modules

import (
	"dfetch/internal/format"
	"fmt"
	"syscall"
)

func formatBytesWithUnit(bytes uint64, divisor float64) string {
	return fmt.Sprintf("%.1f", float64(bytes)/divisor)
}

func Disk(formatstring, mount string) string {
	fields := format.Fields(formatstring)

	var stat syscall.Statfs_t

	if mount == "" {
		mount = "/"
	}

	if err := syscall.Statfs(mount, &stat); err != nil {
		return format.Format(formatstring, format.Values{
			"disk": "unknown",
		})
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)

	if total == 0 {
		return format.Format(formatstring, format.Values{
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
		unit = "B"
		divisor = 1
	}

	usedValue := formatBytesWithUnit(used, divisor)
	totalValue := formatBytesWithUnit(total, divisor)

	values := format.Values{}

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

	return format.Format(formatstring, values)
}
