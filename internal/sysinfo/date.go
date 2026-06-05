package sysinfo

import "time"

func Date() string {
	return time.Now().Format("01-02-2006")
}
