package modules

import (
	"dfetch/internal/config"
	"dfetch/internal/helpers"
	"os"
	"strings"
)

type wmElements struct {
	wmName      string
	wmVersion   string
	sessionType string
}

func WindowManager(format string) string {
	formatList := config.ExtractFormat(format)

	var needsName, needsVersion bool

	for _, value := range formatList {
		switch value {
		case "name":
			needsName = true
		case "version":
			needsName = true
			needsVersion = true
		}
	}

	var elm wmElements

	elm.sessionType = os.Getenv("XDG_SESSION_TYPE")

	if elm.sessionType == "" {
		switch {
		case os.Getenv("WAYLAND_DISPLAY") != "":
			elm.sessionType = "wayland"
		case os.Getenv("DISPLAY") != "":
			elm.sessionType = "x11"
		default:
			return "unknown"
		}
	}

	if needsName {
		switch elm.sessionType {
		case "wayland":
			elm.wmName = WMOnWayland()
		case "x11":
			elm.wmName = WMOnX11()
		default:
			elm.wmName = "unknown"
		}
	}

	if needsVersion && elm.wmName != "unknown" {
		elm.wmVersion = WMVersion(elm.wmName)
	}

	return formatWMOutput(formatList, elm)
}

func formatWMOutput(formatList []string, elm wmElements) string {
	var output []string

	for _, value := range formatList {
		switch value {
		case "name":
			output = append(output, elm.wmName)
		case "version":
			output = append(output, elm.wmVersion)
		case "sessiontype":
			output = append(output, elm.sessionType)
		}
	}

	return strings.Join(output, " ")
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

	switch helpers.ProcessExists(
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

	switch helpers.ProcessExists(
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

func WMVersion(name string) string {
	switch name {
	case "Hyprland":
		return helpers.CommandVersion("hyprctl", "version")
	case "Sway":
		return helpers.CommandVersion("sway", "--version")
	case "Niri":
		return helpers.CommandVersion("niri", "--version")
	case "Wayfire":
		return helpers.CommandVersion("wayfire", "--version")
	case "labwc":
		return helpers.CommandVersion("labwc", "--version")
	case "KWin":
		return helpers.CommandVersion("kwin_wayland", "--version")
	case "Mutter":
		return helpers.CommandVersion("gnome-shell", "--version")
	case "River":
		return helpers.CommandVersion("river", "--version")
	case "COSMIC":
		return helpers.CommandVersion("cosmic-comp", "--version")
	case "i3":
		return helpers.CommandVersion("i3", "--version")
	case "bspwm":
		return helpers.CommandVersion("bspwm", "--version")
	case "Xfwm4":
		return helpers.CommandVersion("xfwm4", "--version")
	case "Openbox":
		return helpers.CommandVersion("openbox", "--version")
	case "awesome":
		return helpers.CommandVersion("awesome", "--version")
	case "fluxbox":
		return helpers.CommandVersion("fluxbox", "-v")
	case "IceWM":
		return helpers.CommandVersion("icewm", "--version")
	default:
		return ""
	}
}

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
