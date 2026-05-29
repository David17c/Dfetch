package model

import sysinfo "dfetch/internal/sysinfo"

type SystemInfo struct {
	DistroName   string
	ID           string
	Kernel       string
	CPU          string
	Memory       string
	Username     string
	Hostname     string
	LocalIP      string
	IPVersion    string
	Uptime       string
	Battery      int
	BatteryState string
	DE           string
	SessionType  string
	Shell        string
}

func CollectSystemInfo() SystemInfo {
	DistroName, id := sysinfo.Distro()
	localIP := sysinfo.LocalIP()
	battery, batteryStatus := sysinfo.Battery()

	de, sessionType := sysinfo.DesktopEnvironment()

	return SystemInfo{
		DistroName:   DistroName,
		ID:           id,
		Kernel:       sysinfo.Kernel(),
		CPU:          sysinfo.Cpu(),
		Memory:       sysinfo.Memory(),
		Username:     sysinfo.Username(),
		Hostname:     sysinfo.Hostname(),
		LocalIP:      localIP,
		Uptime:       sysinfo.Uptime(),
		Battery:      battery,
		BatteryState: batteryStatus,
		DE:           de,
		SessionType:  sessionType,
		Shell:        sysinfo.Shell(),
	}
}
