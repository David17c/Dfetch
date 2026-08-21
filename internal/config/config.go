package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"dfetch/internal/cmd"
)

type Config struct {
	Logo    ASCIIConfig `json:"ascii"`
	Modules []Module    `json:"modules"`
}

type ASCIIConfig struct {
	Enabled       bool   `json:"enabled"`
	Path          string `json:"path"`
	PaddingTop    int    `json:"padding_top"`
	PaddingBottom int    `json:"padding_bottom"`
}

type Module struct {
	Name      string `json:"name"`
	Label     string `json:"label,omitempty"`
	Color     string `json:"color,omitempty"`
	Format    string `json:"format,omitempty"`
	Separator string `json:"separator,omitempty"`

	// disk module only
	Mount string `json:"mount,omitempty"`
}

// return path to default config
func configPath(opts cmd.Options) (string, error) {
	var configDir string

	if opts.SetConfig != "" {
		_, err := os.Stat(opts.SetConfig)
		if err != nil {
			return "", err
		}
		configDir = opts.SetConfig
		return configDir, nil
	}
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("unable to find config directory: %w", err)
	}

	return filepath.Join(configDir, "dfetch", "dfetch.json"), nil
}

// Read the config file put everyting in the struct and validate the config
func ReadConfig(opts cmd.Options) (Config, error) {
	path, err := configPath(opts)
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("unable to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("invalid config JSON: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// validate modules only use allowed options
func (c Config) Validate() error {

	for i, module := range c.Modules {
		if module.Name == "" {
			return fmt.Errorf("module %d is missing a name", i+1)
		}

		if module.Mount != "" && module.Name != "disk" {
			return fmt.Errorf("module %s does cannot have a mount", module.Name)
		}

		// Empty line module
		if module.Name == "emptyline" {
			if module.HasOptions() {
				return fmt.Errorf("%s module does not support any options", module.Name)
			}
			continue
		}

		if module.Name == "color" {
			if module.HasOptions() {
				return fmt.Errorf("%s module does not support any options", module.Name)
			}
			continue
		}

		// Text module
		if module.Name == "text" {
			if module.Format == "" {
				return fmt.Errorf("text module is missing text")
			}

			if module.Label != "" || module.Separator != "" {
				return fmt.Errorf("Text module only supports format option")
			}
			continue
		}

		// Userinfo module
		if module.Name == "userinfo" {
			if module.Label != "" {
				return fmt.Errorf("module 'userinfo' cannot have a label")
			}

			if module.Separator != "" {
				return fmt.Errorf("module 'userinfo' cannot have a separator")
			}
		}

		if module.Name == "disk" && module.Mount != "" {
			info, err := os.Stat(module.Mount)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf(
						"disk mount '%s' does not exist",
						module.Mount,
					)
				}

				return fmt.Errorf(
					"unable to access disk mount '%s': %w",
					module.Mount,
					err,
				)
			}

			if !info.IsDir() {
				return fmt.Errorf(
					"disk mount '%s' is not a directory",
					module.Mount,
				)
			}
		}
	}

	return nil
}

// create default config file
func CreateConfigFile(opts cmd.Options, distroID string, distroIDLike string) error {
	path, err := configPath(opts)
	if err != nil {
		return err
	}

	configDir := filepath.Dir(path)

	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("unable to create config directory: %w", err)
	}

	if _, err := os.Stat(path); err == nil {
		if opts.ResetConfig {
			RemoveConfigFile(opts)
		} else {
			return nil
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("unable to check config file: %w", err)
	}

	// Decide what color to use for default config based on the distro
	colorToUse := mapLogoToColor(distroID)

	if colorToUse == "" {
		colorToUse = mapLogoToColor(distroIDLike)
		if colorToUse == "" {
			colorToUse = "bold_cyan"
		}
	}

	//The default config
	defaultConfig := Config{
		Logo: ASCIIConfig{
			Enabled:       true,
			Path:          "builtin",
			PaddingTop:    1,
			PaddingBottom: 1,
		},

		Modules: []Module{
			{Name: "emptyline"},

			{Name: "os", Label: "OS", Color: colorToUse, Separator: ": ", Format: "{name}"},
			{Name: "host", Label: "Host", Color: colorToUse, Separator: ": ", Format: "{host}"},
			{Name: "kernel", Label: "Kernel", Color: colorToUse, Separator: ": ", Format: "{version}"},
			{Name: "bios", Label: "BIOS", Color: colorToUse, Separator: ": ", Format: "{bios}"},

			{Name: "emptyline"},

			{Name: "cpu", Label: "CPU", Color: colorToUse, Separator: ": ", Format: "{short}"},
			{Name: "memory", Label: "RAM", Color: colorToUse, Separator: ": ", Format: "{used} / {total} {unit} ({percent}%)"},
			{Name: "disk", Label: "Disk", Color: colorToUse, Separator: ": ", Format: "{used} / {total} {unit} ({percent}%)", Mount: "/"},
			{Name: "board", Label: "Board", Color: colorToUse, Separator: ": ", Format: "{board}"},

			{Name: "emptyline"},

			{Name: "packages", Label: "Pkgs", Color: colorToUse, Separator: ": ", Format: "{packages}"},
			{Name: "shell", Label: "Shell", Color: colorToUse, Separator: ": ", Format: "{name} {version}"},
			{Name: "terminal", Label: "Term", Color: colorToUse, Separator: ": ", Format: "{name} {version}"},
			{Name: "de", Label: "DE", Color: colorToUse, Separator: ": ", Format: "{de}"},
			{Name: "wm", Label: "WM", Color: colorToUse, Separator: ": ", Format: "{name} ({sessiontype})"},
			{Name: "uptime", Label: "Uptime", Color: colorToUse, Separator: ": ", Format: "{uptime}"},
			{Name: "local_ip", Label: "Local IP", Color: colorToUse, Separator: ": ", Format: "{address}"},
			{Name: "locale", Label: "Lang", Color: colorToUse, Separator: ": ", Format: "{locale}"},

			{Name: "emptyline"},

			{Name: "color"},

			{Name: "emptyline"},
		},
	}

	data, err := json.MarshalIndent(defaultConfig, "", "    ")
	if err != nil {
		return fmt.Errorf("unable to encode default config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("unable to write config file: %w", err)
	}

	fmt.Printf("succesfully created config file '%s'.\n", path)

	return nil
}

// Remove existing config
func RemoveConfigFile(opts cmd.Options) {
	path, err := configPath(opts)
	if err != nil {
		return
	}

	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return
		}

		fmt.Printf("unable to remove config file: %s\n", err)
		return
	}
}

// helper function used to validate modules that don't support any options
func (m Module) HasOptions() bool {
	return m.Label != "" ||
		m.Color != "" ||
		m.Format != "" ||
		m.Separator != "" ||
		m.Mount != ""
}

// map of accentcolor -> distro
func mapLogoToColor(ID string) string {
	colors := map[string]string{
		"almalinux":           "bold_yellow",
		"alpine":              "bold_magenta",
		"arch":                "bold_cyan",
		"artix":               "bold_cyan",
		"bazzite":             "bold_bright_blue",
		"cachyos":             "bold_cyan",
		"centos":              "bold_green",
		"clear-linux-os":      "bold_blue",
		"debian":              "bold_red",
		"deepin":              "bold_green",
		"elementary":          "bold_blue",
		"endeavouros":         "bold_magenta",
		"fedora":              "bold_blue",
		"garuda":              "bold_red",
		"gentoo":              "bold_magenta",
		"kali":                "bold_blue",
		"linuxmint":           "bold_green",
		"manjaro":             "bold_green",
		"mx":                  "bold_white",
		"nobara":              "bold_white",
		"nixos":               "bold_magenta",
		"opensuse-leap":       "bold_green",
		"opensuse-tumbleweed": "bold_green",
		"opensuse-slowroll":   "bold_green",
		"parrot":              "bold_cyan",
		"pclinuxos":           "bold_white",
		"peppermint":          "bold_white",
		"popos":               "bold_cyan",
		"qubes":               "bold_magenta",
		"rhel":                "bold_red",
		"rocky":               "bold_green",
		"slackware":           "bold_bright_blue",
		"tuxedo":              "bold_red",
		"ubuntu":              "bold_yellow",
		"vanilla":             "bold_yellow",
		"void":                "bold_white",
		"zorin":               "bold_blue",
	}

	return colors[ID]
}
