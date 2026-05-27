# Dfetch

A minimal and practical system information tool focused on clean, compact output and fast startup times. It is designed to provide useful system information while being lightweight enough to launch instantly with your terminal.

## Why use this?

Dfetch does not try to compete with heavily customizable alternatives like [Neofetch](https://github.com/dylanaraps/neofetch) or [Fasfetch](https://github.com/fastfetch-cli/fastfetch). The project exists mainly as a fun and learning project for myself, while still being useful for people who prefer minimal, practical tools with good defaults.

## Installation

This program does not currently provide official packages for any platform. You can either build it from source or [download the latest prebuilt binaries](https://github.com/David17c/Dfetch/releases).

## Customization

When you first run Dfetch it creates a config file in `~/.config/Dfetch/Dfetch.conf`.

```
// Lines starting with `//` are comments and are ignored by Dfetch.
// In the System Information section you can change what info is displayed and in what order.

// ------------------------
// Color
//------------------------

// ASCII color
headercolor: default
labelcolor: default
infocolor: default
asciicolor: default

// Available colors:
// black, red, green, yellow, blue,
// magenta, cyan, white,
// bright_black, bright_red,
// bright_green, bright_yellow,
// bright_blue, bright_magenta,
// bright_cyan, bright_white

// ------------------------
// System Information
// ------------------------

os
kernel
uptime
cpu
memory
localip
// shell
// de
// battery
```

## Supported distros

```txt
- Debian
- Arch
- Fedora
- Ubuntu
- Linux Mint
```
