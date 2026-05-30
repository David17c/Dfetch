package sysinfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Memory() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	var memTotal uint64
	var memAvailable uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			memTotal, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "unknown"
			}

		case "MemAvailable:":
			memAvailable, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "unknown"
			}
		}

		if memTotal != 0 && memAvailable != 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "unknown"
	}

	if memTotal == 0 || memAvailable == 0 {
		return "unknown"
	}

	memUsed := memTotal - memAvailable
	usedPercent := float64(memUsed) / float64(memTotal) * 100

	const kbPerMB = 1024
	const kbPerGB = 1024 * 1024
	const kbPerTB = 1024 * 1024 * 1024

	switch {
	case memTotal >= kbPerTB:
		return fmt.Sprintf(
			"%.2f / %.2f TB (%.0f%%)",
			float64(memUsed)/float64(kbPerTB),
			float64(memTotal)/float64(kbPerTB),
			usedPercent,
		)

	case memTotal >= kbPerGB:
		return fmt.Sprintf(
			"%.2f / %.2f GB (%.0f%%)",
			float64(memUsed)/float64(kbPerGB),
			float64(memTotal)/float64(kbPerGB),
			usedPercent,
		)

	case memTotal >= kbPerMB:
		return fmt.Sprintf(
			"%.0f / %.0f MB (%.0f%%)",
			float64(memUsed)/float64(kbPerMB),
			float64(memTotal)/float64(kbPerMB),
			usedPercent,
		)

	default:
		return fmt.Sprintf(
			"%d / %d KB (%.0f%%)",
			memUsed,
			memTotal,
			usedPercent,
		)
	}
}
