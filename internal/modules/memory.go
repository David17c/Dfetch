package modules

import (
	"bufio"
	"dfetch/internal/config"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Memory(format string) string {
	fields := config.Fields(format)

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return config.Format(format, config.Values{
			"memory": "unknown",
		})
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
				return config.Format(format, config.Values{
					"memory": "unknown",
				})
			}

		case "MemAvailable:":
			memAvailable, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return config.Format(format, config.Values{
					"memory": "unknown",
				})
			}
		}

		if memTotal != 0 && memAvailable != 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return config.Format(format, config.Values{
			"memory": "unknown",
		})
	}

	if memTotal == 0 || memAvailable == 0 {
		return config.Format(format, config.Values{
			"memory": "unknown",
		})
	}

	memUsed := memTotal - memAvailable
	usedPercent := float64(memUsed) / float64(memTotal) * 100

	const kbPerMB = 1024
	const kbPerGB = 1024 * 1024
	const kbPerTB = 1024 * 1024 * 1024

	var (
		used  string
		total string
		unit  string
	)

	switch {
	case memTotal >= kbPerTB:
		unit = "TB"
		used = fmt.Sprintf("%.2f", float64(memUsed)/float64(kbPerTB))
		total = fmt.Sprintf("%.2f", float64(memTotal)/float64(kbPerTB))

	case memTotal >= kbPerGB:
		unit = "GB"
		used = fmt.Sprintf("%.2f", float64(memUsed)/float64(kbPerGB))
		total = fmt.Sprintf("%.2f", float64(memTotal)/float64(kbPerGB))

	case memTotal >= kbPerMB:
		unit = "MB"
		used = fmt.Sprintf("%.0f", float64(memUsed)/float64(kbPerMB))
		total = fmt.Sprintf("%.0f", float64(memTotal)/float64(kbPerMB))

	default:
		unit = "KB"
		used = strconv.FormatUint(memUsed, 10)
		total = strconv.FormatUint(memTotal, 10)
	}

	memory := used + " / " + total + " " + unit

	values := config.Values{}

	for _, field := range fields {
		switch field {
		case "memory":
			values["memory"] = memory

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
