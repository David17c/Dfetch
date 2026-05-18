package main

import (
	"Dfetch/internal/config"
	"Dfetch/internal/customization"
	"fmt"
	"os"
)

func main() {
	lines, color := config.ReadConfig()

	noColor := false

	if len(os.Args) > 1 {
		if os.Args[1] == "--no-color" {
			noColor = true
			color = ""
		} else {
			fmt.Println("Invalid flag!")
			return
		}
	}

	sys := collectSystemInfo()

	asciiLines, color := loadASCII(sys.ID, color, noColor)

	color = customization.GetColorCode(color)

	infoLines := buildInfoLines(sys, lines)

	printOutput(asciiLines, infoLines, color)
}
