package output

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"dfetch/internal/config"
	"dfetch/internal/modules"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(s string) int {
	clean := ansiRegex.ReplaceAllString(s, "")
	return utf8.RuneCountInString(clean)
}

func BuildInfoLines(info []modules.ModuleOutput, NoColor bool) []string {
	lines := make([]string, 0, len(info))

	for _, module := range info {
		if module.Name == "emptyline" {
			lines = append(lines, "")
			continue
		}

		if module.Value == "" {
			continue
		}

		if module.Name == "userinfo" {
			if module.Color != "" {
				lines = append(
					lines,
					config.GetColorCode(module.Color, NoColor)+module.Value+config.GetColorCode("reset", NoColor),
				)
			} else {
				lines = append(lines, module.Value)
			}

			continue
		}

		label := module.Label

		if module.Name == "disk" {
			mount := module.Mount
			if mount == "" {
				mount = "/"
			}

			label = fmt.Sprintf("%s (%s)", label, mount)
		}

		lines = append(lines, field(
			label,
			module.Color,
			module.Separator,
			module.Value,
			NoColor,
		))
	}

	return lines
}

func field(label, color, separator, value string, NoColor bool) string {
	if label == "" && separator == "" {
		return value
	}

	if color == "" {
		return fmt.Sprintf("%s%s %s", label, separator, value)
	}

	return fmt.Sprintf(
		"%s%s\x1b[0m%s %s",
		config.GetColorCode(color, NoColor),
		label,
		separator,
		value,
	)
}

func PrintOutput(asciiLines, infoLines []string, NoColor bool) {
	renderedAscii := make([]string, len(asciiLines))

	for i, line := range asciiLines {
		renderedAscii[i] = ApplyColorTags(line, NoColor)
	}

	if len(renderedAscii) == 0 {
		for _, line := range infoLines {
			fmt.Println(line)
		}
		return
	}

	width := getMaxWidth(renderedAscii)
	total := max(len(asciiLines), len(infoLines))

	for i := 0; i < total; i++ {
		var left string

		if i < len(renderedAscii) {
			left = renderedAscii[i]
		}

		padding := width - visibleLen(left) + 2
		if padding < 0 {
			padding = 0
		}

		if i >= len(infoLines) {
			fmt.Printf("%s\n", left)
			continue
		}

		rightLines := strings.Split(infoLines[i], "\n")

		for j, right := range rightLines {
			if j == 0 {
				fmt.Printf("%s%s%s\x1b[0m\n",
					left,
					strings.Repeat(" ", padding),
					right,
				)
			} else {
				fmt.Printf("%s%s\x1b[0m\n",
					strings.Repeat(" ", width+2),
					right,
				)
			}
		}
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

func ApplyColorTags(line string, NoColor bool) string {

	result := colorTagRE.ReplaceAllStringFunc(line, func(tag string) string {
		name := strings.TrimSuffix(strings.TrimPrefix(tag, "${"), "}")

		if strings.EqualFold(name, "reset") {
			return "\x1b[0m"
		}

		return config.GetColorCode(name, NoColor)
	})

	return result + "\x1b[0m"
}
