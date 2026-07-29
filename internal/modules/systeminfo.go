package modules

import (
	"dfetch/internal/config"
	"fmt"
)

type ModuleOutput struct {
	Name      string
	Label     string
	Color     string
	Separator string
	Value     string

	// Disk only
	Mount string
}

func CollectSystemInfo(modules []config.Module, distroName string) []ModuleOutput {
	var output []ModuleOutput

	for _, module := range modules {
		var value string

		switch module.Name {
		case "userinfo":
			username, hostname := Userinfo()

			if module.Color != "" {
				c := config.GetColorCode(module.Color)
				r := "\x1b[0m"

				value = c + username + r + "@" + c + hostname + r
			} else {
				value = username + "@" + hostname
			}

		case "os":
			value = distroName

		case "kernel":
			value = Kernel(module.Format)

		case "cpu":
			value = Cpu(module.Format)

		case "memory":
			value = Memory(module.Format)

		case "swap":
			value = Swap(module.Format)

		case "local_ip":
			value = Local_IP()

		case "uptime":
			value = Uptime(module.Format)

		case "battery":
			value = Battery()

		case "bios":
			value = Bios(module.Format)

		case "desktop":
			value = DesktopEnvironment(module.Format)

		case "shell":
			value = Shell(module.Format)

		case "terminal":
			value = Terminal(module.Format)

		case "disk":
			value = Disk(module.Format, module.Mount)

		case "datetime":
			value = DateTime(module.Format)

		case "packages":
			value = Packages(module.Format)

		case "host":
			value = Host(module.Format)

		case "board":
			value = Board()

		case "text":
			value = module.Text

		case "emptyline":
			output = append(output, ModuleOutput{
				Name: "emptyline",
			})
			continue

		default:
			fmt.Printf("Unknown module '%s'\n", module.Name)
			continue
		}

		output = append(output, ModuleOutput{
			Name:      module.Name,
			Label:     module.Label,
			Color:     module.Color,
			Separator: module.Separator,
			Value:     value,
			Mount:     module.Mount,
		})
	}

	return output
}
