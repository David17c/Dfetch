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

	// Read ASCII art
	file := fmt.Sprintf("/home/david/Documents/Programmeer-projecten/Dfetch/logo/%s.txt", strings.ToLower(ID))
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	asciiLines := strings.Split(string(data), "\n")

	// Prepare system info lines
	infoLines := []string{
		fmt.Sprintf("OS: %s", prettyName),
		fmt.Sprintf("Kernel: %s", kernel),
		fmt.Sprintf("CPU: %s", cpu),
		fmt.Sprintf("Memory: %s", mem),
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
