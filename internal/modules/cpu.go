package modules

import (
	"bufio"
	"dfetch/internal/format"
	"os"
	"os/exec"
	"strings"
)

func Cpu(formatstring string) string {
	fields := format.Fields(formatstring)

	needsName := false
	needsShort := false

	for _, field := range fields {
		switch field {
		case "name":
			needsName = true

		case "short":
			needsShort = true
		}
	}

	if !needsName && !needsShort {
		return format.Format(formatstring, format.Values{})
	}

	name := cpuName()
	if name == "" {
		name = "unknown"
	}

	values := format.Values{}

	for _, field := range fields {
		switch field {
		case "name":
			values["name"] = name

		case "short":
			values["short"] = normalizeCPUName(name)
		}
	}

	return format.Format(formatstring, values)
}

func cpuName() string {
	file, err := os.Open("/proc/cpuinfo")
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			if cpu, ok := strings.CutPrefix(scanner.Text(), "model name\t: "); ok {
				return strings.TrimSpace(cpu)
			}
		}

		if err := scanner.Err(); err != nil {
			return ""
		}
	}

	out, err := exec.Command("lscpu").Output()
	if err != nil {
		return ""
	}

	for _, line := range strings.Split(string(out), "\n") {
		if model, ok := strings.CutPrefix(line, "Model name:"); ok {
			return strings.TrimSpace(model)
		}
	}

	return ""
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
