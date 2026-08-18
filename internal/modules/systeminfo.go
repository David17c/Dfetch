package modules

import (
	"dfetch/internal/cmd"
	"dfetch/internal/config"
	"fmt"
)

type ModuleOutput struct {
	Name      string
	Label     string
	Color     string
	Separator string
	Value     string
	Mount     string
}

func CollectSystemInfo(modules []config.Module, distroName string, opts cmd.Options) []ModuleOutput {
	var output []ModuleOutput

	for _, module := range modules {
		var value string

		switch module.Name {
		case "userinfo":
			value = Userinfo(module.Format, module.Color, opts.NoColor)

		case "os":
			value = config.Format(module.Format, config.Values{
				"name": distroName,
			})

		case "kernel":
			value = Kernel(module.Format)

		case "cpu":
			value = Cpu(module.Format)

		case "memory":
			value = Memory(module.Format)

		case "swap":
			value = Swap(module.Format)

		case "local_ip":
			value = LocalIP(module.Format)

		case "uptime":
			value = Uptime(module.Format)

		case "battery":
			value = Battery(module.Format)

		case "bios":
			value = Bios(module.Format)

		case "de":
			value = DesktopEnvironment(module.Format)

		case "wm":
			value = WindowManager(module.Format)

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

		case "locale":
			value = Locale(module.Format)

		case "board":
			value = Board(module.Format)

		case "text":
			value = module.Format

		case "emptyline":
			output = append(output, ModuleOutput{
				Name: "emptyline",
			})
			continue

		case "color":
			value = Color(opts.NoColor)

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
