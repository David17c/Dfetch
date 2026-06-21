package modules

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Packages() string {
	var (
		cmd  *exec.Cmd
		name string
	)

	switch {
	case exists("dpkg-query"):
		name = "dpkg"
		cmd = exec.Command("dpkg-query", "-f", "${binary:Package}\n", "-W")

	case exists("rpm"):
		name = "rpm"
		cmd = exec.Command("rpm", "-qa")

	case exists("pacman"):
		name = "pacman"
		cmd = exec.Command("pacman", "-Qq")

	case exists("apk"):
		name = "apk"
		cmd = exec.Command("apk", "info")

	case exists("xbps-query"):
		name = "xbps"
		cmd = exec.Command("xbps-query", "-l")

	case exists("eopkg"):
		name = "eopkg"
		cmd = exec.Command("eopkg", "list-installed")

	case exists("pkg"):
		name = "pkg"
		cmd = exec.Command("pkg", "info")

	case exists("pkg_info"):
		name = "pkg_info"
		cmd = exec.Command("pkg_info")

	default:
		return "Unknown package manager"
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "unknown"
	}

	count := len(strings.Fields(out.String()))
	return fmt.Sprintf("%d (%s)", count, name)
}

func exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
