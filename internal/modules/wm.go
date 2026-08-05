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

	processes := getRunningProcesses()

	switch {
	case hasProcess(processes, "kwin_wayland"):
		return "KWin"

	case hasProcess(processes, "gnome-shell"):
		return "Mutter"

	case hasProcess(processes, "river"):
		return "River"

	case hasProcess(processes, "cosmic-comp"):
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

	processes := getRunningProcesses()

	switch {
	case hasProcess(processes, "kwin_x11"):
		return "KWin"

	case hasProcess(processes, "gnome-shell"):
		return "Mutter"

	case hasProcess(processes, "xfwm4"):
		return "Xfwm4"

	case hasProcess(processes, "i3"):
		return "i3"

	case hasProcess(processes, "bspwm"):
		return "bspwm"

	case hasProcess(processes, "openbox"):
		return "Openbox"

	case hasProcess(processes, "awesome"):
		return "awesome"

	case hasProcess(processes, "fluxbox"):
		return "Fluxbox"

	case hasProcess(processes, "icewm"):
		return "IceWM"

	case hasProcess(processes, "dwm"):
		return "dwm"
	}

	return "unknown"
}

func getRunningProcesses() map[string]bool {
	processes := make(map[string]bool)

	entries, err := os.ReadDir("/proc")
	if err != nil {
		return processes
	}

	for _, entry := range entries {
		name := entry.Name()

		if name == "" || name[0] < '0' || name[0] > '9' {
			continue
		}

		data, err := os.ReadFile("/proc/" + name + "/comm")
		if err != nil {
			continue
		}

		process := strings.TrimSpace(string(data))
		if process != "" {
			processes[process] = true
		}
	}

	return processes
}

func hasProcess(processes map[string]bool, name string) bool {
	return processes[name]
}
