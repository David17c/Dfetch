// getsysinfo/desktopenviroment.go
package getsysinfo

import "os"

func DesktopEnvironment() (string, string) {
	de := os.Getenv("DESKTOP_SESSION")
	if de == "" {
		de = "unknown"
	}

	sessionType := os.Getenv("XDG_SESSION_TYPE")
	if sessionType == "" {
		sessionType = "unknown"
	}

	return de, sessionType
}
