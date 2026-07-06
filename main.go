package main

import (
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
	"fmt"
	"log"
	"os"
)

func main() {

	// Version command
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" {
			fmt.Printf("Dfetch 1.2.0\n")
			return
		}
	}

	// Get distro name and id
	distroName, id := modules.Distro()

	// Read or create config file
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Collect necessary system info
	sys := modules.CollectSystemInfo(cfg.EnabledModules)

	// Prepare the ASCII art
	asciiLines := output.LoadASCII(output.LogoFS, id, cfg)

	// Build the info lines
	infoLines := output.BuildInfoLines(sys, *cfg, distroName)

	// Put everything together and print it
	output.PrintOutput(asciiLines, infoLines)
}
