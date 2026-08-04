package cmd

import (
	"fmt"
	"os"
)

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
