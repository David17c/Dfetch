package sysinfo

import (
	"bufio"
	"os"
	"strings"
)

func Cpu() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cpu string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "model name	: ") {
			cpu = strings.TrimPrefix(line, "model name	: ")

			return cpu
		}
	}
	return "unknown"
}
