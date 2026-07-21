package modules

import (
	"os"
	"os/user"
)

func Userinfo() (string, string) {

	hostname := Hostname()
	username := Username()

	return username, hostname
}

func Hostname() string {
	if hostname, err := os.Hostname(); err == nil && hostname != "" {
		return hostname
	}

	for _, key := range []string{
		"HOST",
		"HOSTNAME",
	} {
		if hostname := os.Getenv(key); hostname != "" {
			return hostname
		}
	}

	return "unknown"
}

func Username() string {
	for _, key := range []string{"USER", "LOGNAME", "USERNAME"} {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	if u, err := user.Current(); err == nil && u.Username != "" {
		return u.Username
	}

	return "unknown"
}
