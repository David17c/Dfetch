package main

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	NoASCII        bool
	ResetConfig    bool
	ListAllModules bool
}

func ParseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.NoASCII, "no-ascii", false, "Disable ASCII art")
	flag.BoolVar(&flags.ResetConfig, "reset-config", false, "Regenerate the config file")
	flag.BoolVar(&flags.ListAllModules, "list-all-modules", false, "Print a list of available modules")

	flag.Parse()

	if flags.ListAllModules {
		ListAllModules()
	}

	return flags
}

func ListAllModules() {
	fmt.Print(`| Module         | Description                           |
| -------------- | ------------------------------------- |
| userinfo       | Username and hostname                 |
| os             | Operating system                      |
| kernel         | Current kernel                        |
| cpu            | Processor information                 |
| memory         | Memory usage                          |
| swap           | Swap usage                            |
| local_ip       | Local IP address                      |
| locale         | System locale settings                |
| uptime         | System uptime                         |
| battery        | Battery information                   |
| bios           | BIOS information                      |
| desktop        | Desktop environment or window manager |
| shell          | Current shell                         |
| terminal       | Current terminal                      |
| disk           | Disk usage                            |
| time           | Current time                          |
| date           | Current date                          |
| packages       | Installed packages                    |
| host           | Device model                          |
| motherboard    | Motherboard name                      |
| emptyline      | Inserts a blank line                  |
| text           | Custom text                           |
`)
	os.Exit(0)
}
