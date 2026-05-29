package sysinfo

import "os"

func DesktopEnvironment() (string, string) {
	var DE string
	var sessionType string

	for _, key := range []string{
		"DESKTOP_SESSION",
		"GDMSESSION",
		"XDG_CURRENT_DESKTOP",
	} {
		if value := os.Getenv(key); value != "" {
			DE = value
			break
		}
	}

	for _, key := range []string{
		"XDG_SESSION_TYPE",
		"WAYLAND_DISPLAY",
		"DISPLAY",
	} {
		if value := os.Getenv(key); value != "" {
			sessionType = value
			break
		}
	}

	return DE, sessionType
}
