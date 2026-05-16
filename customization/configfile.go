package customization

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ConfigFile() ([]string, string, error) {

	// Get config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, "", err
	}

	appConfigDir := filepath.Join(configDir, "Dfetch")

	// Create config directory if missing
	err = os.MkdirAll(appConfigDir, 0700)
	if err != nil {
		return nil, "", err
	}

	configFile := filepath.Join(appConfigDir, "Dfetch.conf")

	// Create file if it doesn't exist
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(configFile)
			if err != nil {
				return nil, "", err
			}
			defer file.Close()

			_, err = file.WriteString(
				"// Config file for Dfetch\n" +
					"// Lines starting with '//' will be ignored\n" +
					"// Default settings can be restored by deleting this file\n\n" +
					"// color: blue\n" +
					"os\n" +
					"kernel\n" +
					"cpu\n" +
					"memory\n" +
					"localip\n" +
					"uptime\n" +
					"//battery\n",
			)
			if err != nil {
				return nil, "", err
			}
		} else {
			return nil, "", err
		}
	}

	// Open file for reading
	file, err := os.Open(configFile)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	var lines []string
	var color string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		// if line starts with "color:" store it
		if strings.HasPrefix(line, "color:") {
			color = strings.TrimSpace(strings.TrimPrefix(line, "color:"))
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, "", err
	}

	return lines, color, nil
}
