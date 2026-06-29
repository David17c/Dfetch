package modules

import "fmt"

type Modules struct {
	DistroName  string
	ID          string
	Kernel      string
	CPU         string
	RAM         string
	Swap        string
	Userinfo    string
	Local_IP    string
	Uptime      string
	Battery     string
	Desktop     string
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

		case "ram":
			sys.RAM = RAM()

		case "local_ip":
			sys.Local_IP = Local_IP()

		case "uptime":
			sys.Uptime = Uptime()

		case "battery":
			sys.Battery = Battery()

		case "desktop":
			sys.Desktop = DesktopEnvironment()

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

		case "swap":
			sys.Swap = Swap()

		default:
			fmt.Printf("Unable to find module '%s'\n", module)
		}

	}

	return sys
}
