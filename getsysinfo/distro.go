// getsysinfo/getdistro.go
package getsysinfo

import (
	"bufio"
	"os"
	"strings"
)

func Distro() (string, string) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		file, err = os.Open("/usr/lib/os-release")
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		distroName string
		distroID   string
	)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "NAME="):
			if distroName == "" {
				distroName = strings.Trim(strings.TrimPrefix(line, "NAME="), `"`)
			}
		case strings.HasPrefix(line, "PRETTY_NAME="):
			distroName = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), `"`)
		case strings.HasPrefix(line, "ID="):
			distroID = strings.Trim(strings.TrimPrefix(line, "ID="), `"`)
		}
	}

	if err := scanner.Err(); err != nil {
		return "unknown", "unknown"
	}

	if distroID == "" {
		distroID = "unknown"
	}

	if distroName != "" {
		return distroName, distroID
	}
	return "unknown", "unknown"
}
