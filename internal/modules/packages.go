package modules

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"unicode"
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

func Packages(format string) string {
	countOnce.Do(func() {
		pm := getPackageManager()
		if pm == nil {
			result = "Unknown package manager"
			return
		}

		var pmCount int
		var err error

		// Special handling for NixOS
		if pm.name == "nix" {
			pmCount, err = countNixPackages()
		} else {
			pmCount, err = countPackagesFromCommand(pm)
		}

		if err != nil {
			result = "unknown"
			return
		}

		flatpakCount, _ := countFlatpakPackages()
		snapCount, _ := countSnapPackages()

		total := pmCount + flatpakCount + snapCount

		if format == "short" {
			result = fmt.Sprintf("%d", total)
			return
		}

		parts := []string{
			fmt.Sprintf("%d %s", pmCount, pm.name),
		}

		if flatpakCount > 0 {
			parts = append(parts, fmt.Sprintf("%d flatpak", flatpakCount))
		}

		if snapCount > 0 {
			parts = append(parts, fmt.Sprintf("%d snap", snapCount))
		}

		result = fmt.Sprintf("%d (%s)", total, strings.Join(parts, ", "))
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
	profiles := []string{
		"/nix/var/nix/profiles/default",
		"/run/current-system",
	}

	if home, err := os.UserHomeDir(); err == nil {
		profiles = append(profiles,
			filepath.Join(home, ".nix-profile"),
			filepath.Join(userStateDir(home), "nix", "profile"),
		)
	}

	if user := os.Getenv("USER"); user != "" {
		profiles = append(profiles, filepath.Join("/etc/profiles/per-user", user))
	}

	total := 0
	found := false
	seen := make(map[string]struct{})

	for _, profile := range profiles {
		resolved := profile
		if realPath, err := filepath.EvalSymlinks(profile); err == nil {
			resolved = realPath
		}

		if _, ok := seen[resolved]; ok {
			continue
		}
		seen[resolved] = struct{}{}

		count, ok, err := countNixProfilePackages(profile)
		if err != nil {
			return 0, err
		}
		if !ok {
			continue
		}

		found = true
		total += count
	}

	if !found {
		return 0, errors.New("no nix profiles found")
	}

	return total, nil
}

func userStateDir(home string) string {
	if stateHome := os.Getenv("XDG_STATE_HOME"); stateHome != "" {
		return stateHome
	}

	return filepath.Join(home, ".local", "state")
}

func countNixProfilePackages(profile string) (int, bool, error) {
	if info, err := os.Stat(profile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return 0, false, nil
		}
		return 0, false, err
	} else if !info.IsDir() {
		return 0, false, nil
	}

	out, err := exec.Command("nix-store", "--query", "--requisites", profile).Output()
	if err != nil {
		return 0, false, err
	}

	count := 0
	for _, line := range bytes.Split(out, []byte{'\n'}) {
		if isValidNixPackagePath(string(line)) {
			count++
		}
	}

	return count, true, nil
}

func isValidNixPackagePath(path string) bool {
	if path == "" {
		return false
	}

	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return false
	}

	name := filepath.Base(path)
	if strings.HasPrefix(name, "nixos-system-nixos-") ||
		strings.HasSuffix(name, "-doc") ||
		strings.HasSuffix(name, "-man") ||
		strings.HasSuffix(name, "-info") ||
		strings.HasSuffix(name, "-dev") ||
		strings.HasSuffix(name, "-bin") {
		return false
	}

	return containsVersion(name)
}

func containsVersion(name string) bool {
	state := 0
	for _, char := range name {
		switch state {
		case 0:
			if unicode.IsDigit(char) {
				state = 1
			}
		case 1:
			if unicode.IsDigit(char) {
				continue
			}
			if char == '.' {
				state = 2
			} else {
				state = 0
			}
		case 2:
			if unicode.IsDigit(char) {
				state = 3
			} else {
				state = 0
			}
		case 3:
			return true
		}
	}

	return state == 3
}

func getPackageManager() *packageManager {
	detectOnce.Do(func() {
		// Check for NixOS first via /etc/os-release
		if isNixOS() {
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

func countFlatpakPackages() (int, error) {
	if !exists("flatpak") {
		return 0, nil
	}

	out, err := exec.Command("flatpak", "list", "--app").Output()
	if err != nil {
		return 0, err
	}

	count := bytes.Count(out, []byte{'\n'})
	if len(out) > 0 && out[len(out)-1] != '\n' {
		count++
	}

	return count, nil
}

func countSnapPackages() (int, error) {
	if !exists("snap") {
		return 0, nil
	}

	out, err := exec.Command("snap", "list").Output()
	if err != nil {
		return 0, err
	}

	lines := bytes.Count(out, []byte{'\n'})
	if len(out) > 0 && out[len(out)-1] != '\n' {
		lines++
	}

	if lines > 0 {
		lines--
	}

	if lines < 0 {
		lines = 0
	}

	return lines, nil
}
