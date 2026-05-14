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
		panic(err)
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
