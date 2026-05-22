package main

import (
	"dfetch/internal/assets"
	"dfetch/internal/config"
	"dfetch/internal/customization"
	"dfetch/internal/model"
	"dfetch/internal/render"
)

func main() {
	lines, color := config.ReadConfig()

	sys := model.CollectSystemInfo()

	asciiLines, color := assets.LoadASCII(
		assets.LogoFS,
		sys.ID,
		color,
	)

	color = customization.GetColorCode(color)

	infoLines := render.BuildInfoLines(sys, lines)

	render.PrintOutput(asciiLines, infoLines, color)
}
