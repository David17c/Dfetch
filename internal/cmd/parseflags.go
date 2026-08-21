package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Options struct {
	NoLogo      bool
	NoColor     bool
	ResetConfig bool
	ListModules bool
	SetLogo     string
	SetConfig   string
}

func ParseFlags() Options {
	var opts Options

	fs := pflag.NewFlagSet("Dfetch", pflag.ExitOnError)

	fs.BoolVarP(&opts.NoLogo, "no-logo", "l", false, "Disable ASCII art for this run")
	fs.BoolVarP(&opts.NoColor, "no-color", "C", false, "Disable color output for this run")
	fs.StringVarP(&opts.SetLogo, "set-logo", "L", "", "Set the logo from a custom ASCII art file or distro name")
	fs.StringVarP(&opts.SetConfig, "set-config", "c", "", "Use the specified config file")
	fs.BoolVar(&opts.ResetConfig, "regen-config", false, "Regenerate the default config file")
	fs.BoolVarP(&opts.ListModules, "list-modules", "m", false, "List all available modules")

	fs.Parse(os.Args[1:])

	if opts.ListModules {
		listModules()
		os.Exit(0)
	}

	return opts
}

func listModules() {
	fmt.Print(`
 Module       Description
 ------------ -------------------------------
 userinfo     Username and hostname
 os           Operating system
 kernel       Current kernel
 cpu          Processor information
 memory       Memory usage
 swap         Swap usage
 local_ip     Local IP address
 locale       System locale settings
 uptime       System uptime
 battery      Battery information
 bios         BIOS information
 de           Desktop environment
 wm           Window manager
 shell        Current shell
 terminal     Current terminal
 disk         Disk usage
 time         Current time
 date         Current date
 packages     Installed packages
 host         Device model
 board        Motherboard name
 emptyline    Inserts a blank line
 text         Custom text
 color        Display terminal color palette
`)
}
