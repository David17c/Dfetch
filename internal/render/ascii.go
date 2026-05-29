package render

import (
	"bufio"
	"embed"
	"fmt"
	"strings"
)

func LoadASCII(fs embed.FS, distroID, asciicolor string) ([]string, string) {

	file := fmt.Sprintf("logo/%s.txt", strings.ToLower(distroID))

	// If distro logo doesn't exist use Linux logo
	if _, err := fs.Open(file); err != nil {
		file = "logo/linux.txt"
	}

	f, err := fs.Open(file)
	if err != nil {
		return nil, asciicolor
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "asciicolor:") {
			if asciicolor == "" || asciicolor == "default" {
				asciicolor = strings.TrimSpace(strings.TrimPrefix(line, "asciicolor:"))
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
