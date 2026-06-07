# Dfetch

A clean and practical system information tool focused on clean, easy to understand output and fast startup times. It is designed to provide useful system information while being lightweight enough to launch instantly with your terminal.

<table>
  <tr>
    <td><img src="image/default_output_example_linuxmint.png" alt="default Linux mint output example" width="100%"></td>
    <td><img src="image/default_output_example_pop.png" alt="default Pop!_OS output example" width="100%"></td>
  </tr>
  <tr>
    <td><img src="image/default_output_example_arch.png" alt="default Arch output example" width="100%"></td>
    <td><img src="image/default_output_example_tux.png" alt="default Ubuntu output example" width="100%"></td>
  </tr>
</table>

## Why use this?

Dfetch does not try to compete with heavily customizable alternatives like [Neofetch](https://github.com/dylanaraps/neofetch) or [Fastfetch](https://github.com/fastfetch-cli/fastfetch). The project exists mainly as a fun project for myself, while still being useful for those who prefer clean, easy to configure tools with good defaults.


## Installation

Currently no official package for any platform is provided. You can either build Dfetch from source or [download the latest prebuilt binaries](https://github.com/David17c/Dfetch/releases).

## Customization

`~/.config/Dfetch/Dfetch.conf`

```
// Lines starting with `//` are comments and are ignored by Dfetch.
// In the modules section you can change what info is displayed and in what order.

// 'Emptyline' module can be used to get an empty line in between modules
modules {
	userinfo
	os
	kernel
	uptime
	shell
	de
	terminal
	cpu
	memory
	disk
	// battery
	localip
	// time
	// date
}

asciisize: default
// Ascii size can be either 'big', 'default' or 'small'. Default is big.

customascii: default
// Set a custom ascii logo by providing a path to the txt file containing it.

accentcolor: default
// Color used by the info labels

// Available colors:
// black, red, green, yellow, blue,
// magenta, cyan, white,
// bright_black, bright_red,
// bright_green, bright_yellow,
// bright_blue, bright_magenta,
// bright_cyan, bright_white
```

## Supported Linux distros

```txt
- Arch
- CachyOS
- Debian
- Fedora
- Linux Mint
- OpenSUSE Leap
- OpenSUSE Tumbleweed
- Pop! OS
- Ubuntu
```

## How to make a custom ascii art

first put the ASCII art you want to use into a `txt` file. it should look something like this.

```
MMMMMMMMMMMMMMMMMMMMMMMMMmds+.
MMm                       ymNMd+
MMd      /++                -sNMd:
MMNso\   dMM    `.::-. .-::.   hMN:
ddddMMh  dMM   :hNMNMNhNMNMNh:  NMm
    NMm  dMM   NMN/-+MMM+-/NMN  dMM
    NMm  dMM   MMm   MMM   dMM  dMM
    NMm  dMM   MMm   MMM   dMM  dMM
    NMm  dMM   mmd   mmm   yMM  dMM
    NMm  dMM.              ydm  dMM
    hMM  +MMd/-------...-:sdds  dMM
    -NMm  :hNMNNNmdddddddddy/   dMM
    -dMNs                      dMM
     \dMNmy+/:-------------:/yMMM
       \ydNMMMMMMMMMMMMMMMMMMMMM
```

Now you can give it color by adding color tags. For a list of supported colors look at the default config file.

```
${green}MMMMMMMMMMMMMMMMMMMMMMMMMmds+.
${green}MMm                       ymNMd+
${green}MMd      ${white}/++                ${green}-sNMd:
${green}MMNso\   ${white}dMM    `.::-. .-::.${green}   hMN:
${green}ddddMMh  ${white}dMM   :hNMNMNhNMNMNh:  ${green}NMm
${green}    NMm  ${white}dMM   NMN/-+MMM+-/NMN  ${green}dMM
${green}    NMm  ${white}dMM   MMm   MMM   dMM  ${green}dMM
${green}    NMm  ${white}dMM   MMm   MMM   dMM  ${green}dMM
${green}    NMm  ${white}dMM   mmd   mmm   yMM  ${green}dMM
${green}    NMm  ${white}dMM.              ydm  ${green}dMM
${green}    hMM  ${white}+MMd/-------...-:sdds  ${green}dMM
${green}    -NMm  ${white}:hNMNNNmdddddddddy/   ${green}dMM
${green}    -dMNs                      dMM
${green}     \dMNmy+/:-------------:/yMMM
${green}       \ydNMMMMMMMMMMMMMMMMMMMMM
```

At the bottom of the file add an accentcolor: `accentcolor: green`. This is the color given to the info module labels.


Now in the config file add / edit `customascii: PATH_TO_FILE`. Dfetch should now be using your ASCII art.
