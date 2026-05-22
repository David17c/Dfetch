package assets

import (
	"bufio"
	"embed"
	"fmt"
	"strings"
)

func LoadASCII(fs embed.FS, distroID, color string) ([]string, string) {

	file := fmt.Sprintf("logo/%s.txt", strings.ToLower(distroID))

	// If distro logo also doesn't exist, fallback to linux
	if _, err := fs.Open(file); err != nil {
		file = "logo/linux.txt"
	}

	f, err := fs.Open(file)
	if err != nil {
		return nil, color
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "color:") {
			if color == "" {
				color = strings.TrimSpace(strings.TrimPrefix(line, "color:"))
			}
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, color
	}

	return lines, color
}
