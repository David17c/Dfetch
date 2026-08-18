package modules

import (
	"bytes"
	"dfetch/internal/format"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
		Count: countRpm,
	},
	{
		Name:  "snap",
		Count: countSnap,
	},
	{
		Name:  "flatpak",
		Path:  "/var/lib/flatpak/app",
		Dir:   true,
		Count: countFlatpak,
	},
}

func Packages(formatstring string) string {
	fields := format.Fields(formatstring)

	needsPackages := false
	needsTotal := false
	neededManagers := make(map[string]bool)

	for _, field := range fields {
		switch field {
		case "packages":
			needsPackages = true

		case "total":
			needsTotal = true

		case "dpkg", "pacman", "apk", "eopkg", "rpm", "snap", "flatpak":
			neededManagers[field] = true
		}
	}

	if needsPackages || needsTotal {
		for _, pm := range PackageManagers {
			neededManagers[pm.Name] = true
		}
	}

	if len(neededManagers) == 0 {
		return format.Format(formatstring, format.Values{})
	}

	detected := getPackageManagers()
	counts := make(map[string]int)

	for _, pm := range detected {
		if !neededManagers[pm.Name] {
			continue
		}

		counts[pm.Name] = pm.Count()
	}

	total := 0

	for _, count := range counts {
		total += count
	}

	values := format.Values{}

	for _, field := range fields {
		switch field {
		case "packages":
			var results []string

			for _, pm := range detected {
				if !neededManagers[pm.Name] {
					continue
				}

				count := counts[pm.Name]
				if count > 0 {
					results = append(
						results,
						fmt.Sprintf("%d %s", count, pm.Name),
					)
				}
			}

			values["packages"] = strings.Join(results, ", ")

		case "total":
			values["total"] = strconv.Itoa(total)

		case "dpkg", "pacman", "apk", "eopkg", "rpm", "snap", "flatpak":
			values[field] = strconv.Itoa(counts[field])
		}
	}

	return format.Format(formatstring, values)
}

func getPackageManagers() []PackageManager {
	var detected []PackageManager

	for _, pm := range PackageManagers {
		switch pm.Name {
		case "rpm", "snap":
			if _, err := exec.LookPath(pm.Name); err != nil {
				continue
			}

			detected = append(detected, pm)
			continue
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
	out, err := exec.Command("rpm", "-qa").Output()
	if err != nil {
		return 0
	}

	return bytes.Count(out, []byte{'\n'})
}

func countSnap() int {
	out, err := exec.Command("snap", "list").Output()
	if err != nil {
		return 0
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) <= 1 {
		return 0
	}

	return len(lines) - 1
}

func countFlatpak() int {
	count := 0
	paths := []string{"/var/lib/flatpak/app"}

	if homeDir, err := os.UserHomeDir(); err == nil {
		paths = append(
			paths,
			filepath.Join(homeDir, ".local/share/flatpak/app"),
		)
	}

	seen := make(map[string]bool)

	for _, path := range paths {
		entries, err := os.ReadDir(path)
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
