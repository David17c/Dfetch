// getsysinfo/getusername.go
package getsysinfo

import (
	"os/user"
)

func Username() string {
	currentUser, err := user.Current()
	if err != nil {
		return "unknown"
	}

	return currentUser.Username
}
