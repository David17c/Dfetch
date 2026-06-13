package modules

import "fmt"

type Modules struct {
	DistroName  string
	ID          string
	Kernel      string
	CPU         string
	Memory      string
	Userinfo    string
	LocalIP     string
	Uptime      string
	Battery     string
	DE          string
	Shell       string
	Terminal    string
	colorterm   string
	Disk        string
	Time        string
	Date        string
	Emptyline   string
	Packages    string
	Host        string
	MotherBoard string
}

func CollectSystemInfo(enabledModules []string) Modules {
	var sys Modules

	sys.DistroName, sys.ID = Distro()

	for _, module := range enabledModules {
		switch module {

		case "userinfo":
			sys.Userinfo = Userinfo()

		case "os":
			continue

		case "kernel":
			sys.Kernel = Kernel()

		case "cpu":
			sys.CPU = Cpu()

		case "memory":
			sys.Memory = Ram()

		case "localip":
			sys.LocalIP = LocalIP()

		case "uptime":
			sys.Uptime = Uptime()

		case "battery":
			sys.Battery = Battery()

		case "de":
			sys.DE = DesktopEnvironment()

		case "shell":
			sys.Shell = Shell()

		case "terminal":
			sys.Terminal = Terminal()

		case "disk":
			sys.Disk = Disk()

		case "time":
			sys.Time = Time()

		case "date":
			sys.Date = Date()

		case "emptyline":
			sys.Emptyline = ""

		case "packages":
			sys.Packages = Packages()

		case "host":
			sys.Host = Host()

		case "motherboard":
			sys.MotherBoard = MotherBoard()

		default:
			fmt.Printf("Unable to find module '%s'\n", module)
		}

	}

	return sys
}
