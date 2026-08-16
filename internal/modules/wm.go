package modules

import (
	"dfetch/internal/config"
	"dfetch/internal/helpers"
	"os"
)

func WindowManager(format string) string {
	fields := config.Fields(format)
	values := config.Values{}

	var needsName, needsVersion, needsSessionType bool

	for _, field := range fields {
		switch field {
		case "name":
			needsName = true

		case "version":
			needsName = true
			needsVersion = true

		case "sessiontype":
			needsSessionType = true
		}
	}

	var sessionType string

	// Session type is only needed when the format asks for it
	// or when it is needed to detect the window manager.
	if needsSessionType || needsName {
		sessionType = os.Getenv("XDG_SESSION_TYPE")

		if sessionType == "" {
			switch {
			case os.Getenv("WAYLAND_DISPLAY") != "":
				sessionType = "wayland"

			case os.Getenv("DISPLAY") != "":
				sessionType = "x11"

			default:
				sessionType = "unknown"
			}
		}
	}

	if needsName {
		var name string

		switch sessionType {
		case "wayland":
			name = WMOnWayland()

		case "x11":
			name = WMOnX11()

		default:
			name = "unknown"
		}

		values["name"] = name

		if needsVersion && name != "unknown" {
			values["version"] = WMVersion(name)
		}
	}

	if needsSessionType {
		values["sessiontype"] = sessionType
	}

	return config.Format(format, values)
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
