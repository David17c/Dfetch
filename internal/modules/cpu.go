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
					return normalizeCPUName(cpu)
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
				return normalizeCPUName(model)
			}
			return model
		}
	}

	return "unknown"
}

var cpuReplacer = strings.NewReplacer(
    "(R)", "",
    "(TM)", "",
    "®", "",
    "™", "",
    " CPU", "",
    " Processor", "",
    " APU", "",
    " with Radeon Vega Graphics", "",
    " with Radeon Graphics", "",
    " with Radeon", "",
)

func normalizeCPUName(name string) string {
    name = cpuReplacer.Replace(name)

    if before, _, found := strings.Cut(name, " @ "); found {
        name = before
    }

    if before, _, found := strings.Cut(strings.ToLower(name), " w/"); found {
        name = name[:len(before)]
    }

    return strings.Join(strings.Fields(name), " ")
}