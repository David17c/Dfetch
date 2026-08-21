// checks if a process exists

package helpers

import (
	"os"
	"strconv"
	"strings"
)

func ProcessName(pid int) string {
	return ProcessNameFromPID(strconv.Itoa(pid))
}

func ProcessNameFromPID(pid string) string {
	data, err := os.ReadFile("/proc/" + pid + "/comm")
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func ParentPID(pid int) int {
	data, err := os.ReadFile("/proc/" + strconv.Itoa(pid) + "/stat")
	if err != nil {
		return 0
	}

	line := string(data)
	end := strings.LastIndexByte(line, ')')
	if end < 0 || end+2 >= len(line) {
		return 0
	}

	fields := strings.Fields(line[end+2:])
	if len(fields) < 2 {
		return 0
	}

	ppid, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0
	}

	return ppid
}

func ProcessInParentChain(names ...string) string {
	wanted := make(map[string]struct{}, len(names))
	for _, name := range names {
		wanted[name] = struct{}{}
	}

	pid := os.Getpid()
	seen := make(map[int]struct{})

	for pid > 1 {
		if _, ok := seen[pid]; ok {
			break
		}
		seen[pid] = struct{}{}

		name := ProcessName(pid)
		if _, ok := wanted[name]; ok {
			return name
		}

		next := ParentPID(pid)
		if next <= 0 || next == pid {
			break
		}

		pid = next
	}

	return ""
}

func ProcessExists(names ...string) string {
	wanted := make(map[string]struct{}, len(names))
	for _, name := range names {
		wanted[name] = struct{}{}
	}

	entries, err := os.ReadDir("/proc")
	if err != nil {
		return ""
	}

	for _, entry := range entries {
		pid := entry.Name()
		if !isPID(pid) {
			continue
		}

		name := ProcessNameFromPID(pid)
		if _, ok := wanted[name]; ok {
			return name
		}
	}

	return ""
}

func isPID(name string) bool {
	if name == "" {
		return false
	}

	for i := range name {
		if name[i] < '0' || name[i] > '9' {
			return false
		}
	}

	return true
}
