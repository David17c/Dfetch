package cmd

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type Flags struct {
	NoASCII        bool
	NoColor        bool
	ResetConfig    bool
	ListAllModules bool
	SetASCII       string
	SetConfig      string
}

func ParseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.NoASCII, "no-ascii", false, "Disable ASCII art")
	flag.BoolVar(&flags.NoColor, "no-color", false, "Disable color")
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

func ListAllModules() {
	fmt.Print(`
| Module        | Description                           |
| --------------| --------------------------------------|
| userinfo      | Username and hostname                 |
| os            | Operating system                      |
| kernel        | Current kernel                        |
| cpu           | Processor information                 |
| memory        | Memory usage                          |
| swap          | Swap usage                            |
| local_ip      | Local IP address                      |
| locale        | Systems locale settings               |
| uptime        | System uptime                         |
| battery       | Battery information                   |
| bios          | BIOS information                      |
| de            | Desktop environment                   |
| wm            | Window manager                        |
| shell         | Current shell                         |
| terminal      | Current terminal                      |
| disk          | Disk usage                            |
| time          | Current time                          |
| date          | Current date                          |
| packages      | Installed packages                    |
| host          | Device model                          |
| board         | Motherboard name                      |
| emptyline     | Inserts a blank line                  |
| text          | Custom text                           |

`)
	os.Exit(0)
}
