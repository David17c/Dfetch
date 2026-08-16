package modules

import (
	"dfetch/internal/config"
	"dfetch/internal/helpers"
	"os"
	"path/filepath"
	"strings"
)

type shellElements struct {
	shellName    string
	shellVersion string
}

func Shell(format string) string {
	shellPath := os.Getenv("SHELL")
	shell := strings.ToLower(filepath.Base(shellPath))

	if shell == "" {
		return "unknown"
	}

	var elm shellElements

	switch shell {
	case "bash":
		elm.shellName = "Bash"
		elm.shellVersion = helpers.CommandVersion(shellPath, "--version")

	case "zsh":
		elm.shellName = "Zsh"
		elm.shellVersion = helpers.CommandVersion(shellPath, "--version")

	case "fish":
		elm.shellName = "Fish"
		elm.shellVersion = helpers.CommandVersion(shellPath, "--version")

	case "dash":
		elm.shellName = "Dash"
		elm.shellVersion = helpers.CommandVersion(shellPath, "-V")

	case "ksh":
		elm.shellName = "Ksh"
		elm.shellVersion = helpers.CommandVersion(shellPath, "--version")

	default:
		elm.shellName = shell
	}

	return formatShellOutput(config.ExtractFormat(format), elm)
}

func formatShellOutput(formatList []string, elm shellElements) string {
	var output []string

	for _, value := range formatList {
		switch value {
		case "name":
			output = append(output, elm.shellName)
		case "version":
			output = append(output, elm.shellVersion)
		}
	}

	return strings.Join(output, " ")
}
