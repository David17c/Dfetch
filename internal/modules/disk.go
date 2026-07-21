package modules

import (
	"fmt"
	"syscall"
)

func formatBytesWithUnit(bytes uint64, divisor float64) string {
	return fmt.Sprintf("%.1f", float64(bytes)/divisor)
}

func Disk() string {
	var stat syscall.Statfs_t

	if err := syscall.Statfs("/", &stat); err != nil {
		return "Unknown"
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free
	percent := float64(used) / float64(total) * 100

	const (
		KB = 1000
		MB = KB * 1000
		GB = MB * 1000
		TB = GB * 1000
	)

	var unit string
	var divisor float64

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
		return fmt.Sprintf("%d / %d B", used, total)
	}

	return fmt.Sprintf("%s / %s %s (%.0f%%)",
		formatBytesWithUnit(used, divisor),
		formatBytesWithUnit(total, divisor),
		unit,
		percent,
	)
}
