package modules

import (
	"dfetch/internal/format"
	"os"
	"strings"
)

func DesktopEnvironment(formatstring string) string {
	fields := format.Fields(formatstring)

	needsDE := false

	for _, field := range fields {
		switch field {
		case "de", "name":
			needsDE = true
		}
	}

	if !needsDE {
		return format.Format(formatstring, format.Values{})
	}

	name := detectDesktopEnvironment()

	values := format.Values{}

	for _, field := range fields {
		switch field {
		case "de", "name":
			values[field] = name
		}
	}

	return format.Format(formatstring, values)
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
	return "unknown"
}
