package gpu

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//go:embed pci.ids
var pciIDs embed.FS

type database struct {
	Vendors map[string]string
	Devices map[string]map[string]string
}

var (
	db       *database
	once     sync.Once
	name     string
	cardRegexp = regexp.MustCompile(`^card\d+$`)
)

func init() {
	f, err := pciIDs.Open("pci.ids")
	if err != nil {
		return
	}
	defer f.Close()

	db, _ = loadDatabase(f)
}

func loadDatabase(r io.Reader) (*database, error) {
	db := &database{
		Vendors: make(map[string]string),
		Devices: make(map[string]map[string]string),
	}

	var vendor string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[0] == '#' {
			continue
		}

		if line[0] != '\t' {
			fields := strings.Fields(line)
			if len(fields) < 2 {
				continue
			}

			vendor = strings.ToLower(fields[0])
			db.Vendors[vendor] = strings.Join(fields[1:], " ")
			db.Devices[vendor] = make(map[string]string)
			continue
		}

		if strings.HasPrefix(line, "\t\t") {
			continue
		}

		fields := strings.Fields(strings.TrimLeft(line, "\t"))
		if len(fields) < 2 || vendor == "" {
			continue
		}

		device := strings.ToLower(fields[0])
		db.Devices[vendor][device] = strings.Join(fields[1:], " ")
	}

	return db, scanner.Err()
}

func readID(path string) (string, bool) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", false
	}

	id := strings.TrimSpace(string(b))
	id = strings.TrimPrefix(strings.ToLower(id), "0x")
	
	if len(id) > 0 && len(id) < 4 {
		id = fmt.Sprintf("%04s", id)
	}

	return id, true
}

func readBool(path string) bool {
	b, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(b)) == "1"
}

type gpuInfo struct {
	vendor string
	device string
}

func detectGPU() (gpuInfo, bool) {
	cards, err := filepath.Glob("/sys/class/drm/card*")
	if err != nil {
		return gpuInfo{}, false
	}

	var fallback *gpuInfo

	for _, card := range cards {
		if !cardRegexp.MatchString(filepath.Base(card)) {
			continue
		}

		vendor, ok1 := readID(filepath.Join(card, "device/vendor"))
		device, ok2 := readID(filepath.Join(card, "device/device"))
		if !ok1 || !ok2 {
			continue
		}

		info := gpuInfo{
			vendor: vendor,
			device: device,
		}

		if readBool(filepath.Join(card, "device/boot_vga")) {
			return info, true
		}

		if fallback == nil {
			fallback = &info
		}
	}

	if fallback != nil {
		return *fallback, true
	}

	return gpuInfo{}, false
}

func lookupName(info gpuInfo) string {
	if db == nil {
		return fmt.Sprintf("%s:%s", info.vendor, info.device)
	}

	if devices, ok := db.Devices[info.vendor]; ok {
		if device, ok := devices[info.device]; ok {
			return device
		}
	}

	if vendor, ok := db.Vendors[info.vendor]; ok {
		return fmt.Sprintf("%s (%s)", vendor, info.device)
	}

	return fmt.Sprintf("%s:%s", info.vendor, info.device)
}

func Name() string {
	once.Do(func() {
		info, ok := detectGPU()
		if !ok {
			name = "Unknown"
			return
		}

		name = lookupName(info)
	})

	return name
}