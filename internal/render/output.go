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

func BuildInfoLines(sys model.SystemInfo, configLines []string, headercolor string, infocolor string, labelcolor string) []string {

	userInfo := fmt.Sprintf("\x1b[1m%s[%s@%s]\x1b[0m", headercolor, sys.Username, sys.Hostname)

	infoMap := map[string]string{
		"os": fmt.Sprintf(
			"%sOS:\x1b[0m %s%s\x1b[0m",
			labelcolor,
			infocolor,
			sys.DistroName,
		),

		"kernel": fmt.Sprintf(
			"%sKernel:\x1b[0m %s%s\x1b[0m",
			labelcolor,
			infocolor,
			sys.Kernel,
		),

		"cpu": fmt.Sprintf(
			"%sCPU:\x1b[0m %s%s\x1b[0m",
			labelcolor,
			infocolor,
			sys.CPU,
		),

		"memory": fmt.Sprintf(
			"%sMemory:\x1b[0m %s%s\x1b[0m",
			labelcolor,
			infocolor,
			sys.Memory,
		),

		"localip": fmt.Sprintf(
			"%sLocal IP (%s):\x1b[0m %s%s\x1b[0m",
			labelcolor,
			sys.IPVersion,
			infocolor,
			sys.LocalIP,
		),

		"uptime": fmt.Sprintf(
			"%sUptime:\x1b[0m %s%s\x1b[0m",
			labelcolor,
			infocolor,
			sys.Uptime,
		),

		"battery": fmt.Sprintf(
			"%sBattery:\x1b[0m %s%d%% [%s]\x1b[0m",
			labelcolor,
			infocolor,
			sys.Battery,
			sys.BatteryState,
		),

		"de": fmt.Sprintf(
			"%sDE:\x1b[0m %s%s (%s)\x1b[0m",
			labelcolor,
			infocolor,
			sys.DE,
			sys.SessionType,
		),

		"shell": fmt.Sprintf(
			"%sShell:\x1b[0m %s%s\x1b[0m",
			labelcolor,
			infocolor,
			sys.Shell,
		),
	}

	infoLines := []string{
		userInfo,
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
