# Dfetch

Dfetch is a command-line utility inspired by [Neofetch](https://github.com/dylanaraps/neofetch?utm_source=chatgpt.com) and written in Go. It’s designed to be simple and easy to configure without overwhelming the user with endless options. Dfetch focuses on a clean default, a minimal and compact look, and extremely fast performance. It also keeps the displayed information focused and avoids external dependencies.


## Installation

You can install Dfetch from the releases page:

[Dfetch Releases](https://github.com/David17c/Dfetch/releases)

Download the version that matches your Linux distribution.

If no compatible release is available for your distro, using the steps below you can compile Dfetch from source instead.

### Step 1

Go to the main GitHub repository page:

[Dfetch GitHub Repository](https://github.com/David17c/Dfetch)

Click the green Code button, then select Download ZIP.

### Step 2

Unzip the file you just downloaded and navigate to it in the terminal using the `cd` command.

```bash
cd ~/Downloads/Dfetch
```

### Step 3

Install the [Go](https://go.dev/) programming language. This step is different depending on your distro.

Debian / Ubuntu:

```bash
sudo apt install golang-go
```

Arch:

```bash
sudo pacman -S go
```

Fedora:

```bash
sudo dnf install golang
```

You can verify it's installed by running:

```bash
go version
```

### Step 4

Now that you have Go installed and have navigated to the root of the project folder, you're nearly done.

Run:

```bash
go build -ldflags="-s -w" -trimpath -o build/Dfetch
```

To compile the program and store the executable file in the `Dfetch/build` folder.

You can now execute the file you just created and use the program.

## Usage

Run Dfetch using:

```bash
dfetch
```

Or, if you compiled it from source:

```bash
./build/Dfetch
```

## Customization

When you first run Dfetch it creates a config file in `~/.config/Dfetch/Dfetch.conf`.

Default config file:

```
// Lines starting with `//` are comments and are ignored by Dfetch.
// In the System Information section you can change what info is displayed and in what order.

// ------------------------
// Color
//------------------------

// ASCII color
labelcolor: default
infocolor: default
headercolor: default
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
//shell
//de
battery
```

## Supported information:

```txt
- Hostname
- Username
- Distro / OS
- CPU
- Memory
- Battery
- DE (desktop enviroment)
- Kernel
- Local IP
- Shell
- Uptime
```

## Supported distros

The following is a list of supported distros. Dfetch definitely works on many more distros, however these distros have been tested and have their own ASCII art. More distros will be added in the future.

```txt
- Debian
- Arch
- Fedora
- Ubuntu
- Linux Mint
```

## Supported colors:
```txt

- Black
- Red
- Green
- Yellow
- Blue
- Magenta
- Cyan
- White

- Bright_black / gray / grey
- Bright_red
- Bright_green
- Bright_yellow
- Bright_blue
- Bright_magenta
- Bright_cyan
- Bright_white
```