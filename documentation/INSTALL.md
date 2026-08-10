# Installation

* [Debian-based distributions](#debian)
* [Fedora-based distributions](#fedora)
* [Arch-based distributions](#arch)
* [Nix](#nix)
* [Build from source](#build-from-source)

## Debian

Dfetch provides prebuilt packages for Debian and Debian-based distributions.

Download the `.deb` package from the [Releases page](https://github.com/David17c/Dfetch/releases) and install it. Alternatively, you can download the prebuilt binary from the Releases page, which works on Debian-based distributions as well.

## Fedora

Dfetch provides prebuilt packages for Fedora and Fedora-based distributions.

Download the `.rpm` package from the [Releases page](https://github.com/David17c/Dfetch/releases) and install it. Alternatively, you can download the prebuilt binary from the Releases page, which works on Fedora-based distributions as well.

## Arch

Dfetch does not currently provide a prebuilt package for Arch or Arch-based distributions. Instead, you can either:

* Download the prebuilt binary from the [Releases page](https://github.com/David17c/Dfetch/releases), or
* Build Dfetch from source by following the [build instructions](#build-from-source).

## Nix

### Install with Flakes

Add Dfetch as a flake input:

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    dfetch.url = "github:David17c/Dfetch";
  };
}
```

Then add the package to your system configuration:

```nix
{ inputs, pkgs, ... }:

{
  environment.systemPackages = [
    inputs.dfetch.packages.${pkgs.stdenv.hostPlatform.system}.default
  ];
}
```

Rebuild your system:

```bash
sudo nixos-rebuild switch --flake /path/to/your/flake#your-hostname
```

### NixOS Module

Alternatively, the flake provides a NixOS module:

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

### Other Nix systems

You can run Dfetch directly without installing it:

```bash
nix run github:David17c/Dfetch
```

Or enter a development shell with the required build tools:

```bash
nix develop github:David17c/Dfetch
```

## Build from source

Dfetch is written in Go and has no external build dependencies beyond the Go toolchain.

### Requirements

- Go 1.24.4+
- Git

### Build

Clone the repository and build the binary:

```bash
git clone https://github.com/David17c/Dfetch.git
cd Dfetch
go build .
```