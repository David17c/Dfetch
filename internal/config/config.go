package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	EnabledModules []string
	AsciiColor     string
	AccentColor    string
	AsciiSize      string
	CustomAscii    string
}

func configPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "dfetch", "dfetch.conf"), nil
}

func ReadConfig() (*Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := CreateConfigFile(); err != nil {
			return nil, err
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "asciicolor:"):
			cfg.AsciiColor = strings.TrimSpace(
				strings.TrimPrefix(line, "asciicolor:"),
			)

		case strings.HasPrefix(line, "accentcolor:"):
			cfg.AccentColor = strings.TrimSpace(
				strings.TrimPrefix(line, "accentcolor:"),
			)

		case strings.HasPrefix(line, "asciisize:"):
			cfg.AsciiSize = strings.TrimSpace(
				strings.TrimPrefix(line, "asciisize:"),
			)

		case strings.HasPrefix(line, "customascii:"):
			cfg.CustomAscii = strings.TrimSpace(
				strings.TrimPrefix(line, "customascii:"),
			)

		default:
			cfg.EnabledModules = append(cfg.EnabledModules, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func CreateConfigFile() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	configDir := filepath.Dir(path)

	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		var config strings.Builder

		config.WriteString(
			"// Lines starting with `//` are comments and are ignored by Dfetch.\n" +
				"// In the System Information section you can change what info is displayed and in what order.\n\n" +
				"//------------------------\n" +
				"// Colors\n\n" +
				"asciicolor: default\n" +
				"accentcolor: default\n\n" +
				"// Available colors:\n" +
				"// black, red, green, yellow, blue,\n" +
				"// magenta, cyan, white,\n" +
				"// bright_black, bright_red,\n" +
				"// bright_green, bright_yellow,\n" +
				"// bright_blue, bright_magenta,\n" +
				"// bright_cyan, bright_white\n\n" +
				"// ------------------------\n" +
				"// System info modules\n\n" +
				"os\n" +
				"kernel\n" +
				"uptime\n" +
				"shell\n" +
				"de\n" +
				"terminal\n" +
				"cpu\n" +
				"memory\n" +
				"disk\n" +
				"localip\n" +
				"// battery\n" +
				"// time\n" +
				"// date\n\n" +
				"// ------------------------\n" +
				"// Options\n\n" +
				"asciisize: default\n" +
				"// Ascii size can be either 'big', 'default' or 'small'. Default is big.\n\n" +
				"customascii: default\n" +
				"// Set your own custom ascii logo by providing a path to it.",
		)

		if err := os.WriteFile(path, []byte(config.String()), 0600); err != nil {
			return err
		}
	}

	return nil
}
