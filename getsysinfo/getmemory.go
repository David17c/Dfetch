package getsysinfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetMem() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var memTotal int
	var memAvailable int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			val, err := strconv.Atoi(fields[1])
			if err != nil {
				return ""
			}
			memTotal = val

		case "MemAvailable:":
			val, err := strconv.Atoi(fields[1])
			if err != nil {
				return ""
			}
			memAvailable = val
		}

		if memTotal != 0 && memAvailable != 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	memUsed := memTotal - memAvailable
	percentageUsed := float64(memUsed) / float64(memTotal) * 100

	switch {
	case memUsed < 1024:
		return fmt.Sprintf("%d / %d KB (%.0f%%)", memUsed, memTotal, percentageUsed)

	case memUsed < 1024*1024:
		return fmt.Sprintf(
			"%.2f / %.2f MB (%.0f%%)",
			float64(memUsed)/1024,
			float64(memTotal)/1024,
			percentageUsed,
		)

	default:
		return fmt.Sprintf(
			"%.2f / %.2f GB (%.0f%%)",
			float64(memUsed)/(1024*1024),
			float64(memTotal)/(1024*1024),
			percentageUsed,
		)
	}
}
