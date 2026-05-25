package render

import (
	"dfetch/internal/model"
	"fmt"
	"regexp"
	"strings"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(s string) int {
	clean := ansiRegex.ReplaceAllString(s, "")
	return len(clean)
}

func BuildInfoLines(sys model.SystemInfo, configLines []string) []string {
	userInfo := fmt.Sprintf("\x1b[1m%s@%s\x1b[0m", sys.Username, sys.Hostname)
	separator := strings.Repeat("-", visibleLen(userInfo))

	infoMap := map[string]string{
		"os":      fmt.Sprintf("OS: %s", sys.DistroName),
		"kernel":  fmt.Sprintf("Kernel: %s", sys.Kernel),
		"cpu":     fmt.Sprintf("CPU: %s", sys.CPU),
		"memory":  fmt.Sprintf("Memory: %s", sys.Memory),
		"localip": fmt.Sprintf("Local IP (%s): %s", sys.IPVersion, sys.LocalIP),
		"uptime":  fmt.Sprintf("Uptime: %s", sys.Uptime),
		"battery": fmt.Sprintf("Battery: %d%% [%s]", sys.Battery, sys.BatteryState),
		"de":      fmt.Sprintf("DE: %s (%s)", sys.DE, sys.SessionType),
		"shell":   fmt.Sprintf("Shell: %s", sys.Shell),
	}

	infoLines := []string{
		userInfo,
		separator,
	}

	for _, line := range configLines {
		line = strings.TrimSpace(strings.ToLower(line))

		if value, exists := infoMap[line]; exists {
			infoLines = append(infoLines, value)
		}
	}

	return infoLines
}

func PrintOutput(asciiLines, infoLines []string, asciicolor string) {
	maxLen := getMaxWidth(asciiLines)

	totalLines := len(asciiLines)
	if len(infoLines) > totalLines {
		totalLines = len(infoLines)
	}

	for i := 0; i < totalLines; i++ {
		left := ""
		right := ""

		if i < len(asciiLines) {
			left = asciiLines[i]
		}

		if i < len(infoLines) {
			right = infoLines[i]
		}

		fmt.Printf("\x1b[1m%s%-*s\x1b[0m %s\n", asciicolor, maxLen, left, right)
	}
}

func getMaxWidth(lines []string) int {
	maxLen := 0

	for _, line := range lines {
		if visibleLen(line) > maxLen {
			maxLen = visibleLen(line)
		}
	}

	return maxLen
}
