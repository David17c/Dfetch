package sysinfo

import (
	"os"
	"os/user"
)

func Username() string {
	currentUser, err := user.Current()
	if err == nil && currentUser.Username != "" {
		return currentUser.Username
	}

	for _, key := range []string{
		"USER",
		"LOGNAME",
		"USERNAME",
	} {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	return "unknown"
}
