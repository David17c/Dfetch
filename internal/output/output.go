package output

import (
	"dfetch/internal/sysinfo"
	"fmt"
	"regexp"
	"strings"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(s string) int {
	return len(ansiRegex.ReplaceAllString(s, ""))
}

func BuildInfoLines(sys sysinfo.SystemInfo, configLines []string, accent string) []string {
	info := map[string]string{
		"os":       field(accent, "OS", sys.DistroName),
		"kernel":   field(accent, "Kernel", sys.Kernel),
		"cpu":      field(accent, "CPU", sys.CPU),
		"memory":   field(accent, "RAM", sys.Memory),
		"localip":  field(accent, "IP", sys.LocalIP),
		"uptime":   field(accent, "Uptime", sys.Uptime),
		"shell":    field(accent, "Shell", sys.Shell),
		"terminal": field(accent, "Terminal", sys.Terminal),
		"battery":  field(accent, "Battery", sys.Battery),
		"de":       field(accent, "DE", sys.DE),
		"disk":     field(accent, "Disk", sys.Disk),
		"time":     field(accent, "Time", sys.Time),
		"date":     field(accent, "Date", sys.Date),
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

		fmt.Printf("%s%-*s\x1b[0m %s\n", asciiColor, width, left, right)
	}
}

func getMaxWidth(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		if w := visibleLen(line); w > maxWidth {
			maxWidth = w + 4
		}
	}
	return maxWidth
}
