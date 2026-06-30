package output

import (
	"bufio"
	"dfetch/internal/config"
	"embed"
	"fmt"
	"os"
	"strings"
)

func LoadASCII(fs embed.FS, id string, cfg *config.Config) []string {
	var scanner *bufio.Scanner
	var closer func()

	// custom file
	if cfg.CustomAscii != "" && cfg.CustomAscii != "default" {
		if f, err := os.Open(cfg.CustomAscii); err == nil {
			scanner = bufio.NewScanner(f)
			closer = func() { f.Close() }
		} else {
			fmt.Printf("Error: Custom ascii path '%s' doesn't seem to exist.\n", cfg.CustomAscii)
		}
	}

	// Build in ascii
	if scanner == nil {
		var file string

		file = fmt.Sprintf("logo/%s.txt", strings.ToLower(id))

		// Fallback to Linux logo
		if _, err := fs.Open(file); err != nil {
			file = "logo/linux.txt"
		}

		// Fallback to no ascii
		f, err := fs.Open(file)
		if err != nil {
			return nil
		}

		scanner = bufio.NewScanner(f)
		closer = func() { f.Close() }
	}

	defer closer()

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lower := strings.ToLower(line)

		// grab the colors not set in the config file from the ascii file
		if strings.HasPrefix(lower, "label_color:") &&
			(cfg.LabelColor == "" || cfg.LabelColor == "default") {

			cfg.LabelColor = strings.TrimSpace(line[len("label_color:"):])
			continue
		}

		if strings.HasPrefix(lower, "userinfo_color:") &&
			(cfg.UserinfoColor == "" || cfg.UserinfoColor == "default") {

			cfg.UserinfoColor = strings.TrimSpace(line[len("userinfo_color:"):])
			continue
		}

		if strings.HasPrefix(lower, "info_color:") &&
			(cfg.Info_color == "" || cfg.Info_color == "default") {

			cfg.Info_color = strings.TrimSpace(line[len("info_color:"):])
			continue
		}

		lines = append(lines, line)
	}

	return lines
}
