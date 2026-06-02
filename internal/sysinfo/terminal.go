package sysinfo

import (
	"fmt"
	"os"
)

func Terminal() string {
	term := os.Getenv("TERM")
	colorterm := os.Getenv("COLORTERM")

	if term == "" {
		return "unknown"
	}
	if colorterm == "" {
		return term
	}

	return fmt.Sprintf("%s [%s]", term, colorterm)
}
