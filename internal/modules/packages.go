package modules

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type PackageManager struct {
	Name  string
	Path  string
	Dir   bool
	Count func() int
}

var PackageManagers = []PackageManager{
	{
		Name:  "dpkg",
		Path:  "/var/lib/dpkg/info",
		Dir:   true,
		Count: countDpkg,
	},
	{
		Name:  "pacman",
		Path:  "/var/lib/pacman/local",
		Dir:   true,
		Count: countPacman,
	},
	{
		Name:  "apk",
		Path:  "/lib/apk/db/installed",
		Dir:   false,
		Count: countApk,
	},
	{
		Name:  "eopkg",
		Path:  "/var/lib/eopkg/package",
		Dir:   true,
		Count: countEopkg,
	},
	{
		Name:  "rpm",
		Path:  "/var/lib/rpm",
		Dir:   true,
		Count: countRpm,
	},
	{
		Name:  "snap",
		Path:  "/var/lib/snapd/snaps",
		Dir:   true,
		Count: countSnap,
	},
	{
		Name:  "flatpak",
		Path:  "/var/lib/flatpak/app",
		Dir:   true,
		Count: countFlatpak,
	},
}

func Packages(format string) string {
	var results []string
	total := 0

	for _, pm := range getPackageManagers() {
		count := pm.Count()

		if count > 0 {
			total += count
			results = append(results, fmt.Sprintf("%s %d", pm.Name, count))
		}
	}

	if strings.ToLower(format) == "short" {
		return fmt.Sprintf("%d", total)
	}

	return strings.Join(results, ", ")
}

func getPackageManagers() []PackageManager {
	var detected []PackageManager

	for _, pm := range PackageManagers {
		switch pm.Name {
		case "rpm", "dpkg", "pacman", "apk", "eopkg", "flatpak", "snap":
			if _, err := exec.LookPath(pm.Name); err != nil {
				continue
			}
		}

		if pm.Dir {
			if dirExists(pm.Path) {
				detected = append(detected, pm)
			}
		} else {
			if fileExists(pm.Path) {
				detected = append(detected, pm)
			}
		}
	}

	return detected
}

func countDpkg() int {
	data, err := os.ReadFile("/var/lib/dpkg/status")
	if err != nil {
		return 0
	}

	return bytes.Count(data, []byte("Status: install ok installed"))
}

func countPacman() int {
	entries, err := os.ReadDir("/var/lib/pacman/local")
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "ALPM_DB_VERSION" {
			count++
		}
	}
	return count
}

func countApk() int {
	data, err := os.ReadFile("/lib/apk/db/installed")
	if err != nil {
		return 0
	}

	count := 0
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "P:") {
			count++
		}
	}
	return count
}

func countEopkg() int {
	entries, err := os.ReadDir("/var/lib/eopkg/package")
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			count++
		}
	}
	return count
}

func countRpm() int {
	if _, err := os.Stat("/var/lib/rpm/Packages"); err != nil {
		return 0
	}

	out, err := exec.Command("rpm", "-qa").Output()
	if err != nil {
		return 0
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return 0
	}

	return len(lines)
}

func countSnap() int {
	entries, err := os.ReadDir("/var/lib/snapd/snaps")
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".snap") {
			count++
		}
	}
	return count
}

func countFlatpak() int {
	count := 0
	paths := []string{"/var/lib/flatpak/app"}

	if homeDir, err := os.UserHomeDir(); err == nil {
		paths = append(paths, filepath.Join(homeDir, ".local/share/flatpak/app"))
	}

	seen := make(map[string]bool)

	for _, p := range paths {
		entries, err := os.ReadDir(p)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() && !seen[entry.Name()] {
				seen[entry.Name()] = true
				count++
			}
		}
	}

	return count
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
