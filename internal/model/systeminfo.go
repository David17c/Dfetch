package model

import (
	sysinfo "dfetch/internal/sysinfo"
)

type SystemInfo struct {
	DistroName   string
	ID           string
	Kernel       string
	CPU          string
	Memory       string
	Username     string
	Hostname     string
	LocalIP      string
	Uptime       string
	Battery      int
	BatteryState string
	DE           string
	SessionType  string
	Shell        string
}

func CollectSystemInfo(enabledModules []string) SystemInfo {
	var sys SystemInfo

	// Always collected
	sys.DistroName, sys.ID = sysinfo.Distro()
	sys.Username = sysinfo.Username()
	sys.Hostname = sysinfo.Hostname()

	// Optional modules
	for _, module := range enabledModules {
		switch module {

		case "kernel":
			sys.Kernel = sysinfo.Kernel()

		case "cpu":
			sys.CPU = sysinfo.Cpu()

		case "memory":
			sys.Memory = sysinfo.Memory()

		case "localip":
			sys.LocalIP = sysinfo.LocalIP()

		case "uptime":
			sys.Uptime = sysinfo.Uptime()

		case "battery":
			sys.Battery, sys.BatteryState = sysinfo.Battery()

		case "de":
			sys.DE, sys.SessionType = sysinfo.DesktopEnvironment()

		case "shell":
			sys.Shell = sysinfo.Shell()
		}
	}

	return sys
}
