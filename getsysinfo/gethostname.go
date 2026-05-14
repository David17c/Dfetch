package getsysinfo

import (
	"io"
	"os"
	"strings"
)

func GetHostname() string {
	file, err := os.Open("/etc/hostname")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(content))
}
