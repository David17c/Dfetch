package cmd

import (
	"fmt"
	"os"

	flags "github.com/spf13/pflag"
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

	flags.BoolVar(&opts.NoLogo, "no-logo", false, "Disable ASCII art for this run")
	flags.BoolVar(&opts.NoColor, "no-color", false, "Disable color output for this run")
	flags.StringVar(&opts.SetLogo, "set-logo", "", "Set the logo from a custom ASCII art file or distro name")
	flags.StringVar(&opts.SetConfig, "set-config", "", "Use the specified config file")
	flags.BoolVar(&opts.ResetConfig, "reset-config", false, "Regenerate the default config file")
	flags.BoolVar(&opts.ListModules, "list-modules", false, "List all available modules")

	flags.Parse()

	if opts.ListModules {
		ListAllModules()
	}

	return opts
}

func ListAllModules() {
	fmt.Print(`
 Module       Description                           
 ------------ ----------------------------------
 userinfo     Username and hostname                 
 os           Operating system                      
 kernel       Current kernel                        
 cpu          Processor information                 
 memory       Memory usage                          
 swap         Swap usage                            
 local_ip     Local IP address                      
 locale       Systems locale settings               
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
 color        Display terminals color pallet        

`)
	os.Exit(0)
}
