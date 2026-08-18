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

	// Collect distro info
	distro, err := modules.Distro()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dfetch: %v\n", err)
		os.Exit(1)
	}

	// Create the default config file if it doesn't exist
	if err := config.CreateConfigFile(opts, distro.ID, distro.IDLike); err != nil {
		fmt.Fprintf(os.Stderr, "Dfetch: %v\n", err)
		os.Exit(1)
	}

	// Read the config
	cfg, err := config.ReadConfig(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dfetch: %v\n", err)
		os.Exit(1)
	}

	// Collect required system info
	sys := modules.CollectSystemInfo(cfg.Modules, distro.DisplayName(), opts)

	// Load ascii art
	asciiLines := output.LoadLogo(distro, cfg, opts)

	// Build output lines
	infoLines := output.BuildInfoLines(sys, opts.NoColor)

	// Print result
	output.PrintOutput(asciiLines, infoLines, opts.NoColor)
}
