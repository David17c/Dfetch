package modules

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func Cpu() string {
	file, err := os.Open("/proc/cpuinfo")
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if cpu, ok := strings.CutPrefix(scanner.Text(), "model name\t: "); ok {
				return cpu
			}
		}
	}

	// Fallback that only works with lscpu installed otherwise its skipped
	out, err := exec.Command("lscpu").Output()
	if err != nil {
		return "unknown"
	}

	for _, line := range strings.Split(string(out), "\n") {
		if model, ok := strings.CutPrefix(line, "Model name:"); ok {
			return strings.TrimSpace(model)
		}
	}

	return "unknown"
}
