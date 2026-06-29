package output

import (
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(s string) int {
	clean := ansiRegex.ReplaceAllString(s, "")
	return utf8.RuneCountInString(clean)
}

func BuildInfoLines(sys modules.Modules, enabledModules []string, accent string) []string {
	accent = config.GetColorCode(accent)

	info := map[string]string{
		"userinfo":    fmt.Sprintf("%s%s", accent, sys.Userinfo),
		"os":          field(accent, "OS", sys.DistroName),
		"kernel":      field(accent, "Kernel", sys.Kernel),
		"cpu":         field(accent, "CPU", sys.CPU),
		"ram":         field(accent, "RAM", sys.RAM),
		"swap":        field(accent, "Swap", sys.Swap),
		"local_ip":    field(accent, "Local IP", sys.Local_IP),
		"uptime":      field(accent, "Uptime", sys.Uptime),
		"shell":       field(accent, "Shell", sys.Shell),
		"terminal":    field(accent, "Terminal", sys.Terminal),
		"battery":     field(accent, "Battery", sys.Battery),
		"desktop":     field(accent, "Desktop", sys.Desktop),
		"disk":        field(accent, "Disk", sys.Disk),
		"time":        field(accent, "Time", sys.Time),
		"date":        field(accent, "Date", sys.Date),
		"packages":    field(accent, "Packages", sys.Packages),
		"host":        field(accent, "Host", sys.Host),
		"motherboard": field(accent, "Motherboard", sys.MotherBoard),
		"emptyline":   "",
	}

	lines := make([]string, 0, len(enabledModules))

	for _, key := range enabledModules {
		k := strings.ToLower(strings.TrimSpace(key))
		if v, ok := info[k]; ok {
			lines = append(lines, v)
		}
	}

	return lines
}

func field(color, label, value string) string {
	return fmt.Sprintf("%s%s:\x1b[0m %s", color, label, value)
}

func PrintOutput(asciiLines, infoLines []string) {
	var renderedAscii []string
	renderedAscii = make([]string, len(asciiLines))

	for i, line := range asciiLines {
		renderedAscii[i] = ApplyColorTags(line)
	}

	width := getMaxWidth(renderedAscii)
	total := max(len(asciiLines), len(infoLines))

	for i := 0; i < total; i++ {
		var left, right string

		if i < len(asciiLines) {
			left = ApplyColorTags(asciiLines[i])
		}
		if i < len(infoLines) {
			right = infoLines[i]
		}

		padding := width - visibleLen(left)
		if padding < 0 {
			padding = 0
		}

		fmt.Printf("%s%s %s\x1b[0m\n",
			left,
			strings.Repeat(" ", padding+2),
			right,
		)
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

var colorTagRE = regexp.MustCompile(`\$\{([^}]+)\}`)

func ApplyColorTags(line string) string {
	result := colorTagRE.ReplaceAllStringFunc(line, func(tag string) string {
		name := strings.TrimSuffix(strings.TrimPrefix(tag, "${"), "}")

		if strings.EqualFold(name, "reset") {
			return "\x1b[0m"
		}

		return config.GetColorCode(name)
	})

	return result + "\x1b[0m"
}
