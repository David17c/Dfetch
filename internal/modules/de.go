package modules

import (
	"os"
	"strings"
)

func DesktopEnvironment(format string) string {
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
			if format == "short" {
				return "GNOME"
			}
			return "GNOME"

		case "gnome-classic":
			if format == "short" {
				return "GNOME Classic"
			}
			return "GNOME Classic"

		case "kde", "plasma":
			if format == "short" {
				return "KDE Plasma"
			}
			return "KDE Plasma"

		case "xfce":
			if format == "short" {
				return "XFCE"
			}
			return "XFCE"

		case "x-cinnamon", "cinnamon":
			if format == "short" {
				return "Cinnamon"
			}
			return "Cinnamon"

		case "mate":
			if format == "short" {
				return "MATE"
			}
			return "MATE"

		case "lxqt":
			if format == "short" {
				return "LXQt"
			}
			return "LXQt"

		case "lxde":
			if format == "short" {
				return "LXDE"
			}
			return "LXDE"

		case "unity":
			if format == "short" {
				return "Unity"
			}
			return "Unity"

		case "pantheon":
			if format == "short" {
				return "Pantheon"
			}
			return "Pantheon"

		case "cosmic":
			if format == "short" {
				return "COSMIC"
			}
			return "COSMIC"

		case "budgie":
			if format == "short" {
				return "Budgie"
			}
			return "Budgie"

		case "deepin":
			if format == "short" {
				return "Deepin"
			}
			return "Deepin"
		}
	}

	return id
}
