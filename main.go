package main

import (
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
	"os"
	"fmt"
)

func main() {
	// Create the default config file if it doesn't exist
	if err := config.CreateConfigFile(); err != nil {
	    fmt.Fprintln(os.Stderr, err)
    	os.Exit(1)
	}

	// Read the config
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
    	os.Exit(1)
	}

	// Get distro name and ID
	distroName, id := modules.Distro("")

	// Collect system information based on the configured modules
	sys := modules.CollectSystemInfo(cfg.Modules, distroName)

	// Load the ASCII art
	asciiLines := output.LoadASCII(output.LogoFS, id, cfg)

	// Build the output lines
	infoLines := output.BuildInfoLines(sys)

	// Print final result
	output.PrintOutput(asciiLines, infoLines)
}
