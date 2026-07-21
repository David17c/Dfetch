package modules

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func Kernel() string {
	var uts unix.Utsname

	if err := unix.Uname(&uts); err != nil {
		return "unknown"
	}

	return fmt.Sprintf("%s %s",
		charsToString(uts.Sysname[:]),
		charsToString(uts.Release[:]),
	)
}

func charsToString(ca []byte) string {
	n := 0
	for n < len(ca) && ca[n] != 0 {
		n++
	}
	return string(ca[:n])
}
