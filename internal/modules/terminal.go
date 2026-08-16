package modules

import (
	"dfetch/internal/config"
	"dfetch/internal/helpers"
	"os"
	"strings"
)

type terminalElements struct {
	terminalName    string
	terminalVersion string
}

type terminalDetector struct {
	env  string
	name string
}

var detectors = []terminalDetector{
	{"ALACRITTY_SOCKET", "Alacritty"},
	{"GNOME_TERMINAL_SCREEN", "GNOME Terminal"},
	{"KITTY_PID", "Kitty"},
	{"GHOSTTY_RESOURCES_DIR", "Ghostty"},
	{"KONSOLE_VERSION", "Konsole"},
	{"TILIX_ID", "Tilix"},
	{"TERMINATOR_UUID", "Terminator"},
}

func detectTerminal() string {
	for _, t := range detectors {
		if os.Getenv(t.env) != "" {
			return t.name
		}
	}

	switch strings.ToLower(strings.TrimSpace(os.Getenv("TERM_PROGRAM"))) {
	case "vscode":
		return "Code"
	case "wezterm":
		return "WezTerm"
	case "ghostty":
		return "Ghostty"
	case "hyper":
		return "Hyper"
	case "tabby":
		return "Tabby"
	case "rio":
		return "Rio"
	case "iterm.app":
		return "iTerm2"
	case "apple_terminal":
		return "Terminal"
	case "warpterminal":
		return "Warp"
	}

	if proc := helpers.ProcessInParentChain(
		"gnome-terminal-server",
		"konsole",
		"xfce4-terminal",
		"mate-terminal",
		"lxterminal",
		"tilix",
		"terminator",
		"ptyxis",
		"kgx",
		"deepin-terminal",
		"yakuake",
		"cool-retro-term",
		"termite",
		"st",
		"qterminal",
		"urxvt",
		"rxvt",
		"xterm",
		"foot",
		"alacritty",
		"kitty",
		"ghostty",
		"wezterm-gui",
		"rio",
	); proc != "" {
		switch proc {
		case "gnome-terminal-server":
			return "GNOME Terminal"
		case "xfce4-terminal":
			return "Xfce Terminal"
		case "mate-terminal":
			return "MATE Terminal"
		case "lxterminal":
			return "LXTerminal"
		case "ptyxis":
			return "Ptyxis"
		case "kgx":
			return "GNOME Console"
		case "deepin-terminal":
			return "Deepin Terminal"
		case "cool-retro-term":
			return "Cool Retro Term"
		case "qterminal":
			return "QTerminal"
		case "urxvt", "rxvt":
			return "rxvt"
		case "st":
			return "st"
		case "wezterm-gui":
			return "WezTerm"
		default:
			return proc
		}
	}

	term := strings.TrimSpace(os.Getenv("TERM"))

	switch {
	case strings.HasPrefix(term, "foot"):
		return "Foot"
	case strings.HasPrefix(term, "xterm"):
		return "XTerm"
	case strings.HasPrefix(term, "rxvt"):
		return "rxvt"
	case strings.HasPrefix(term, "screen"):
		return "Screen"
	case strings.HasPrefix(term, "tmux"):
		return "tmux"
	case term == "":
		return "unknown"
	default:
		return term
	}
}

func terminalVersion(name string) string {
	var command string
	var args string

	switch name {
	case "Alacritty":
		command = "alacritty"
		args = "--version"

	case "Kitty":
		command = "kitty"
		args = "--version"

	case "WezTerm":
		command = "wezterm"
		args = "--version"

	case "Ghostty":
		command = "ghostty"
		args = "version"

	case "Konsole":
		command = "konsole"
		args = "--version"

	default:
		return ""
	}

	version := helpers.CommandVersion(command, args)

	fields := strings.Fields(version)
	if len(fields) >= 2 {
		return fields[len(fields)-1]
	}

	return version
}

func Terminal(format string) string {
	formatList := config.ExtractFormat(format)

	elm := terminalElements{
		terminalName: detectTerminal(),
	}

	for _, value := range formatList {
		if value == "version" {
			elm.terminalVersion = terminalVersion(elm.terminalName)
			break
		}
	}

	return formatTerminalOutput(formatList, elm)
}

func formatTerminalOutput(formatList []string, elm terminalElements) string {
	var output []string

	for _, value := range formatList {
		switch value {
		case "name":
			output = append(output, elm.terminalName)

		case "version":
			if elm.terminalVersion != "" {
				output = append(output, elm.terminalVersion)
			}
		}
	}

	return strings.Join(output, " ")
}
