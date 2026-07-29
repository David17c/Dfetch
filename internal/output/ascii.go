package output

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"

	"dfetch/internal/config"
)

func LoadASCII(fs embed.FS, distroID string, cfg config.Config) []string {
	if !cfg.ASCII.Enabled {
		return nil
	}

	var scanner *bufio.Scanner
	var closeFn func()

	if cfg.ASCII.Path == "builtin" {
		path := fmt.Sprintf("logo/%s.txt", strings.ToLower(distroID))

		f, err := fs.Open(path)
		if err != nil {
			return nil
		}

		scanner = bufio.NewScanner(f)
		closeFn = func() {
			f.Close()
		}
	} else {
		f, err := os.Open(cfg.ASCII.Path)
		if err != nil {
			return nil
		}

		scanner = bufio.NewScanner(f)
		closeFn = func() {
			f.Close()
		}
	}

	defer closeFn()

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if cfg.ASCII.PaddingTop > 0 {
		padding := make([]string, cfg.ASCII.PaddingTop)
		lines = append(padding, lines...)
	}

	if cfg.ASCII.PaddingBottom > 0 {
		padding := make([]string, cfg.ASCII.PaddingBottom)
		lines = append(lines, padding...)
	}

	return lines
}
