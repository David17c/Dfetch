package modules

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type packageManager struct {
	name string
	bin  string
	args []string
}

var packageManagers = []packageManager{
	{"dpkg", "dpkg-query", []string{"-f", "${binary:Package}\n", "-W"}},
	{"rpm", "rpm", []string{"-qa"}},
	{"pacman", "pacman", []string{"-Qq"}},
	{"apk", "apk", []string{"info"}},
	{"xbps", "xbps-query", []string{"-l"}},
	{"eopkg", "eopkg", []string{"list-installed"}},
	{"pkg", "pkg", []string{"info"}},
	{"pkg_info", "pkg_info", nil},
	{"nix", "nix", []string{"profile", "list"}},
}

var (
	detectOnce sync.Once
	detected   *packageManager

	countOnce sync.Once
	result    string
)

func Packages() string {
	countOnce.Do(func() {
		pm := getPackageManager()
		if pm == nil {
			result = "Unknown package manager"
			return
		}

		var count int
		var err error

		// Special handling for NixOS
		if pm.name == "nix" {
			count, err = countNixPackages()
		} else {
			count, err = countPackagesFromCommand(pm)
		}

		if err != nil {
			result = "unknown"
			return
		}

		result = fmt.Sprintf("%s - %d", pm.name, count)
	})

	return result
}

func countPackagesFromCommand(pm *packageManager) (int, error) {
	out, err := exec.Command(pm.bin, pm.args...).Output()
	if err != nil {
		return 0, err
	}

	count := bytes.Count(out, []byte{'\n'})

	if len(out) > 0 && out[len(out)-1] != '\n' {
		count++
	}

	return count, nil
}

func countNixPackages() (int, error) {
	// Use nix profile list to count installed packages
	out, err := exec.Command("nix", "profile", "list").Output()
	if err != nil {
		return 0, err
	}

	// Count output lines (each line is a package)
	count := bytes.Count(out, []byte{'\n'})

	// If output doesn't end with newline, increment count
	if len(out) > 0 && out[len(out)-1] != '\n' {
		count++
	}

	// Subtract 1 if there's a header line or handle empty output
	if count > 0 {
		count--
	}

	return count, nil
}

func getPackageManager() *packageManager {
	detectOnce.Do(func() {
		// Check for NixOS first via /etc/os-release
		if isNixOS() && exists("nix") {
			detected = &packageManagers[8] // nix package manager
			return
		}

		// Fall back to other package managers
		for i := range packageManagers {
			if i == 8 { // Skip nix here, already handled above
				continue
			}
			if exists(packageManagers[i].bin) {
				detected = &packageManagers[i]
				return
			}
		}
	})

	return detected
}

func isNixOS() bool {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return false
	}
	return bytes.Contains(data, []byte("ID=nixos"))
}

func exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
