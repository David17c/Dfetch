package main

import (
	"dfetch/internal/assets"
	"dfetch/internal/config"
	"dfetch/internal/customization"
	"dfetch/internal/model"
	"dfetch/internal/render"
)

func main() {

	// Read / create the config file
	lines, asciicolor, headercolor, infocolor, labelcolor := config.ReadConfig()

	// Collect the users system info
	sys := model.CollectSystemInfo()

	// Load and format the ascii art
	asciiLines, asciicolor := assets.LoadASCII(
		assets.LogoFS,
		sys.ID,
		asciicolor,
	)

	if labelcolor == "" || labelcolor == "default" {
		labelcolor = asciicolor
	}

	if headercolor == "" || headercolor == "default" {
		headercolor = asciicolor
	}

	// Get the colors corresponding ascii codes
	asciicolor = customization.GetColorCode(asciicolor)
	headercolor = customization.GetColorCode(headercolor)
	infocolor = customization.GetColorCode(infocolor)
	labelcolor = customization.GetColorCode(labelcolor)

	// Build the info lines
	infoLines := render.BuildInfoLines(
		sys,
		lines,
		headercolor,
		infocolor,
		labelcolor,
	)

	// Put everything together and print it
	render.PrintOutput(asciiLines, infoLines, asciicolor)
}
