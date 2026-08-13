package modules

import (
	"os"
	"strconv"
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
	{"KONSOLE_VERSION", "Konsole"},
	{"TILIX_ID", "Tilix"},
	{"TERMINATOR_UUID", "Terminator"},
}

func processName(pid int) string {
	data, err := os.ReadFile("/proc/" + strconv.Itoa(pid) + "/comm")
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func parentPID(pid int) int {
	data, err := os.ReadFile("/proc/" + strconv.Itoa(pid) + "/stat")
	if err != nil {
		return 0
	}

	line := string(data)

	end := strings.LastIndex(line, ")")
	if end == -1 || end+2 >= len(line) {
		return 0
	}

	fields := strings.Fields(line[end+2:])
	if len(fields) < 2 {
		return 0
	}

	ppid, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0
	}

	return ppid
}

func processInParentChain(names ...string) string {
	wanted := make(map[string]struct{}, len(names))

	for _, name := range names {
		wanted[name] = struct{}{}
	}

	pid := os.Getpid()

	seen := make(map[int]struct{})

	for pid > 1 {
		if _, ok := seen[pid]; ok {
			break
		}

		seen[pid] = struct{}{}

		proc := processName(pid)

		if _, ok := wanted[proc]; ok {
			return proc
		}

		next := parentPID(pid)
		if next <= 0 || next == pid {
			break
		}

		pid = next
	}

	return ""
}

func Terminal() string {
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

	if proc := processInParentChain(
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
