package getsysinfo

import (
	"os"
	"path/filepath"
)

func Shell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}

	return filepath.Base(shell)
}
