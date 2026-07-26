package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Distro() (string, string) {
	prettyName, id, err := parseOSRelease("/etc/os-release")
	if err == nil && prettyName != "" {
		return prettyName, id
	}

	// First fallback
	prettyName, id, err = parseOSRelease("/usr/lib/os-release")
	if err == nil && prettyName != "" {
		return prettyName, id
	}

	// Second fallback
	data, err := os.ReadFile("/etc/issue")
	if err == nil {
		prettyName := strings.TrimSpace(string(data))

		if idx := strings.IndexRune(prettyName, '\\'); idx != -1 {
			prettyName = strings.TrimSpace(prettyName[:idx])
		}

		if prettyName != "" {
			return prettyName, ""
		}
	}

	return "unknown", "unknown"
}

func parseOSRelease(path string) (string, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		distroName string
		prettyName string
		name       string
		version    string
		id         string
	)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "PRETTY_NAME="):
			prettyName = strings.Trim(
				strings.TrimPrefix(line, "PRETTY_NAME="),
				"\"",
			)

		case strings.HasPrefix(line, "NAME="):
			name = strings.Trim(
				strings.TrimPrefix(line, "NAME="),
				"\"",
			)

		case strings.HasPrefix(line, "VERSION="):
			version = strings.Trim(
				strings.TrimPrefix(line, "VERSION="),
				"\"",
			)

		case strings.HasPrefix(line, "ID="):
			id = strings.Trim(
				strings.TrimPrefix(line, "ID="),
				"\"",
			)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	if name != "" && version != "" {
		distroName = fmt.Sprintf("%s %s", name, version)
	} else if prettyName != "" {
		distroName = prettyName
	} else if name != "" {
		distroName = name
	}

	if id == "" {
		if name != "" {
			id = strings.ToLower(name)
		} else {
			id = "unknown"
		}
	}

	return distroName, id, nil
}
