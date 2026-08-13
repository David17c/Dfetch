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
		fmt.Fprintf(os.Stderr, "Dfetch: %v\n", err)
		os.Exit(1)
	}

	// Read / create the config
	cfg, err := config.ReadConfig(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dfetch: %v\n", err)
		os.Exit(1)
	}

	// Collect distro info
	distro, err := modules.Distro()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dfetch: %v\n", err)
		os.Exit(1)
	}

	// Collect required system info
	sys := modules.CollectSystemInfo(cfg.Modules, distro.DisplayName(), opts.NoColor)

	// Load ascii art
	asciiLines := output.LoadLogo(distro, cfg, opts)

	// Build output lines
	infoLines := output.BuildInfoLines(sys, opts.NoColor)

	// Print result
	output.PrintOutput(asciiLines, infoLines, opts.NoColor)
}
