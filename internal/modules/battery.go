package modules

import (
	"dfetch/internal/format"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Battery(formatstring string) string {
	fields := format.Fields(formatstring)

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
		return format.Format(formatstring, format.Values{})
	}

	batPath, err := findBattery()
	if err != nil {
		return format.Format(formatstring, format.Values{
			"percent": "unknown",
			"status":  "unknown",
		})
	}

	// Assume the battery is present unless the "present" file
	// explicitly says otherwise.
	presentPath := filepath.Join(batPath, "present")
	if _, err := os.Stat(presentPath); err == nil {
		present, err := readInt(presentPath)
		if err != nil || present != 1 {
			return format.Format(formatstring, format.Values{
				"percent": "unknown",
				"status":  "No battery present",
			})
		}
	}

	capacity, err := readInt(filepath.Join(batPath, "capacity"))
	if err != nil {
		return format.Format(formatstring, format.Values{
			"percent": "unknown",
			"status":  "unknown",
		})
	}

	values := format.Values{
		"percent": fmt.Sprintf("%d%%", capacity),
	}

	if needsStatus {
		status, err := readString(filepath.Join(batPath, "status"))
		if err != nil || status == "" {
			status = "unknown"
		}

		values["status"] = status
	}

	return format.Format(formatstring, values)
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
