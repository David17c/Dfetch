package modules

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func Cpu(format string) string {
	file, err := os.Open("/proc/cpuinfo")
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if cpu, ok := strings.CutPrefix(scanner.Text(), "model name\t: "); ok {
				if format == "short" {
					return cleanCPUName(cpu)
				}
				return cpu
			}
		}
	}

	out, err := exec.Command("lscpu").Output()
	if err != nil {
		return "unknown"
	}

	for _, line := range strings.Split(string(out), "\n") {
		if model, ok := strings.CutPrefix(line, "Model name:"); ok {
			model = strings.TrimSpace(model)

			if format == "short" {
				return cleanCPUName(model)
			}
			return model
		}
	}

	return "unknown"
}

func cleanCPUName(name string) string {
	replacements := []string{
		"(R)", "",
		"(TM)", "",
		" CPU", "",
		" Processor", "",
		" APU", "",
		" with Radeon Graphics", "",
		" with Radeon Vega Graphics", "",
		" with Radeon", "",
	}

	for i := 0; i < len(replacements); i += 2 {
		name = strings.ReplaceAll(name, replacements[i], replacements[i+1])
	}

	// Remove everything after " @ " (clock speed).
	if i := strings.Index(name, " @ "); i != -1 {
		name = name[:i]
	}

	// Remove everything after " w/" or " W/".
	lower := strings.ToLower(name)
	if i := strings.Index(lower, " w/"); i != -1 {
		name = name[:i]
	}

	// Collapse duplicate whitespace.
	return strings.Join(strings.Fields(name), " ")
}
