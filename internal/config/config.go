package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	EnabledModules []string

	AsciiColor  string
	AccentColor string
	AsciiSize   string
	CustomAscii string
}

func configPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "dfetch", "dfetch.conf"), nil
}

func ReadConfig() (*Config, error) {
	var inModules bool

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

	// Give option variables default values
	cfg := &Config{
		AsciiSize:   "default",
		AsciiColor:  "default",
		AccentColor: "default",
		CustomAscii: "default",
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip comments
		if idx := strings.Index(line, "//"); idx != -1 {
			line = strings.TrimSpace(line[:idx])
		}

		// Skip empty lines
		if line == "" {
			continue
		}

		// Detects start of modules block
		if strings.EqualFold(line, "modules {") {
			inModules = true
			continue
		}

		// Collects contents of modules block and the detects the end of this block
		if inModules {
			if line == "}" {
				inModules = false
				continue
			}

			cfg.EnabledModules = append(cfg.EnabledModules, strings.ToLower(line))
			continue
		}

		if idx := strings.Index(line, ":"); idx != -1 {
			key := strings.ToLower(strings.TrimSpace(line[:idx]))
			value := strings.ToLower(strings.TrimSpace(line[idx+1:]))

			switch key {
			case "asciicolor":
				cfg.AsciiColor = value
			case "accentcolor":
				cfg.AccentColor = value
			case "asciisize":
				if value != "default" && value != "small" && value != "big" {
					return nil, fmt.Errorf("invalid asciisize in config")
				}
				cfg.AsciiSize = value
			case "customascii":
				cfg.CustomAscii = value
			}

			continue
		}
	}

	// Error if scanner fails
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Error if modules block was never closed
	if inModules {
		return nil, fmt.Errorf("unterminated block in config")
	}

	return cfg, nil
}

func CreateConfigFile() error {
	// Detects default config path
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
				"// In the modules section you can change what info is displayed and in what order.\n\n" +
				"modules {\n" +
				"	userinfo\n" +
				"	os\n" +
				"	kernel\n" +
				"	uptime\n" +
				"	shell\n" +
				"	de\n" +
				"	terminal\n" +
				"	cpu\n" +
				"	memory\n" +
				"	disk\n" +
				"	// battery\n" +
				"	localip\n" +
				"	// time\n" +
				"	// date\n" +
				"}\n\n" +
				"asciisize: default\n" +
				"// Ascii size can be either 'big', 'default' or 'small'. Default is big.\n\n" +
				"customascii: default\n" +
				"// Set a custom ascii logo by providing a path to the txt file containing it.\n\n" +
				"asciicolor: default\n" +
				"// Color of ascii art\n\n" +
				"accentcolor: default\n" +
				"// Color used by the info labels\n\n" +
				"// Available colors:\n" +
				"// black, red, green, yellow, blue,\n" +
				"// magenta, cyan, white,\n" +
				"// bright_black, bright_red,\n" +
				"// bright_green, bright_yellow,\n" +
				"// bright_blue, bright_magenta,\n" +
				"// bright_cyan, bright_white\n\n",
		)

		if err := os.WriteFile(path, []byte(config.String()), 0600); err != nil {
			return err
		}
	}

	return nil
}
