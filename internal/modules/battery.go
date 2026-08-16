package modules

import (
	"dfetch/internal/config"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Battery(format string) string {
	fields := config.Fields(format)

	needsPercent := false
	needsStatus := false

	for _, field := range fields {
		switch field {
		case "percent":
			needsPercent = true

		case "status":
			needsStatus = true
		}
	}

	if !needsPercent && !needsStatus {
		return config.Format(format, config.Values{})
	}

	batPath, err := findBattery()
	if err != nil {
		return config.Format(format, config.Values{
			"percent": "unknown",
			"status":  "unknown",
		})
	}

	present, err := readInt(filepath.Join(batPath, "present"))
	if err != nil || present != 1 {
		return config.Format(format, config.Values{
			"percent": "unknown",
			"status":  "No battery present",
		})
	}

	capacity, err := readInt(filepath.Join(batPath, "capacity"))
	if err != nil {
		return config.Format(format, config.Values{
			"percent": "unknown",
			"status":  "unknown",
		})
	}

	values := config.Values{
		"percent": fmt.Sprintf("%d%%", capacity),
	}

	if needsStatus {
		status, err := readString(filepath.Join(batPath, "status"))
		if err != nil || status == "" {
			status = "unknown"
		}

		values["status"] = status
	}

	return config.Format(format, values)
}

func findBattery() (string, error) {
	const base = "/sys/class/power_supply"

	entries, err := os.ReadDir(base)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		typePath := filepath.Join(base, entry.Name(), "type")

		b, err := os.ReadFile(typePath)
		if err != nil {
			continue
		}

		if strings.TrimSpace(string(b)) == "Battery" {
			return filepath.Join(base, entry.Name()), nil
		}
	}

	return "", os.ErrNotExist
}

func readInt(path string) (int, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(string(b)))
}

func readString(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
