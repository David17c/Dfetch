package getsysinfo

import (
	"os/user"
)

func GetUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		return "unknown"
	}

	return currentUser.Username
}
