package modules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	var SwapTotal uint64
	var SwapFree uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "SwapTotal:":
			SwapTotal, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "unknown"
			}

		case "SwapFree:":
			SwapFree, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "unknown"
			}
		}

		if SwapTotal != 0 && SwapFree != 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "unknown"
	}

	if SwapTotal == 0 || SwapFree == 0 {
		return "unknown"
	}

	SwapUsed := SwapTotal - SwapFree
	usedPercent := float64(SwapUsed) / float64(SwapTotal) * 100

	const kbPerMB = 1024
	const kbPerGB = 1024 * 1024
	const kbPerTB = 1024 * 1024 * 1024

	switch {
	case SwapTotal >= kbPerTB:
		return fmt.Sprintf(
			"%.2f / %.2f TB (%.0f%%)",
			float64(SwapUsed)/float64(kbPerTB),
			float64(SwapTotal)/float64(kbPerTB),
			usedPercent,
		)

	case SwapTotal >= kbPerGB:
		return fmt.Sprintf(
			"%.2f / %.2f GB (%.0f%%)",
			float64(SwapUsed)/float64(kbPerGB),
			float64(SwapTotal)/float64(kbPerGB),
			usedPercent,
		)

	case SwapTotal >= kbPerMB:
		return fmt.Sprintf(
			"%.0f / %.0f MB (%.0f%%)",
			float64(SwapUsed)/float64(kbPerMB),
			float64(SwapTotal)/float64(kbPerMB),
			usedPercent,
		)

	default:
		return fmt.Sprintf(
			"%d / %d KB (%.0f%%)",
			SwapUsed,
			SwapTotal,
			usedPercent,
		)
	}
}
