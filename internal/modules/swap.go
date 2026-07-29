package modules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(format string) string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	var swapTotal uint64
	var swapFree uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "SwapTotal:":
			swapTotal, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "unknown"
			}

		case "SwapFree:":
			swapFree, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "unknown"
			}
		}

		if swapTotal != 0 && swapFree != 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "unknown"
	}

	if swapTotal == 0 || swapFree == 0 {
		return "unknown"
	}

	swapUsed := swapTotal - swapFree
	usedPercent := float64(swapUsed) / float64(swapTotal) * 100

	const kbPerMB = 1024
	const kbPerGB = 1024 * 1024
	const kbPerTB = 1024 * 1024 * 1024

	var base string

	switch {
	case swapTotal >= kbPerTB:
		base = fmt.Sprintf(
			"%.2f / %.2f TB",
			float64(swapUsed)/float64(kbPerTB),
			float64(swapTotal)/float64(kbPerTB),
		)
	case swapTotal >= kbPerGB:
		base = fmt.Sprintf(
			"%.2f / %.2f GB",
			float64(swapUsed)/float64(kbPerGB),
			float64(swapTotal)/float64(kbPerGB),
		)
	case swapTotal >= kbPerMB:
		base = fmt.Sprintf(
			"%.0f / %.0f MB",
			float64(swapUsed)/float64(kbPerMB),
			float64(swapTotal)/float64(kbPerMB),
		)
	default:
		base = fmt.Sprintf("%d / %d KB", swapUsed, swapTotal)
	}

	if format != "short" {
		base += fmt.Sprintf(" (%.0f%%)", usedPercent)
	}

	return base
}
