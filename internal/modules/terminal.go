package modules

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type terminalDetector struct {
	env    string
	name   string
	binary string
	parse  func(string) string
}

var ghosttyVersion = func(out string) string {
	line := strings.TrimSpace(strings.SplitN(out, "\n", 2)[0])
	return strings.TrimPrefix(line, "Ghostty ")
}

var detectors = []terminalDetector{
	{"ALACRITTY_SOCKET", "Alacritty", "alacritty", extractVersion},
	{"GNOME_TERMINAL_SCREEN", "GNOME Terminal", "gnome-terminal", extractVersion},
	{"KITTY_PID", "Kitty", "kitty", extractVersion},
	{"GHOSTTY_RESOURCES_DIR", "Ghostty", "ghostty", ghosttyVersion},
}

func terminalVersion(name, binary string, parse func(string) string, format string) string {
	if format == "short" {
		return name
	}

	out, err := exec.Command(binary, "--version").Output()
	if err != nil {
		return name
	}

	if parse == nil {
		parse = extractVersion
	}

	if v := parse(string(out)); v != "" {
		return fmt.Sprintf("%s %s", name, v)
	}

	return name
}

func Terminal(format string) string {
	for _, t := range detectors {
		if os.Getenv(t.env) != "" {
			return terminalVersion(t.name, t.binary, t.parse, format)
		}
	}

	if v := os.Getenv("KONSOLE_VERSION"); v != "" {
		if format == "short" {
			return "Konsole"
		}
		return "Konsole " + v
	}

	switch strings.ToLower(strings.TrimSpace(os.Getenv("TERM_PROGRAM"))) {
	case "":
	case "vscode":
		return "Code"
	case "wezterm":
		return terminalVersion("WezTerm", "wezterm", nil, format)
	case "ghostty":
		return terminalVersion("Ghostty", "ghostty", ghosttyVersion, format)
	default:
		return os.Getenv("TERM_PROGRAM")
	}

	term := os.Getenv("TERM")

	if strings.HasPrefix(term, "foot") {
		return terminalVersion("Foot", "foot", nil, format)
	}

	if term == "xterm" {
		if format == "short" {
			return "XTerm"
		}

		out, err := exec.Command("xterm", "-v").Output()
		if err != nil {
			return "XTerm"
		}

		var version int
		if _, err := fmt.Sscanf(strings.TrimSpace(string(out)), "XTerm(%d)", &version); err == nil {
			return fmt.Sprintf("XTerm %d", version)
		}

		return "XTerm"
	}

	if term == "" {
		return "unknown"
	}

	return term
}
