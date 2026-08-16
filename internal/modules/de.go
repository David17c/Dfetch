package modules

import (
	"dfetch/internal/config"
	"os"
	"strings"
)

func DesktopEnvironment(format string) string {
	fields := config.Fields(format)

	needsDE := false

	for _, field := range fields {
		switch field {
		case "de", "name":
			needsDE = true
		}
	}

	if !needsDE {
		return config.Format(format, config.Values{})
	}

	name := detectDesktopEnvironment()

	values := config.Values{}

	for _, field := range fields {
		switch field {
		case "de", "name":
			values[field] = name
		}
	}

	return config.Format(format, values)
}

func detectDesktopEnvironment() string {
	id := os.Getenv("XDG_CURRENT_DESKTOP")

	if id == "" {
		id = os.Getenv("DESKTOP_SESSION")
	}

	if id == "" {
		id = os.Getenv("GDMSESSION")
	}

	if id == "" {
		return "unknown"
	}

	for _, de := range strings.FieldsFunc(id, func(r rune) bool {
		return r == ':' || r == ';'
	}) {
		switch strings.ToLower(strings.TrimSpace(de)) {
		case "gnome":
			return "GNOME"

		case "gnome-classic":
			return "GNOME Classic"

		case "kde", "plasma":
			return "KDE Plasma"

		case "xfce":
			return "XFCE"

		case "x-cinnamon", "cinnamon":
			return "Cinnamon"

		case "mate":
			return "MATE"

		case "lxqt":
			return "LXQt"

		case "lxde":
			return "LXDE"

		case "unity":
			return "Unity"

		case "pantheon":
			return "Pantheon"

		case "cosmic":
			return "COSMIC"

		case "budgie":
			return "Budgie"

		case "deepin":
			return "Deepin"
		}
	}

	return id
}
