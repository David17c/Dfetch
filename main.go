package main

import (
	"dfetch/internal/assets"
	"dfetch/internal/config"
	"dfetch/internal/customization"
	"dfetch/internal/model"
	"dfetch/internal/render"
)

func main() {
	lines, asciicolor, headercolor, infocolor, labelcolor := config.ReadConfig()

	sys := model.CollectSystemInfo()

	asciiLines, asciicolor := assets.LoadASCII(
		assets.LogoFS,
		sys.ID,
		asciicolor,
	)

	asciicolor = customization.GetColorCode(asciicolor)
	headercolor = customization.GetColorCode(headercolor)
	infocolor = customization.GetColorCode(infocolor)
	labelcolor = customization.GetColorCode(labelcolor)

	infoLines := render.BuildInfoLines(
		sys,
		lines,
		headercolor,
		infocolor,
		labelcolor,
	)

	render.PrintOutput(asciiLines, infoLines, asciicolor)
}
