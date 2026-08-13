package main

import (
	"fmt"
	"os"

	"dfetch/internal/cmd"
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
)

func main() {
	// Parse command-line flags
	opts := cmd.ParseFlags()

	// Create the default config file if it doesn't exist
	if err := config.CreateConfigFile(opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Read the config
	cfg, err := config.ReadConfig(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Get distro information
	distro, err := modules.Distro()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Collect system information based on the configured modules
	sys := modules.CollectSystemInfo(cfg.Modules, distro.DisplayName(), opts.NoColor)

	// Load the ASCII art
	asciiLines := output.LoadASCII(distro, cfg, opts)

	// Build the output lines
	infoLines := output.BuildInfoLines(sys, opts.NoColor)

	// Print final result
	output.PrintOutput(asciiLines, infoLines, opts.NoColor)
}
