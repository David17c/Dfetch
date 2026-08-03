package main

import (
	"fmt"
	"os"

	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
)

func main() {

	// Read flags and put them in a struct
	flags := ParseFlags()

	// Remove existing config file if user wants to regenerate it
	if flags.ResetConfig {
		config.RemoveConfigFile()
	}

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
	asciiLines := output.LoadASCII(output.LogoFS, id, cfg, flags.NoASCII)

	// Build the output lines
	infoLines := output.BuildInfoLines(sys)

	// Print final result
	output.PrintOutput(asciiLines, infoLines)
}
