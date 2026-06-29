package modules

import (
	"bytes"
	"fmt"
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

		out, err := exec.Command(pm.bin, pm.args...).Output()
		if err != nil {
			result = "unknown"
			return
		}

		count := bytes.Count(out, []byte{'\n'})

		if len(out) > 0 && out[len(out)-1] != '\n' {
			count++
		}

		result = fmt.Sprintf("%s - %d", pm.name, count)
	})

	return result
}

func getPackageManager() *packageManager {
	detectOnce.Do(func() {
		for i := range packageManagers {
			if exists(packageManagers[i].bin) {
				detected = &packageManagers[i]
				return
			}
		}
	})

	return detected
}

func exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
