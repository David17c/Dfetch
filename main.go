package main

import (
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
	"log"
)

func main() {
	// Create the default config if it doesn't exist.
	if err := config.CreateConfigFile(); err != nil {
		log.Fatal(err)
	}

	// Read the config.
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Get distro name and ID.
	distroName, id := modules.Distro("")

	// Collect system information based on the configured modules.
	sys := modules.CollectSystemInfo(cfg.Modules, distroName)

	// Load the ASCII art.
	asciiLines := output.LoadASCII(output.LogoFS, id, cfg)

	// Build the output lines.
	infoLines := output.BuildInfoLines(sys)

	// Print everything.
	output.PrintOutput(asciiLines, infoLines)

	_ = distroName
}
