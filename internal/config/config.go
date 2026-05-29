package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ReadConfig() ([]string, string, string) {
	home, _ := os.UserHomeDir()
	configpath := filepath.Join(home, ".config", "dfetch", "dfetch.conf")

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		err = CreateConfigFile()
		if err != nil {
			return nil, "", ""
		}
	}

	file, err := os.Open(configpath)
	if err != nil {
		return nil, "", ""
	}
	defer file.Close()

	var enabledmodules []string

	var asciicolor string
	var accentcolor string

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
		case strings.HasPrefix(line, "accentcolor:"):
			accentcolor = strings.TrimSpace(strings.TrimPrefix(line, "accentcolor:"))
			continue
		}

		enabledmodules = append(enabledmodules, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, "", ""
	}

	return enabledmodules, asciicolor, accentcolor
}

func CreateConfigFile() error {

	// Create config file / directory if doesnt exist
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appConfigDir := filepath.Join(configDir, "dfetch")

	if err := os.MkdirAll(appConfigDir, 0700); err != nil {
		return err
	}

	configFile := filepath.Join(appConfigDir, "dfetch.conf")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {

		var config strings.Builder

		// Default config file
		config.WriteString(
			"// Lines starting with `//` are comments and are ignored by Dfetch.\n" +
				"// In the System Information section you can change what info is displayed and in what order.\n\n" +
				"// ------------------------\n" +
				"// Color\n" +
				"//------------------------\n\n" +
				"// ASCII color\n" +
				"asciicolor: default\n" +
				"accentcolor: default\n" +
				"// Available colors:\n" +
				"// black, red, green, yellow, blue,\n" +
				"// magenta, cyan, white,\n" +
				"// bright_black, bright_red,\n" +
				"// bright_green, bright_yellow,\n" +
				"// bright_blue, bright_magenta,\n" +
				"// bright_cyan, bright_white\n\n" +
				"// ------------------------\n" +
				"// System Information\n" +
				"// ------------------------\n\n" +
				"os\n" +
				"kernel\n" +
				"uptime\n" +
				"cpu\n" +
				"memory\n" +
				"localip\n" +
				"// shell\n" +
				"// de\n" +
				"// battery\n",
		)

		err = os.WriteFile(configFile, []byte(config.String()), 0600)
		if err != nil {
			return err
		}
	}

	return nil
}
