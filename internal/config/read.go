package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ReadConfig() ([]string, string, string) {
	home, _ := os.UserHomeDir()
	configpath := filepath.Join(home, ".config", "Dfetch", "Dfetch.conf")

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		err := CreateConfigFile()
		if err != nil {
			return nil, "", ""
		}
	}

	file, err := os.Open(configpath)
	if err != nil {
		return nil, "", ""
	}
	defer file.Close()

	var lines []string
	var color string
	var ASCII string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "color:"):
			color = strings.TrimSpace(strings.TrimPrefix(line, "color:"))
			continue
		case strings.HasPrefix(line, "ASCII:"):
			ASCII = strings.TrimSpace(strings.TrimPrefix(line, "ASCII:"))
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, "", ""
	}

	return lines, color, ASCII
}
