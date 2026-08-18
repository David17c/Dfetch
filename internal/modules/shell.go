package modules

import (
	"dfetch/internal/format"
	"dfetch/internal/helpers"
	"os"
	"path/filepath"
	"strings"
)

type shellElements struct {
	name    string
	version string
}

func Shell(formatstring string) string {
	shellPath := os.Getenv("SHELL")
	shell := strings.ToLower(filepath.Base(shellPath))

	if shell == "" {
		return format.Format(formatstring, format.Values{
			"name":    "unknown",
			"version": "unknown",
		})
	}

	var elm shellElements

	switch shell {
	case "bash":
		elm.name = "Bash"
		elm.version = helpers.CommandVersion(shellPath, "--version")

	case "zsh":
		elm.name = "Zsh"
		elm.version = helpers.CommandVersion(shellPath, "--version")

	case "fish":
		elm.name = "Fish"
		elm.version = helpers.CommandVersion(shellPath, "--version")

	case "dash":
		elm.name = "Dash"
		elm.version = helpers.CommandVersion(shellPath, "-V")

	case "ksh":
		elm.name = "Ksh"
		elm.version = helpers.CommandVersion(shellPath, "--version")

	default:
		elm.name = shell
	}

	return format.Format(formatstring, map[string]string{
		"name":    elm.name,
		"version": elm.version,
	})
}
