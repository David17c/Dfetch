package modules

import (
	"dfetch/internal/format"
	"os"
	"os/user"
)

func Userinfo(formatstring string, color string, noColor bool) string {
	fields := format.Fields(formatstring)
	values := format.Values{}

	var username, hostname string

	for _, field := range fields {
		switch field {
		case "username":
			if username == "" {
				username = Username()
			}
			values["username"] = username

		case "hostname":
			if hostname == "" {
				hostname = Hostname()
			}
			values["hostname"] = hostname
		}
	}

	if color != "" {
		c := format.GetColorCode(color, noColor)
		r := format.GetColorCode("reset", noColor)

		if _, ok := values["username"]; ok {
			values["username"] = c + values["username"] + r
		}

		if _, ok := values["hostname"]; ok {
			values["hostname"] = c + values["hostname"] + r
		}
	}

	return format.Format(formatstring, values)
}

func Hostname() string {
	if hostname, err := os.Hostname(); err == nil && hostname != "" {
		return hostname
	}

	for _, key := range []string{
		"HOST",
		"HOSTNAME",
	} {
		if hostname := os.Getenv(key); hostname != "" {
			return hostname
		}
	}

	return "unknown"
}

func Username() string {
	for _, key := range []string{
		"USER",
		"LOGNAME",
		"USERNAME",
	} {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	if u, err := user.Current(); err == nil && u.Username != "" {
		return u.Username
	}

	return "unknown"
}
