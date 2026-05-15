package main

import (
	"Dfetch/getsysinfo"
	"fmt"
	"os"
	"strings"
)

func main() {
	prettyName, ID := getsysinfo.GetDistro()
	kernel := getsysinfo.GetKernel()
	cpu := getsysinfo.GetCpu()
	mem := getsysinfo.GetMem()
	username := getsysinfo.GetUsername()
	hostname := getsysinfo.GetHostname()
	localip, version := getsysinfo.GetLocalIP()

	// Read ASCII art
	file := fmt.Sprintf("/home/david/Documents/Programmeer-projecten/Dfetch/logo/%s.txt", strings.ToLower(ID))
	data, err := os.ReadFile(file)
	if err != nil {
		// fallback to linux.txt
		file = "/home/david/Documents/Programmeer-projecten/Dfetch/logo/linux.txt"
		data, err = os.ReadFile(file)
		if err != nil {
			// if linux.txt also fails, just ignore the ascii art its not that important anyway.
			data = []byte("")
		}
	}

	asciiLines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

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

		fmt.Printf("%-*s   %s\n", maxLen, left, right)
	}
}

func getANSIColor(name string) string {
	switch strings.ToLower(name) {
	case "black":
		return "\033[30m"
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	default:
		return "\033[0m"
	}
}
