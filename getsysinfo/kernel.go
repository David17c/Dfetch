// getsysinfo/getkernel.go
package getsysinfo

import (
	"bufio"
	"os"
)

func Kernel() string {
	file, err := os.Open("/proc/sys/kernel/osrelease")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		firstLine := scanner.Text()
		return firstLine
	}

	if err := scanner.Err(); err != nil {
		return "unknown"
	}
	return ""
}
