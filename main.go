package main

import (
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
	"log"
)

func main() {
	// Read or create config file
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Collect necessary system info
	sys := modules.CollectSystemInfo(cfg.EnabledModules)

	// Prepare the ASCII art
	asciiLines, accentColor := output.LoadASCII(
		output.LogoFS,
		sys.ID,
		cfg.AsciiSize,
		cfg.CustomAscii,
	)
	if err != nil {
		log.Fatal(err)
	}

	
	if cfg.AccentColor == "" || cfg.AccentColor == "default" {
		cfg.AccentColor = accentColor
	}

	// Build the info lines
	infoLines := output.BuildInfoLines(
		sys,
		cfg.EnabledModules,
		accentColor,
	)

	// Put everything together and print it
	output.PrintOutput(asciiLines, infoLines)
}
