package main

import (
	"dfetch/internal/config"
	"dfetch/internal/model"
	"dfetch/internal/render"
)

func main() {

	// Read / create the config file
	lines, asciicolor, accentcolor := config.ReadConfig()

	// Collect the users system info
	sys := model.CollectSystemInfo()

	// Load and format the ascii art
	asciiLines, asciicolor := render.LoadASCII(
		render.LogoFS,
		sys.ID,
		asciicolor,
	)

	if accentcolor == "" || accentcolor == "default" {
		accentcolor = asciicolor
	}

	// Get the colors corresponding ascii codes
	asciicolor = config.GetColorCode(asciicolor)
	accentcolor = config.GetColorCode(accentcolor)

	// Build the info lines
	infoLines := render.BuildInfoLines(
		sys,
		lines,
		accentcolor,
	)

	// Put everything together and print it
	render.PrintOutput(asciiLines, infoLines, asciicolor)
}
