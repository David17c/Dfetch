package modules

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Battery() string {
	batPath, err := findBattery()
	if err != nil {
		return "unknown"
	}

	present, err := readInt(filepath.Join(batPath, "present"))
	if err != nil || present != 1 {
		return "No battery present"
	}

	capacity, err := readInt(filepath.Join(batPath, "capacity"))
	if err != nil {
		return "No battery present"
	}

	status, err := readString(filepath.Join(batPath, "status"))
	if err != nil {
		status = "unknown"
	}

	if status == "unknown" {
		return fmt.Sprintf("%d%%", capacity)
	}
	return fmt.Sprintf("%d%% (%s)", capacity, status)
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
