package modules

import (
	"dfetch/internal/format"
	"os"
	"strings"
)

func Board(formatstring string) string {
	data, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		return format.Format(formatstring, format.Values{
			"board": "unknown",
		})
	}

	return format.Format(formatstring, format.Values{
		"board": strings.TrimSpace(string(data)),
	})
}
