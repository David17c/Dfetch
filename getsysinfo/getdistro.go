// getsysinfo/getdistro.go
package getsysinfo

import (
	"bufio"
	"os"
	"strings"
)

func GetDistro() (string, string) {

	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "unknown", "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var name string
	var prettyName string
	var ID string

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "NAME="):
			name = strings.Trim(strings.TrimPrefix(line, "NAME="), `"`)
		case strings.HasPrefix(line, "ID="):
			ID = strings.Trim(strings.TrimPrefix(line, "ID="), `"`)
		case strings.HasPrefix(line, "PRETTY_NAME="):
			prettyName = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), `"`)
		}

		if ID != "" && prettyName != "" {
			return prettyName, ID
		}
	}

	if err := scanner.Err(); err != nil {
		return "unknown", "unknown"
	}

	if ID == "" {
		return "unknown", "unknown"
	}

	if name == "" {
		return ID, ""
	}

	return ID, name
}
