# Installation

This guide covers installation instructions for the base operating systems that most supported Linux distributions are derived from. If you're using a derivative distribution, you can usually follow the instructions for its corresponding base distribution.

* [Debian-based distributions](#debian)
* [Fedora-based distributions](#fedora)
* [Arch-based distributions](#arch)
* [NixOS](#nixos)
* [Build from source](#build-from-source)

## Debian

Dfetch provides prebuilt packages for Debian and Debian-based distributions.

Download the `.deb` package from the [Releases page](https://github.com/David17c/Dfetch/releases) and install it. Alternatively, you can download the prebuilt binary from the Releases page, which works on Debian-based distributions as well.

## Fedora

Dfetch provides prebuilt packages for Fedora and Fedora-based distributions.

Download the `.rpm` package from the [Releases page](https://github.com/David17c/Dfetch/releases) and install it. Alternatively, you can download the prebuilt binary from the Releases page, which works on Fedora-based distributions as well.

## Arch-based distributions

Dfetch does not currently provide a prebuilt package for Arch or Arch-based distributions. Instead, you can either:

* Download the prebuilt binary from the [Releases page](https://github.com/David17c/Dfetch/releases), or
* Build Dfetch from source by following the [build instructions](#build-from-source).

## NixOS

### Use the Flake

Add Dfetch to your flake inputs:

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    dfetch.url = "github:David17c/Dfetch";
  };
}
```

Install the package in your NixOS configuration:

```nix
{ inputs, pkgs, ... }:

{
  environment.systemPackages = [
    inputs.dfetch.packages.${pkgs.stdenv.hostPlatform.system}.default
  ];
}
```

Then rebuild:

```bash
sudo nixos-rebuild switch --flake /path/to/your/flake#your-hostname
```

### Use the NixOS Module

The flake also exposes a small NixOS module that installs Dfetch through `programs.dfetch`.

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    dfetch.url = "github:David17c/Dfetch";
  };

  outputs = { self, nixpkgs, dfetch, ... }: {
    nixosConfigurations.your-hostname = nixpkgs.lib.nixosSystem {
      system = "x86_64-linux";
      modules = [
        dfetch.nixosModules.default
        {
          programs.dfetch.enable = true;
        }
      ];
    };
  };
}
```

### Run Without Installing

You can run Dfetch directly from the flake:

```bash
nix run github:David17c/Dfetch
```

### Development Shell

Enter a shell with Go and Git available:

```bash
nix develop github:David17c/Dfetch
```

### Package Counting on Nix

When the `packages` module is enabled in Dfetch's config, Nix package counting checks the standard system and user profile paths and combines their results. Dfetch queries each profile's Nix requisites and filters them with the same package-oriented rules used by Fastfetch, so it is not limited to executable links in `bin`.

- `/run/current-system` for the active NixOS system profile
- `~/.nix-profile` for the user's default Nix profile
- `$XDG_STATE_HOME/nix/profile` or `~/.local/state/nix/profile` for the newer user profile location
- `/etc/profiles/per-user/$USER` for the per-user profile

Missing profile directories are ignored, so Dfetch still works on systems that only have some of these paths.

### Dfetch Configuration

Dfetch reads its runtime configuration from:

```text
~/.config/dfetch/dfetch.conf
```

Include `packages` in the modules block to show package information:

```text
modules {
    userinfo
    os
    kernel
    packages
    memory
    disk
}
```

> [!NOTE]
> Credit to @crispdark for writing the NixOS instructions.

## Build from source

Dfetch is written in Go and has no external build dependencies beyond the Go toolchain.

### Requirements

- Go 1.26.4
- Git

### Build

Clone the repository and build the binary:

```bash
git clone https://github.com/David17c/Dfetch.git
cd Dfetch
go build -o dfetch .
```

Alternatively, install it directly into your Go binary directory:

```bash
go install github.com/David17c/Dfetch@latest
```

If you built the binary with `go build`, you can optionally install it system-wide:

```bash
sudo install -Dm755 dfetch /usr/local/bin/dfetch
```