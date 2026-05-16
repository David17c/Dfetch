// main.go
package main

import (
	"Dfetch/customization"
	"Dfetch/getsysinfo"
	"bufio"
	"embed"
	"fmt"
	"strings"
)

//go:embed logo/*
var logoFS embed.FS

type SystemInfo struct {
	PrettyName   string
	ID           string
	Kernel       string
	CPU          string
	Memory       string
	Username     string
	Hostname     string
	LocalIP      string
	IPVersion    string
	Uptime       string
	Battery      int
	BatteryState string
}

func main() {
	lines, color, err := customization.ConfigFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	sys := collectSystemInfo()

	asciiLines, color := loadASCII(sys.ID, color)

	color = customization.GetColorCode(color)

	infoLines := buildInfoLines(sys, lines)

	printOutput(asciiLines, infoLines, color)
}

func collectSystemInfo() SystemInfo {
	prettyName, id := getsysinfo.Distro()
	localIP, version := getsysinfo.LocalIP()
	battery, batteryStatus := getsysinfo.Battery()

	return SystemInfo{
		PrettyName:   prettyName,
		ID:           id,
		Kernel:       getsysinfo.Kernel(),
		CPU:          getsysinfo.Cpu(),
		Memory:       getsysinfo.Mem(),
		Username:     getsysinfo.Username(),
		Hostname:     getsysinfo.Hostname(),
		LocalIP:      localIP,
		IPVersion:    version,
		Uptime:       getsysinfo.Uptime(),
		Battery:      battery,
		BatteryState: batteryStatus,
	}
}

func loadASCII(distroID, color string) ([]string, string) {
	file := fmt.Sprintf("logo/%s.txt", strings.ToLower(distroID))

	f, err := logoFS.Open(file)
	if err != nil {
		f, err = logoFS.Open("logo/linux.txt")
		if err != nil {
			return []string{}, color
		}
	}

	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		// Load color from ASCII file if config color is empty or false
		if strings.HasPrefix(line, "color:") {
			if color == "" {
				color = strings.TrimSpace(strings.TrimPrefix(line, "color:"))
			}
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, color
	}

	return lines, color
}

func buildInfoLines(sys SystemInfo, configLines []string) []string {
	userInfo := fmt.Sprintf("\x1b[1m%s@%s\x1b[0m", sys.Username, sys.Hostname)
	separator := strings.Repeat("─", len(userInfo))

	infoMap := map[string]string{
		"os":      fmt.Sprintf("OS: %s", sys.PrettyName),
		"kernel":  fmt.Sprintf("Kernel: %s", sys.Kernel),
		"cpu":     fmt.Sprintf("CPU: %s", sys.CPU),
		"memory":  fmt.Sprintf("Memory: %s", sys.Memory),
		"ip":      fmt.Sprintf("Local IP (%s): %s", sys.IPVersion, sys.LocalIP),
		"uptime":  fmt.Sprintf("Uptime: %s", sys.Uptime),
		"battery": fmt.Sprintf("Battery: %d%% [%s]", sys.Battery, sys.BatteryState),
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

func printOutput(asciiLines, infoLines []string, color string) {
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

		fmt.Printf("\x1b[1m%s%-*s\x1b[0m %s\n", color, maxLen, left, right)
	}
}

func getMaxWidth(lines []string) int {
	maxLen := 0

	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	return maxLen
}
