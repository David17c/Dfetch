package config

import (
	"dfetch/internal/getsysinfo"
	"os"
	"path/filepath"
	"strings"
)

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
				"labelcolor: default\n" +
				"infocolor: default\n" +
				"headercolor: default\n" +
				"asciicolor: default\n\n" +
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
				"//shell\n" +
				"//de\n",
		)

		// Enable battery in config if one is present
		_, present := getsysinfo.Battery()
		if present != "unknown" {
			config.WriteString("battery\n")
		} else {
			config.WriteString("//battery\n")
		}

		err = os.WriteFile(configFile, []byte(config.String()), 0600)
		if err != nil {
			return err
		}
	}

	return nil
}
