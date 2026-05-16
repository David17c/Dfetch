package customization

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ConfigFile() ([]string, error) {

	// Get config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	appConfigDir := filepath.Join(configDir, "Dfetch")

	// Create config directory if missing
	err = os.MkdirAll(appConfigDir, 0700)
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(appConfigDir, "Dfetch.conf")

	// Create file if it doesn't exist
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(configFile)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			_, err = file.WriteString(
				"//Config file for Dfetch. Lines starting with '//' will be ignored. Default settings can be restored by removing this file and running Dfetch.\n\nos\nkernel\ncpu\nmemory\nlocalip\nuptime\n",
			)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Open file for reading
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read lines
	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
