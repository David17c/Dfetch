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

//go:embed art/*
var ArtFS embed.FS

func LoadLogo(distro modules.DistroInfo, cfg config.Config, opts cmd.Options) []string {
	var asciiArt []string

	// If the user disabled ASCII art
	if !cfg.Logo.Enabled || opts.NoLogo {
		return nil
	}

	// Try ASCII art specified using a command-line flag
	if opts.SetLogo != "" && opts.SetLogo != "builtin" {
		if lines, err := LoadASCIIByName(opts.SetLogo); err == nil {
			asciiArt = lines
		}
	}

	// Try ASCII art specified in the config file
	if asciiArt == nil && cfg.Logo.Path != "" && cfg.Logo.Path != "builtin" {
		if lines, err := LoadASCIIByName(cfg.Logo.Path); err == nil {
			asciiArt = lines
		}
	}

	// Fall back to distro builtin ASCII art
	if asciiArt == nil {
		for _, name := range []string{distro.ID, distro.IDLike} {
			if name == "" {
				continue
			}

			if lines, err := ReadBuiltinASCII(name); err == nil {
				asciiArt = lines
				break
			}
		}
	}

	// Top padding.
	if cfg.Logo.PaddingTop > 0 {
		padding := make([]string, cfg.Logo.PaddingTop)
		asciiArt = append(padding, asciiArt...)
	}

	// Bottom padding.
	if cfg.Logo.PaddingBottom > 0 {
		padding := make([]string, cfg.Logo.PaddingBottom)
		asciiArt = append(asciiArt, padding...)
	}

	return asciiArt
}

func LoadASCIIByName(name string) ([]string, error) {
	if lines, err := ReadASCII(name); err == nil {
		return lines, nil
	}

	return ReadBuiltinASCII(name)
}

func ReadBuiltinASCII(name string) ([]string, error) {
	return ReadASCII(fmt.Sprintf("art/%s.txt", strings.ToLower(name)))
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

		f, err := ArtFS.Open(asciiPath)
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
