package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ReadConfig() ([]string, string, string, string, string) {
	home, _ := os.UserHomeDir()
	configpath := filepath.Join(home, ".config", "dfetch", "dfetch.conf")

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		err := CreateConfigFile()
		if err != nil {
			return nil, "", "", "", ""
		}
	}

	file, err := os.Open(configpath)
	if err != nil {
		return nil, "", "", "", ""
	}
	defer file.Close()

	var lines []string

	var asciicolor string
	var headercolor string
	var infocolor string
	var labelcolor string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "asciicolor:"):
			asciicolor = strings.TrimSpace(strings.TrimPrefix(line, "asciicolor:"))
			continue
		case strings.HasPrefix(line, "headercolor:"):
			headercolor = strings.TrimSpace(strings.TrimPrefix(line, "headercolor:"))
			continue
		case strings.HasPrefix(line, "infocolor:"):
			infocolor = strings.TrimSpace(strings.TrimPrefix(line, "infocolor:"))
			continue
		case strings.HasPrefix(line, "labelcolor:"):
			labelcolor = strings.TrimSpace(strings.TrimPrefix(line, "labelcolor:"))
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, "", "", "", ""
	}

	return lines, asciicolor, headercolor, infocolor, labelcolor
}
