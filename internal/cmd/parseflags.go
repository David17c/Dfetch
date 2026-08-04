package cmd

import (
	"flag"
)

type Flags struct {
	NoASCII        bool
	ResetConfig    bool
	ListAllModules bool
	SetASCII       string
}

func ParseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.NoASCII, "no-ascii", false, "Disable ASCII art")
	flag.BoolVar(&flags.ResetConfig, "reset-config", false, "Regenerate the config file")
	flag.BoolVar(&flags.ListAllModules, "list-all-modules", false, "Print a list of available modules")
	flag.StringVar(&flags.SetASCII, "set-ascii", "", "Regenerate the config file")

	flag.Parse()

	if flags.ListAllModules {
		ListAllModules()
	}

	return flags
}
