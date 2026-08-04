package output

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"

	"dfetch/internal/cmd"
	"dfetch/internal/config"
	"dfetch/internal/modules"
)

//go:embed logo/*
var LogoFS embed.FS

func LoadASCII(distro modules.DistroInfo, cfg config.Config, flags cmd.Flags) []string {
	var asciiArt []string

	// if the user disabled ascii art
	if !cfg.ASCII.Enabled || flags.NoASCII {
		return nil
	}

	// try ascii art specified using a command-line flag
	if flags.SetASCII != "" && flags.SetASCII != "builtin" {
		if lines, err := ReadASCII(flags.SetASCII); err == nil {
			asciiArt = lines
		}
	}

	// try ascii art specified in the config file
	if asciiArt == nil && cfg.ASCII.Path != "" && cfg.ASCII.Path != "builtin" {
		if lines, err := ReadASCII(cfg.ASCII.Path); err == nil {
			asciiArt = lines
		}
	}

	// use builtin ascii art
	if asciiArt == nil {
		paths := []string{
			fmt.Sprintf("logo/%s.txt", strings.ToLower(distro.ID)),
		}

		if distro.IDLike != "" {
			paths = append(paths,
				fmt.Sprintf("logo/%s.txt", strings.ToLower(distro.IDLike)),
			)
		}

		for _, path := range paths {
			if lines, err := ReadASCII(path); err == nil {
				asciiArt = lines
				break
			}
		}
	}

	// Top padding.
	if cfg.ASCII.PaddingTop > 0 {
		padding := make([]string, cfg.ASCII.PaddingTop)
		asciiArt = append(padding, asciiArt...)
	}

	// Bottom padding.
	if cfg.ASCII.PaddingBottom > 0 {
		padding := make([]string, cfg.ASCII.PaddingBottom)
		asciiArt = append(asciiArt, padding...)
	}

	return asciiArt
}

func ReadASCII(asciiPath string) ([]string, error) {
	var (
		scanner *bufio.Scanner
		closeFn func() error
	)

	if f, err := os.Open(asciiPath); err == nil {
		scanner = bufio.NewScanner(f)
		closeFn = f.Close
	} else {
		f, err := LogoFS.Open(asciiPath)
		if err != nil {
			return nil, err
		}

		scanner = bufio.NewScanner(f)
		closeFn = f.Close
	}

	defer closeFn()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("ASCII file %q is empty", asciiPath)
	}

	return lines, nil
}
