package modules

import (
	"bufio"
	"dfetch/internal/config"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(format string) string {
	fields := config.Fields(format)

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return config.Format(format, config.Values{
			"swap": "unknown",
		})
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
				return config.Format(format, config.Values{
					"swap": "unknown",
				})
			}

		case "SwapFree:":
			swapFree, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return config.Format(format, config.Values{
					"swap": "unknown",
				})
			}
		}

		if swapTotal != 0 && swapFree != 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return config.Format(format, config.Values{
			"swap": "unknown",
		})
	}

	if swapTotal == 0 {
		return config.Format(format, config.Values{
			"swap": "unknown",
		})
	}

	swapUsed := swapTotal - swapFree
	usedPercent := float64(swapUsed) / float64(swapTotal) * 100

	const kbPerMB = 1024
	const kbPerGB = 1024 * 1024
	const kbPerTB = 1024 * 1024 * 1024

	var (
		used  string
		total string
	)

	switch {
	case swapTotal >= kbPerTB:
		used = fmt.Sprintf("%.2f", float64(swapUsed)/float64(kbPerTB))
		total = fmt.Sprintf("%.2f", float64(swapTotal)/float64(kbPerTB))

	case swapTotal >= kbPerGB:
		used = fmt.Sprintf("%.2f", float64(swapUsed)/float64(kbPerGB))
		total = fmt.Sprintf("%.2f", float64(swapTotal)/float64(kbPerGB))

	case swapTotal >= kbPerMB:
		used = fmt.Sprintf("%.0f", float64(swapUsed)/float64(kbPerMB))
		total = fmt.Sprintf("%.0f", float64(swapTotal)/float64(kbPerMB))

	default:
		used = strconv.FormatUint(swapUsed, 10)
		total = strconv.FormatUint(swapTotal, 10)
	}

	var unit string
	switch {
	case swapTotal >= kbPerTB:
		unit = "TB"
	case swapTotal >= kbPerGB:
		unit = "GB"
	case swapTotal >= kbPerMB:
		unit = "MB"
	default:
		unit = "KB"
	}

	swap := used + " / " + total + " " + unit

	values := config.Values{}

	for _, field := range fields {
		switch field {
		case "swap":
			values["swap"] = swap

		case "used":
			values["used"] = used

		case "total":
			values["total"] = total

		case "unit":
			values["unit"] = unit

		case "percent":
			values["percent"] = fmt.Sprintf("%.0f", usedPercent)
		}
	}

	return config.Format(format, values)
}
