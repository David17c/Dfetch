package modules

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

type utsname struct {
	Sysname    [65]byte
	Nodename   [65]byte
	Release    [65]byte
	Version    [65]byte
	Machine    [65]byte
	Domainname [65]byte
}

func cString(b []byte) string {
	i := 0
	for i < len(b) && b[i] != 0 {
		i++
	}
	return string(b[:i])
}

func Kernel() string {
	b, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err == nil {
		kernel := strings.TrimSpace(string(b))
		if kernel != "" {
			return kernel
		}
	}

	var u utsname

	_, _, errno := syscall.Syscall(
		syscall.SYS_UNAME,
		uintptr(unsafe.Pointer(&u)),
		0,
		0,
	)

	if errno == 0 {
		kernel := cString(u.Release[:])
		if kernel != "" {
			return kernel
		}
	}

	return "unknown"
}
