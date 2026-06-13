package output

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"
)

func LoadASCII(fs embed.FS, distroID, asciisize, customascii string) ([]string, string) {
	var scanner *bufio.Scanner
	var closer func()

	// custom file
	if customascii != "" && customascii != "default" {
		if f, err := os.Open(customascii); err == nil {
			scanner = bufio.NewScanner(f)
			closer = func() { f.Close() }
		} else {
			fmt.Printf("Error: Custom ascii path '%s' doesn't seem to exist.\n", customascii)
		}
	}

	// fallback
	if scanner == nil {
		var file string

		if asciisize == "small" {
			file = fmt.Sprintf("logo/%s_small.txt", strings.ToLower(distroID))
		} else {
			file = fmt.Sprintf("logo/%s_big.txt", strings.ToLower(distroID))
		}

		if _, err := fs.Open(file); err != nil {
			if asciisize == "small" {
				file = "logo/linux_small.txt"
			} else {
				file = "logo/linux_big.txt"
			}
		}

		f, err := fs.Open(file)
		if err != nil {
			return nil, ""
		}

		scanner = bufio.NewScanner(f)
		closer = func() { f.Close() }
	}

	defer closer()

	var lines []string
	var accentColor string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(strings.ToLower(line), "accent_color:") {
			accentColor = strings.TrimSpace(strings.TrimPrefix(line, "accent_color:"))
			continue
		}
		lines = append(lines, line)
	}

	return lines, accentColor
}
