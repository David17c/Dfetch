package modules

import (
	"dfetch/internal/format"
	"os"
	"strings"
)

func Locale(formatstring string) string {
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if value := os.Getenv(key); value != "" {
			return format.Format(formatstring, format.Values{
				"locale": strings.TrimSpace(value),
			})
		}
	}
	return format.Format(formatstring, format.Values{
		"locale": "unknown",
	})
}
