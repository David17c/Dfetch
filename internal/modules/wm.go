package modules

import (
	"fmt"
	"os"
	"strings"
)

func windowManager(format string) string {
	sessionType := os.Getenv("XDG_SESSION_TYPE")

	if sessionType == "" {
		switch {
		case os.Getenv("WAYLAND_DISPLAY") != "":
			sessionType = "wayland"
		case os.Getenv("DISPLAY") != "":
			sessionType = "x11"
		default:
			return "unknown"
		}
	}

	switch sessionType {
	case "wayland":
		wm := WMOnWayland()

		if format == "short" {
			return wm
		}

		return fmt.Sprintf("%s (Wayland)", wm)

	case "x11":
		wm := WMOnX11()

		if format == "short" {
			return wm
		}

		return fmt.Sprintf("%s (X11)", wm)

	default:
		return "unknown"
	}
}

func WMOnWayland() string {
	switch {
	case os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") != "":
		return "Hyprland"
	case os.Getenv("SWAYSOCK") != "":
		return "Sway"
	case os.Getenv("NIRI_SOCKET") != "":
		return "Niri"
	case os.Getenv("WAYFIRE_SOCKET") != "":
		return "Wayfire"
	case os.Getenv("LABWC_PID") != "":
		return "labwc"
	}

	switch processExists(
		"kwin_wayland",
		"gnome-shell",
		"river",
		"cosmic-comp",
	) {
	case "kwin_wayland":
		return "KWin"
	case "gnome-shell":
		return "Mutter"
	case "river":
		return "River"
	case "cosmic-comp":
		return "COSMIC"
	}

	return "unknown"
}

func WMOnX11() string {
	switch {
	case os.Getenv("I3SOCK") != "":
		return "i3"
	case os.Getenv("BSPWM_SOCKET") != "":
		return "bspwm"
	}

	switch processExists(
		"kwin_x11",
		"gnome-shell",
		"xfwm4",
		"i3",
		"bspwm",
		"openbox",
		"awesome",
		"fluxbox",
		"icewm",
		"dwm",
	) {
	case "kwin_x11":
		return "KWin"
	case "gnome-shell":
		return "Mutter"
	case "xfwm4":
		return "Xfwm4"
	case "i3":
		return "i3"
	case "bspwm":
		return "bspwm"
	case "openbox":
		return "Openbox"
	case "awesome":
		return "awesome"
	case "fluxbox":
		return "Fluxbox"
	case "icewm":
		return "IceWM"
	case "dwm":
		return "dwm"
	}

	return "unknown"
}

// Check if a process exists
func processExists(names ...string) string {
	wanted := make(map[string]struct{}, len(names))
	for _, name := range names {
		wanted[name] = struct{}{}
	}

	entries, err := os.ReadDir("/proc")
	if err != nil {
		return ""
	}

	for _, entry := range entries {
		pid := entry.Name()

		if pid == "" || pid[0] < '0' || pid[0] > '9' {
			continue
		}

		data, err := os.ReadFile("/proc/" + pid + "/comm")
		if err != nil {
			continue
		}

		proc := strings.TrimSpace(string(data))
		if _, ok := wanted[proc]; ok {
			return proc
		}
	}

	return ""
}
