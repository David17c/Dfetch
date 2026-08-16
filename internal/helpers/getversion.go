package helpers

import (
	"os/exec"
	"regexp"
	"strings"
)

var versionRe = regexp.MustCompile(`\b\d+(?:\.\d+)+\b`)

func CommandVersion(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		return ""
	}

	return extractVersion(string(out))
}

func extractVersion(output string) string {
	return versionRe.FindString(output)
}

func FormatVersion(displayName, cmd string, args ...string) string {
	if version := CommandVersion(cmd, args...); version != "" {
		return strings.Join([]string{displayName, version}, " ")
	}

	return displayName
}
