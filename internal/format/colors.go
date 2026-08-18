package format

import "strings"

func GetColorCode(accentColor string, NoColor bool) string {
	if NoColor {
		return ""
	}

	switch strings.ToLower(accentColor) {
	case "reset":
		return "\x1b[0m"

	case "black", "1":
		return "\x1b[30m"
	case "red", "2":
		return "\x1b[31m"
	case "green", "3":
		return "\x1b[32m"
	case "yellow", "4":
		return "\x1b[33m"
	case "blue", "5":
		return "\x1b[34m"
	case "magenta", "6":
		return "\x1b[35m"
	case "cyan", "7":
		return "\x1b[36m"
	case "white", "8":
		return "\x1b[37m"

	case "bright_black", "9":
		return "\x1b[90m"
	case "bright_red", "10":
		return "\x1b[91m"
	case "bright_green", "11":
		return "\x1b[92m"
	case "bright_yellow", "12":
		return "\x1b[93m"
	case "bright_blue", "13":
		return "\x1b[94m"
	case "bright_magenta", "14":
		return "\x1b[95m"
	case "bright_cyan", "15":
		return "\x1b[96m"
	case "bright_white", "16":
		return "\x1b[97m"

	case "bold_black", "17":
		return "\x1b[1;30m"
	case "bold_red", "18":
		return "\x1b[1;31m"
	case "bold_green", "19":
		return "\x1b[1;32m"
	case "bold_yellow", "20":
		return "\x1b[1;33m"
	case "bold_blue", "21":
		return "\x1b[1;34m"
	case "bold_magenta", "22":
		return "\x1b[1;35m"
	case "bold_cyan", "23":
		return "\x1b[1;36m"
	case "bold_white", "24":
		return "\x1b[1;37m"

	case "bold_bright_black", "25":
		return "\x1b[1;90m"
	case "bold_bright_red", "26":
		return "\x1b[1;91m"
	case "bold_bright_green", "27":
		return "\x1b[1;92m"
	case "bold_bright_yellow", "28":
		return "\x1b[1;93m"
	case "bold_bright_blue", "29":
		return "\x1b[1;94m"
	case "bold_bright_magenta", "30":
		return "\x1b[1;95m"
	case "bold_bright_cyan", "31":
		return "\x1b[1;96m"
	case "bold_bright_white", "32":
		return "\x1b[1;97m"
	}

	return ""
}
