package modules

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func DesktopEnvironment() string {
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

	for _, de := range strings.Split(id, ":") {
		switch strings.ToLower(strings.TrimSpace(de)) {
		case "gnome":
			output, err := exec.Command("gnome-shell", "--version").Output()
			if err != nil {
				return "GNOME"
			}

			fields := strings.Fields(string(output))
			if len(fields) >= 3 {
				return fmt.Sprintf("GNOME %s", fields[len(fields)-1])
			}

			return "GNOME"

		case "kde", "plasma":
			output, err := exec.Command("plasmashell", "--version").Output()
			if err != nil {
				output, err = exec.Command("kf6-config", "--version").Output()
				if err != nil {
					return "KDE Plasma"
				}
			}

			fields := strings.Fields(string(output))
			if len(fields) >= 2 {
				return fmt.Sprintf("KDE Plasma %s", fields[len(fields)-1])
			}

			return "KDE Plasma"

		case "xfce":
			output, err := exec.Command("xfce4-session", "--version").Output()
			if err != nil {
				return "XFCE"
			}
			fields := strings.Fields(string(output))
			if len(fields) >= 2 {
				return fmt.Sprintf("XFCE %s", fields[1])
			}

		case "x-cinnamon", "cinnamon":
			output, err := exec.Command("cinnamon", "--version").Output()
			if err == nil {
				return strings.TrimSpace(string(output))
			}
			return "Cinnamon"

		case "mate":
			output, _ := exec.Command("mate-session", "--version").CombinedOutput()

			re := regexp.MustCompile(`\d+\.\d+(?:\.\d+)?`)
			if version := re.FindString(string(output)); version != "" {
				return fmt.Sprintf("MATE %s", version)
			}

			return "MATE"

		case "lxqt":
			output, err := exec.Command("lxqt-session", "--version").Output()
			if err != nil {
				return "LXQt"
			}
			fields := strings.Fields(string(output))
			if len(fields) >= 2 {
				return fmt.Sprintf("LXQT %s", fields[1])
			}

		case "unity":
			output, err := exec.Command("unity", "--version").Output()
			if err == nil {
				return strings.TrimSpace(string(output))
			}

			return "Unity"
		}
	}
	return id
}

// Support for more desktop enviroments like Cosmic and deepin will be added over time
