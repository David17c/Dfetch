package modules

import (
	"dfetch/internal/config"
	"os"
	"strings"
)

func Locale(format string) string {
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if value := os.Getenv(key); value != "" {
			return config.Format(format, config.Values{
				"locale": strings.TrimSpace(value),
			})
		}
	}
	return config.Format(format, config.Values{
		"locale": "unknown",
	})
}
