package modules

import (
	"fmt"
	"os"
	"strings"
)

func MotherBoard() string {
	name, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		return "unknown"
	}

	name = []byte(strings.TrimSpace(string(name)))

	return fmt.Sprintf("%s", name)
}
