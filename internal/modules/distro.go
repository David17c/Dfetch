package modules

import (
	"bufio"
	"dfetch/internal/config"
	"fmt"
	"os"
	"strings"
)

type DistroInfo struct {
	PrettyName string
	ID         string
	Name       string
	Version    string
	IDLike     string
}

func (d DistroInfo) DisplayName() string {
	switch {
	case d.PrettyName != "":
		return d.PrettyName

	case d.Name != "" && d.Version != "":
		return fmt.Sprintf("%s %s", d.Name, d.Version)

	case d.Name != "":
		return d.Name

	case d.ID != "":
		return d.ID

	default:
		return "unknown"
	}
}

func OS(format string) string {
	info, err := Distro()
	if err != nil {
		return config.Format(format, config.Values{
			"name": "unknown",
		})
	}

	return config.Format(format, config.Values{
		"name": info.DisplayName(),
	})
}

func Distro() (DistroInfo, error) {
	if _, err := os.Stat("/etc/os-release"); err == nil {
		return parseOSRelease("/etc/os-release")
	}

	return parseOSRelease("/usr/lib/os-release")
}

func parseOSRelease(path string) (DistroInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return DistroInfo{}, err
	}
	defer file.Close()

	var info DistroInfo

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = parseOSReleaseValue(value)

		switch key {
		case "PRETTY_NAME":
			info.PrettyName = value
		case "NAME":
			info.Name = value
		case "VERSION":
			info.Version = value
		case "ID":
			info.ID = value
		case "ID_LIKE":
			info.IDLike = value
		}
	}

	if err := scanner.Err(); err != nil {
		return DistroInfo{}, err
	}

	return info, nil
}

func parseOSReleaseValue(value string) string {
	value = strings.TrimSpace(value)

	if len(value) < 2 {
		return value
	}

	quote := value[0]
	if (quote != '"' && quote != '\'') || value[len(value)-1] != quote {
		return value
	}

	value = value[1 : len(value)-1]

	var b strings.Builder
	b.Grow(len(value))

	escaped := false
	for _, r := range value {
		if escaped {
			b.WriteRune(r)
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		b.WriteRune(r)
	}

	if escaped {
		b.WriteByte('\\')
	}

	return b.String()
}
