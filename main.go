package main

import (
	"dfetch/internal/config"
	"dfetch/internal/output"
	"dfetch/internal/sysinfo"
)

func main() {

	// Read or create the config file
	enabledModules, asciicolor, accentcolor := config.ReadConfig()

	// Collect necessary system info
	sys := sysinfo.CollectSystemInfo(enabledModules)

	// Prepare the ASCII art
	asciiLines, asciicolor := output.LoadASCII(
		output.LogoFS,
		sys.ID,
		asciicolor,
	)

	if accentcolor == "" || accentcolor == "default" {
		accentcolor = asciicolor
	}

	// Get the ANSI codes correspondig to the colors
	asciicolor = config.GetColorCode(asciicolor)
	accentcolor = config.GetColorCode(accentcolor)

	// Build the info lines
	infoLines := output.BuildInfoLines(
		sys,
		enabledModules,
		accentcolor,
	)

	// Put everything together and print it
	output.PrintOutput(asciiLines, infoLines, asciicolor)
}
