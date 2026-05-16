// main.go
package main

import (
	"Dfetch/customization"
	"Dfetch/getsysinfo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	prettyName, ID := getsysinfo.Distro()
	kernel := getsysinfo.Kernel()
	cpu := getsysinfo.Cpu()
	mem := getsysinfo.Mem()
	username := getsysinfo.Username()
	hostname := getsysinfo.Hostname()
	localip, version := getsysinfo.LocalIP()
	uptime := getsysinfo.Uptime()

	// Try to get distro specific ASCII art if that fails use Linux penguin ASCII art if that fails skip ASCII art
	file := fmt.Sprintf(
		"./logo/%s.txt",
		strings.ToLower(ID),
	)

	f, err := os.Open(file)
	if err != nil {
		file = "./logo/linux.txt"
		f, err = os.Open(file)
	}

	var data []string
	var color string

	// Read ASCII art file line by line if line starts with "color:" get the color and continue
	if err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {

			line := scanner.Text()

			if strings.HasPrefix(line, "color:") {
				colorName := strings.TrimSpace(strings.TrimPrefix(line, "color:"))
				color = customization.GetColorCode(colorName)
				continue
			}

			data = append(data, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			data = []string{}
		}
	}

	asciiLines := data

	userinfo := fmt.Sprintf("%s@%s", username, hostname)
	separator := strings.Repeat("-", len(userinfo))

	// Prepare system info lines
	infoLines := []string{
		fmt.Sprintf("%s@%s", username, hostname),
		fmt.Sprintf("%s", separator),
		fmt.Sprintf("OS: %s", prettyName),
		fmt.Sprintf("Kernel: %s", kernel),
		fmt.Sprintf("CPU: %s", cpu),
		fmt.Sprintf("Memory: %s", mem),
		fmt.Sprintf("Local IP (%s): %s", version, localip),
		fmt.Sprintf("Uptime: %s", uptime),
	}

	// Find max width of ASCII art for padding
	maxLen := 0
	for _, line := range asciiLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Print side by side
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

		fmt.Printf("%s%-*s\x1b[0m %s\n", color, maxLen, left, right)
	}
}
