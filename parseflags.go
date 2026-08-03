package main

import (
	"flag"
)

type Flags struct {
	NoASCII     bool
	ResetConfig bool
}

func ParseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.NoASCII, "no-ascii", false, "Disable ASCII art")
	flag.BoolVar(&flags.ResetConfig, "reset-config", false, "Regenerate the config file")

	flag.Parse()

	return flags
}
