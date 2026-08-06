package modules

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var versionRe = regexp.MustCompile(`\b\d+\.\d+(?:\.\d+)?\b`)

func shellVersion(displayName, cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		return displayName
	}

	if v := extractVersion(string(out)); v != "" {
		return fmt.Sprintf("%s %s", displayName, v)
	}

	return displayName
}

func extractVersion(output string) string {
	return versionRe.FindString(output)
}

func Shell(format string) string {
	shellPath := os.Getenv("SHELL")
	shell := strings.ToLower(filepath.Base(shellPath))

	if shell == "" {
		return "unknown"
	}

	switch shell {
	case "bash":
		if format == "short" {
			return "Bash"
		}
		return shellVersion("Bash", shellPath, "--version")
	case "zsh":
		if format == "short" {
			return "Zsh"
		}
		return shellVersion("Zsh", shellPath, "--version")
	case "fish":
		if format == "short" {
			return "Fish"
		}
		return shellVersion("Fish", shellPath, "--version")
	case "dash":
		if format == "short" {
			return "Dash"
		}
		return shellVersion("Dash", shellPath, "-V")
	default:
		return shell
	}
}
