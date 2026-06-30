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

func BuildInfoLines(sys modules.Modules, cfg config.Config, distroName string) []string {
	cfg.LabelColor = config.GetColorCode(cfg.LabelColor)
	cfg.UserinfoColor = config.GetColorCode(cfg.UserinfoColor)
	cfg.Info_color = config.GetColorCode(cfg.Info_color)

	fields := map[string]struct {
		label string
		value string
	}{
		"os":          {"OS", distroName},
		"kernel":      {"Kernel", sys.Kernel},
		"cpu":         {"CPU", sys.CPU},
		"memory":      {"Memory", sys.Memory},
		"swap":        {"Swap", sys.Swap},
		"local_ip":    {"Local IP", sys.Local_IP},
		"uptime":      {"Uptime", sys.Uptime},
		"shell":       {"Shell", sys.Shell},
		"terminal":    {"Terminal", sys.Terminal},
		"battery":     {"Battery", sys.Battery},
		"desktop":     {"Desktop", sys.Desktop},
		"disk":        {"Disk", sys.Disk},
		"time":        {"Time", sys.Time},
		"date":        {"Date", sys.Date},
		"packages":    {"Packages", sys.Packages},
		"host":        {"Host", sys.Host},
		"motherboard": {"Motherboard", sys.MotherBoard},
	}

	info := map[string]string{
		"userinfo":  fmt.Sprintf("%s%s\x1b[0m", cfg.UserinfoColor, sys.Userinfo),
		"emptyline": "",
	}

	for key, f := range fields {
		info[key] = field(cfg.LabelColor, cfg.Info_color, f.label, f.value)
	}

	lines := make([]string, 0, len(cfg.EnabledModules))

	for _, key := range cfg.EnabledModules {
		k := strings.ToLower(strings.TrimSpace(key))
		if v, ok := info[k]; ok {
			lines = append(lines, v)
		}
	}

	return lines
}

func field(labelColor, infoColor, label, value string) string {
	return fmt.Sprintf(
		"%s%s:\x1b[0m %s%s\x1b[0m",
		labelColor,
		label,
		infoColor,
		value,
	)
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
