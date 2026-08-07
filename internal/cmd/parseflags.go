package cmd

import (
	flag "github.com/spf13/pflag"
)

type Flags struct {
	NoASCII        bool
	ResetConfig    bool
	ListAllModules bool
	SetASCII       string
	SetConfig      string
}

func ParseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.NoASCII, "no-ascii", false, "Disable ASCII art")
	flag.StringVar(&flags.SetASCII, "set-ascii", "", "Set custom ascii art")
	flag.StringVar(&flags.SetConfig, "set-config", "", "Specify config file to use")
	flag.BoolVar(&flags.ResetConfig, "reset-config", false, "Regenerate the config file")
	flag.BoolVar(&flags.ListAllModules, "list-all-modules", false, "Print a list of available modules")

	flag.Parse()

	if flags.ListAllModules {
		ListAllModules()
	}

	return flags
}
