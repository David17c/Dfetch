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

func Shell() string {
	shellPath := os.Getenv("SHELL")
	shell := strings.ToLower(filepath.Base(shellPath))

	if shell == "" {
		return "unknown"
	}

	switch shell {
	case "bash":
		return shellVersion("Bash", "bash", "--version")
	case "zsh":
		return shellVersion("Zsh", "zsh", "--version")
	case "fish":
		return shellVersion("Fish", "fish", "--version")
	case "dash":
		return shellVersion("Dash", "dash", "-V")
	default:
		return shell
	}
}
