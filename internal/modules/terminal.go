package modules

import (
	"os"
	"strings"
)

type terminalDetector struct {
	env  string
	name string
}

var detectors = []terminalDetector{
	{"ALACRITTY_SOCKET", "Alacritty"},
	{"GNOME_TERMINAL_SCREEN", "GNOME Terminal"},
	{"KITTY_PID", "Kitty"},
	{"GHOSTTY_RESOURCES_DIR", "Ghostty"},
}

func Terminal() string {
	for _, t := range detectors {
		if os.Getenv(t.env) != "" {
			return t.name
		}
	}

	if os.Getenv("KONSOLE_VERSION") != "" {
		return "Konsole"
	}

	switch strings.ToLower(strings.TrimSpace(os.Getenv("TERM_PROGRAM"))) {
	case "":
	case "vscode":
		return "Code"
	case "wezterm":
		return "WezTerm"
	case "ghostty":
		return "Ghostty"
	default:
		return os.Getenv("TERM_PROGRAM")
	}

	term := os.Getenv("TERM")

	if strings.HasPrefix(term, "foot") {
		return "Foot"
	}

	if term == "xterm" {
		return "XTerm"
	}

	if term == "" {
		return "unknown"
	}

	return term
}
