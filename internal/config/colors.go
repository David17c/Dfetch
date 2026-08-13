package config

import (
	"strings"
)

func GetColorCode(accentColor string, NoColor bool) string {
	if NoColor {
		return ""
	}

	switch strings.ToLower(accentColor) {
	// Normal colors
	case "black":
		return "\x1b[30m"
	case "red":
		return "\x1b[31m"
	case "green":
		return "\x1b[32m"
	case "yellow":
		return "\x1b[33m"
	case "blue":
		return "\x1b[34m"
	case "magenta":
		return "\x1b[35m"
	case "cyan":
		return "\x1b[36m"
	case "white":
		return "\x1b[37m"

	// Bright colors
	case "bright_black", "gray", "grey":
		return "\x1b[90m"
	case "bright_red":
		return "\x1b[91m"
	case "bright_green":
		return "\x1b[92m"
	case "bright_yellow":
		return "\x1b[93m"
	case "bright_blue":
		return "\x1b[94m"
	case "bright_magenta":
		return "\x1b[95m"
	case "bright_cyan":
		return "\x1b[96m"
	case "bright_white":
		return "\x1b[97m"

	case "bold_black":
		return "\x1b[1;30m"
	case "bold_red":
		return "\x1b[1;31m"
	case "bold_green":
		return "\x1b[1;32m"
	case "bold_yellow":
		return "\x1b[1;33m"
	case "bold_blue":
		return "\x1b[1;34m"
	case "bold_magenta":
		return "\x1b[1;35m"
	case "bold_cyan":
		return "\x1b[1;36m"
	case "bold_white":
		return "\x1b[1;37m"

	case "bold_bright_black":
		return "\x1b[1;90m"
	case "bold_bright_red":
		return "\x1b[1;91m"
	case "bold_bright_green":
		return "\x1b[1;92m"
	case "bold_bright_yellow":
		return "\x1b[1;93m"
	case "bold_bright_blue":
		return "\x1b[1;94m"
	case "bold_bright_magenta":
		return "\x1b[1;95m"
	case "bold_bright_cyan":
		return "\x1b[1;96m"
	case "bold_bright_white":
		return "\x1b[1;97m"

	default:
		return "\x1b[0m"
	}
}
