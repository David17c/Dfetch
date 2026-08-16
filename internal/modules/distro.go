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

	default:
		return d.ID
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
	info, err := parseOSRelease("/etc/os-release")
	if err == nil && info.ID != "" {
		return info, nil
	}

	info, err = parseOSRelease("/usr/lib/os-release")
	if err == nil && info.ID != "" {
		return info, nil
	}

	data, err := os.ReadFile("/etc/issue")
	if err != nil {
		return DistroInfo{}, err
	}

	name := strings.TrimSpace(string(data))

	if idx := strings.IndexRune(name, '\\'); idx != -1 {
		name = strings.TrimSpace(name[:idx])
	}

	if name == "" {
		return DistroInfo{}, fmt.Errorf("unable to determine distribution")
	}

	return DistroInfo{
		PrettyName: name,
	}, nil
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
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		value = strings.Trim(value, `"`)

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
