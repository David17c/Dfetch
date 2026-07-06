# Dfetch on NixOS

This document provides detailed instructions for installing and configuring Dfetch on NixOS.

## Installation Methods

### Method 1: Using Flakes in Your System Configuration

Add Dfetch to your `flake.nix`:

```nix
inputs = {
  dfetch.url = "github:crispdark/Dfetch";
};
```

Then add to your `environment.systemPackages`:

```nix
environment.systemPackages = with pkgs; [
  dfetch.packages.${pkgs.stdenv.hostPlatform.system}.default
];
```

Rebuild your system:

```bash
sudo nixos-rebuild switch --flake /path/to/your/flake#yourHostname
```

### Method 2: Direct Flake Execution

Run Dfetch directly from the flake without installing it:

```bash
nix run github:crispdark/Dfetch
```

### Method 3: Using the NixOS Module

For more advanced configuration and automatic setup, use the NixOS module.

Add to your `flake.nix` inputs:

```nix
inputs = {
  dfetch.url = "github:crispdark/Dfetch";
};
```

Add to your NixOS configuration imports:

```nix
imports = [
  inputs.dfetch.nixosModules.default
];
```

Enable and configure Dfetch:

```nix
services.dfetch = {
  enable = true;
  
  # Customize which modules to display
  modules = [
    "userinfo"
    "os"
    "kernel"
    "uptime"
    "packages"
    "memory"
    "disk"
  ];
  
  # Set custom colors
  labelColor = "green";
  userinfoColor = "bright_green";
  infoColor = "default";
  
  # Optional: use custom ASCII art
  customAscii = null;
};
```

Rebuild your system:

```bash
sudo nixos-rebuild switch --flake /path/to/your/flake#yourHostname
```

## NixOS Module Options

### Available Modules

- `userinfo` - User and hostname information
- `os` - Operating system name
- `host` - Hostname
- `kernel` - Kernel version
- `uptime` - System uptime
- `shell` - Current shell
- `terminal` - Terminal application
- `desktop` - Desktop environment
- `packages` - Package manager info (now with NixOS support!)
- `cpu` - CPU information
- `memory` - Memory usage
- `swap` - Swap memory usage
- `disk` - Disk usage
- `motherboard` - Motherboard information
- `local_ip` - Local IP address
- `battery` - Battery status (if available)
- `time` - Current time
- `date` - Current date

### Available Colors

```
black, red, green, yellow, blue,
magenta, cyan, white,
bright_black, bright_red,
bright_green, bright_yellow,
bright_blue, bright_magenta,
bright_cyan, bright_white
```

## Packages Support

Dfetch now has full support for NixOS package detection! The `packages` module will automatically detect and count packages from your Nix profile using `nix profile list`.

Add `packages` to your modules list to display package count:

```nix
services.dfetch = {
  enable = true;
  
  modules = [
    "userinfo"
    "os"
    "packages"  # Shows Nix package count
    "memory"
  ];
};
```

## Configuration File

The module automatically generates the configuration file at `~/.config/Dfetch/Dfetch.conf`. You can also manually edit this file:

```
modules {
    userinfo
    os
    kernel
    packages
    memory
    disk
}

label_color: green
userinfo_color: bright_green
info_color: default
custom_ascii: default
```

## Custom ASCII Art

To use custom ASCII art on NixOS:

1. Create your ASCII art file:

```bash
mkdir -p ~/.config/Dfetch
vim ~/.config/Dfetch/custom_logo.txt
```

2. Update your configuration to use it:

```nix
services.dfetch = {
  enable = true;
  customAscii = "/home/youruser/.config/Dfetch/custom_logo.txt";
};
```

Or in the config file:

```
custom_ascii: /home/youruser/.config/Dfetch/custom_logo.txt
```

## Troubleshooting

### Dfetch not found after rebuild

Make sure you've added it to `environment.systemPackages` or enabled the service module. Then rebuild:

```bash
sudo nixos-rebuild switch --flake /path/to/your/flake#yourHostname
```

### Packages showing "unknown"

Ensure `nix` command is available in your PATH. This is usually the case on NixOS by default.

### Module configuration not applied

If you modify the module options, remember to rebuild:

```bash
sudo nixos-rebuild switch --flake /path/to/your/flake#yourHostname
```

### Font/Color issues in terminal

Make sure your terminal supports ANSI colors. Most modern terminals do. If colors look wrong, try a different terminal or adjust the color settings in your Dfetch configuration.

## More Information

For general Dfetch usage and configuration, see the [main README](README.md).
