package modules

import (
	"os"
	"strings"
)

func readDMI(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func Host() string {
	if v, err := readDMI("/sys/devices/virtual/dmi/id/product_family"); err == nil && v != "" {
		return v
	}

	if v, err := readDMI("/sys/devices/virtual/dmi/id/product_name"); err == nil && v != "" {
		return v
	}

	return "unknown"
}
