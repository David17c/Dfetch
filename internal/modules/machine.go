package modules

import (
	"fmt"
	"os"
	"strings"
)

func Host() string {

	productName, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_family")
	if err != nil {
		return "unknown"
	}

	productName = []byte(strings.TrimSpace(string(productName)))

	return fmt.Sprintf("%s", productName)
}

// can also be gotten from hostnamectl command
