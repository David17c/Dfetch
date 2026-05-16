// customization/colors.go
package customization

import (
	"strings"
)

// List of all supported colors
func GetColorCode(color string) string {
	switch strings.ToLower(color) {
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

	default:
		return "\x1b[0m"
	}
}
