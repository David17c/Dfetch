package modules

import (
	"dfetch/internal/config"
	"os"
	"strings"
)

func Board(format string) string {
	data, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		return config.Format(format, config.Values{
			"board": "unknown",
		})
	}

	return config.Format(format, config.Values{
		"board": strings.TrimSpace(string(data)),
	})
}
