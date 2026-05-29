package output

import (
	"dfetch/internal/model"
	"fmt"
	"regexp"
	"strings"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(s string) int {
	return len(ansiRegex.ReplaceAllString(s, ""))
}

func BuildInfoLines(sys model.SystemInfo, configLines []string, accent string) []string {
	info := map[string]string{
		"os":      field(accent, "OS", sys.DistroName),
		"kernel":  field(accent, "Kernel", sys.Kernel),
		"cpu":     field(accent, "CPU", sys.CPU),
		"memory":  field(accent, "Memory", sys.Memory),
		"localip": field(accent, "IP", sys.LocalIP),
		"uptime":  field(accent, "Uptime", sys.Uptime),
		"shell":   field(accent, "Shell", sys.Shell),
		"battery": fmt.Sprintf("%sBattery:\x1b[0m %d%% [%s]", accent, sys.Battery, sys.BatteryState),
		"de":      fmt.Sprintf("%sDE:\x1b[0m %s (%s)", accent, sys.DE, sys.SessionType),
	}

	lines := []string{
		fmt.Sprintf("%s%s@%s\x1b[0m", accent, sys.Username, sys.Hostname),
	}

	for _, key := range configLines {
		if v, ok := info[strings.ToLower(strings.TrimSpace(key))]; ok {
			lines = append(lines, v)
		}
	}

	return lines
}

func field(color, label, value string) string {
	return fmt.Sprintf("%s%s:\x1b[0m %s", color, label, value)
}

func PrintOutput(asciiLines, infoLines []string, asciiColor string) {
	width := getMaxWidth(asciiLines)

	total := max(len(asciiLines), len(infoLines))

	for i := 0; i < total; i++ {
		var left, right string

		if i < len(asciiLines) {
			left = asciiLines[i]
		}
		if i < len(infoLines) {
			right = infoLines[i]
		}

		fmt.Printf("\x1b[1m%s%-*s\x1b[0m %s\n", asciiColor, width, left, right)
	}
}

func getMaxWidth(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		if w := visibleLen(line); w > maxWidth {
			maxWidth = w
		}
	}
	return maxWidth
}
