package getsysinfo

import (
	"os"
)

func Hostname() string {
	if hostname, err := os.Hostname(); err == nil && hostname != "" {
		return hostname
	}

	if hostname := tryenv(); hostname != "" {
		return hostname
	}

	return "unknown"
}

func tryenv() string {
	for _, key := range []string{
		"HOST",
		"HOSTNAME",
	} {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	return ""
}
