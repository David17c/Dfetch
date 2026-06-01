package sysinfo

type SystemInfo struct {
	DistroName string
	ID         string
	Kernel     string
	CPU        string
	Memory     string
	Username   string
	Hostname   string
	LocalIP    string
	Uptime     string
	Battery    string
	DE         string
	Shell      string
	Terminal   string
	colorterm  string
}

func CollectSystemInfo(enabledModules []string) SystemInfo {
	var sys SystemInfo

	sys.DistroName, sys.ID = Distro()
	sys.Username = Username()
	sys.Hostname = Hostname()

	for _, module := range enabledModules {
		switch module {

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
		}
	}

	return sys
}
