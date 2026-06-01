package output

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"
)

func LoadASCII(fs embed.FS, distroID, asciicolor, asciisize, customascii string) ([]string, string) {
	var scanner *bufio.Scanner
	var closer func()

	// Try custom ascii first
	if customascii != "" && customascii != strings.ToLower("default") {
		if f, err := os.Open(customascii); err == nil {
			scanner = bufio.NewScanner(f)
			closer = func() { f.Close() }
		} else {
			fmt.Printf("Error: Custom ascii art path '%s' doesn't seem to exist.\n", customascii)
		}
	}

	// Fallback to default distro logo
	if scanner == nil {
		var file string

		if asciisize == "small" {
			file = fmt.Sprintf("logo/%s_small.txt", strings.ToLower(distroID))
		} else {
			file = fmt.Sprintf("logo/%s_big.txt", strings.ToLower(distroID))
		}

		// Fallback to Linux logo
		if _, err := fs.Open(file); err != nil {
			file = "logo/linux.txt"
		}

		f, err := fs.Open(file)
		if err != nil {
			return nil, asciicolor
		}

		scanner = bufio.NewScanner(f)
		closer = func() { f.Close() }
	}

	defer closer()

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "asciicolor:"):
			if asciicolor == "" || asciicolor == "default" {
				asciicolor = strings.TrimSpace(
					strings.TrimPrefix(line, "asciicolor:"),
				)
			}
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, asciicolor
	}

	return lines, asciicolor
}
