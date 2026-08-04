package output

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"

	"dfetch/internal/config"
	"dfetch/internal/modules"
)

//go:embed logo/*
var LogoFS embed.FS

func LoadASCII(distro modules.DistroInfo, cfg config.Config, noASCII bool) []string {
	if !cfg.ASCII.Enabled || noASCII {
		return nil
	}

	var lines []string

	if cfg.ASCII.Path != "" && cfg.ASCII.Path != "builtin" {
		f, err := os.Open(cfg.ASCII.Path)
		if err == nil {
			defer f.Close()

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				lines = UseBuiltinASCII(distro)
			}
		} else {
			lines = UseBuiltinASCII(distro)
		}
	} else {
		lines = UseBuiltinASCII(distro)
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

func UseBuiltinASCII(distro modules.DistroInfo) []string {
	var lines []string

	path := fmt.Sprintf("logo/%s.txt", strings.ToLower(distro.ID))

	f, err := LogoFS.Open(path)
	if err != nil {
		path = fmt.Sprintf("logo/%s.txt", strings.ToLower(distro.IDLike))
		f, err = LogoFS.Open(path)
		if err != nil {
			return nil
		}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return lines
}
