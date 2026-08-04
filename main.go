package main

import (
	"fmt"
	"os"

	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
)

func main() {
	// Parse command-line flags
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

	// Get distro information
	distro, err := modules.Distro()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Collect system information based on the configured modules
	sys := modules.CollectSystemInfo(cfg.Modules, distro.DisplayName())

	// Load the ASCII art
	asciiLines := output.LoadASCII(distro, cfg, flags.NoASCII)

	// Build the output lines
	infoLines := output.BuildInfoLines(sys)

	// Print final result
	output.PrintOutput(asciiLines, infoLines)
}
