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

	LabelColor    string
	UserinfoColor string
	CustomAscii   string
	Info_color    string
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

	cfg := &Config{
		LabelColor:    "default",
		UserinfoColor: "default",
		CustomAscii:   "default",
		Info_color:    "default",
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawLine := scanner.Text()
		line := strings.TrimSpace(rawLine)

		// Skip comments
		if idx := strings.Index(line, "//"); idx != -1 {
			line = strings.TrimSpace(line[:idx])
		}

		// Detect start of modules block
		if strings.EqualFold(line, "modules {") {
			inModules = true
			continue
		}

		if inModules {
			if line == "}" {
				inModules = false
				continue
			}

			if line == "" {
				cfg.EnabledModules = append(cfg.EnabledModules, "emptyline")
				continue
			}

			cfg.EnabledModules = append(cfg.EnabledModules, strings.ToLower(line))
			continue
		}

		if idx := strings.Index(line, ":"); idx != -1 {
			key := strings.ToLower(strings.TrimSpace(line[:idx]))
			value := strings.ToLower(strings.TrimSpace(line[idx+1:]))

			switch key {
			case "label_color":
				cfg.LabelColor = value
			case "custom_ascii":
				cfg.CustomAscii = value
			case "userinfo_color":
				cfg.UserinfoColor = value
			case "info_color":
				cfg.Info_color = value
			}

			continue
		}

		if line == "" {
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if inModules {
		return nil, fmt.Errorf("unclosed modules block in config file")
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
				"// In the modules section, you can change which information is displayed and in what order.\n\n" +
				"// Insert empty lines in the modules block to get empty lines in the final output.\n" +
				"modules {\n" +
				"    userinfo\n" +
				"    os\n" +
				"    host\n" +
				"    kernel\n" +
				"    uptime\n" +
				"    shell\n" +
				"    terminal\n" +
				"    desktop\n" +
				"    packages\n" +
				"    cpu\n" +
				"    memory\n" +
				"    swap\n" +
				"    disk\n" +
				"    motherboard\n" +
				"    local_ip\n" +
				"    // battery\n" +
				"    // time\n" +
				"    // date\n" +
				"}\n\n" +
				"custom_ascii: default\n" +
				"// Set a custom ASCII logo by providing the path to a text file containing it.\n\n" +
				"label_color: default\n" +
				"// Color used for the information labels.\n\n" +
				"userinfo_color: default\n" +
				"// Color of the userinfo module.\n\n" +
				"info_color: default\n" +
				"// Color of the system info.\n\n" +
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
