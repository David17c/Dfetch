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
	Text      string `json:"text,omitempty"`
	Color     string `json:"color,omitempty"`
	Format    string `json:"format,omitempty"`
	Separator string `json:"separator,omitempty"`

	// Disk module only
	Mount string `json:"mount,omitempty"`
}

func configPath(opts cmd.Options) (string, error) {
	var configDir string

	if opts.SetConfig != "" {
		_, err := os.Stat(opts.SetConfig)
		if err == nil {
			configDir = opts.SetConfig
			return configDir, nil
		}
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
			if module.Text == "" {
				return fmt.Errorf("text module is missing text")
			}

			if module.Label != "" || module.Separator != "" || module.Format != "" {
				return fmt.Errorf("Text module only supports text option")
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
		} else {
			if module.Label == "" {
				return fmt.Errorf(
					"module '%s' is missing a label",
					module.Name,
				)
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

func CreateConfigFile(opts cmd.Options) error {
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

	defaultConfig := Config{
		Logo: ASCIIConfig{
			Enabled:       true,
			Path:          "builtin",
			PaddingTop:    1,
			PaddingBottom: 1,
		},

		Modules: []Module{
			{Name: "emptyline"},

			{Name: "os", Label: "OS", Color: "bold_cyan", Separator: ":"},
			{Name: "host", Label: "Host", Color: "bold_cyan", Separator: ":"},
			{Name: "kernel", Label: "Kernel", Color: "bold_cyan", Format: "short", Separator: ":"},
			{Name: "bios", Label: "BIOS", Color: "bold_cyan", Format: "short", Separator: ":"},

			{Name: "emptyline"},

			{Name: "cpu", Label: "CPU", Color: "bold_cyan", Format: "short", Separator: ":"},
			{Name: "memory", Label: "RAM", Color: "bold_cyan", Separator: ":", Format: "long"},
			{Name: "disk", Label: "Disk", Color: "bold_cyan", Separator: ":", Mount: "/", Format: "long"},
			{Name: "board", Label: "Board", Color: "bold_cyan", Separator: ":"},

			{Name: "emptyline"},

			{Name: "packages", Label: "Pkgs", Color: "bold_cyan", Separator: ":"},
			{Name: "shell", Label: "Shell", Color: "bold_cyan", Separator: ":"},
			{Name: "terminal", Label: "Term", Color: "bold_cyan", Separator: ":"},
			{Name: "de", Label: "DE", Color: "bold_cyan", Separator: ":"},
			{Name: "wm", Label: "WM", Color: "bold_cyan", Separator: ":"},
			{Name: "uptime", Label: "Uptime", Color: "bold_cyan", Separator: ":"},
			{Name: "local_ip", Label: "Local IP", Color: "bold_cyan", Separator: ":"},
			{Name: "locale", Label: "Lang", Color: "bold_cyan", Separator: ":"},

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

// Removes existing config
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

func (m Module) HasOptions() bool {
	return m.Label != "" ||
		m.Text != "" ||
		m.Color != "" ||
		m.Format != "" ||
		m.Separator != "" ||
		m.Mount != ""
}
