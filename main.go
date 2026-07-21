package main

import (
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
	"fmt"
	"log"
	"os"
)

var version = "dev"

func main() {

	// Print version number on request
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println(version)
		return
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
