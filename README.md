# Dfetch

A clean and practical system information tool with easy to understand output and fast startup times. It is designed to provide useful system information while being both lightweight enough to launch instantly with your terminal and very easy to configure.

<table>
  <tr>
    <td><img src="image/output_example_1.png" alt="Output example Linux mint" width="100%"></td>
    <td><img src="image/output_example_2.png" alt="output example Arch" width="100%"></td>
  </tr>
  <tr>
    <td><img src="image/output_example_3.png" alt="Output example Linux mint" width="100%"></td>
    <td><img src="image/output_example_4.png" alt="Output example Pop!_os" width="100%"></td>
  </tr>
</table>

## Why use this?

Dfetch does not try to compete with heavily customizable alternatives like [Neofetch](https://github.com/dylanaraps/neofetch) or [Fastfetch](https://github.com/fastfetch-cli/fastfetch). The project exists mainly as a fun project for myself, while still being useful for those who like: clean, easy to configure tools.


## Installation

Currently no official package for any platform is provided. You can either build Dfetch from source or [download the latest prebuilt binaries](https://github.com/David17c/Dfetch/releases).

## Customization

`~/.config/Dfetch/Dfetch.conf`

```
// Lines starting with `//` are comments, they are ignored by Dfetch.
// In the modules section you can change what info is displayed and in what order.

// 'emptyline' module can be used to get an empty line in between modules
modules {
	userinfo
	os
	Host
	kernel
	uptime
	packages
	shell
	de
	terminal
	cpu
	memory
	disk
	motherboard
	// battery
	// localip
	// time
	// date
}

ascii_size: default
// Only works with build in ascii.
// Ascii size can be either 'big', 'default' or 'small'. Default is big.

custom_ascii: default
// Set a custom ascii logo by providing a path to the txt file containing it.

accent_color: default
// Color used by the info labels

// Available colors:
// black, red, green, yellow, blue,
// magenta, cyan, white,
// bright_black, bright_red,
// bright_green, bright_yellow,
// bright_blue, bright_magenta,
// bright_cyan, bright_white
```

## Supported Operating systems

```txt
- Linux
  - Debian
  - Arch
  - CachyOS
  - Fedora
  - Linux Mint
  - OpenSUSE Leap
  - OpenSUSE Tumbleweed
  - Pop! OS
  - Ubuntu
  - Manjaro
```

While overtime support for more distro's will be added Dfetch does not try to support all distro's instead just focusing on the main ones people actually use.

## How to make a custom ascii art

first put the ASCII art you want to use into a `txt` file. it should look something like this.

```
             ...-:::::-...
          .-MMMMMMMMMMMMMMM-.
      .-MMMM`..-:::::::-..`MMMM-.
    .:MMMM.:MMMMMMMMMMMMMMM:.MMMM:.
   -MMM-M---MMMMMMMMMMMMMMMMMMM.MMM-
  :MMM:MM`  :MMMM:....::-...-MMMM:MMM:
 :MMM:MMM`  :MM:`  ``    ``  `:MMM:MMM:
.MMM.MMMM`  :MM.  -MM.  .MM-  `MMMM.MMM.
:MMM:MMMM`  :MM.  -MM-  .MM:  `MMMM-MMM:
:MMM:MMMM`  :MM.  -MM-  .MM:  `MMMM:MMM:
:MMM:MMMM`  :MM.  -MM-  .MM:  `MMMM-MMM:
.MMM.MMMM`  :MM:--:MM:--:MM:  `MMMM.MMM.
 :MMM:MMM-  `-MMMMMMMMMMMM-`  -MMM-MMM:
  :MMM:MMM:`                `:MMM:MMM:
   .MMM.MMMM:--------------:MMMM.MMM.
     '-MMMM.-MMMMMMMMMMMMMMM-.MMMM-'
       '.-MMMM``--:::::--``MMMM-.'
            '-MMMMMMMMMMMMM-'
               ``-:::::-``
```

Now you can give it color by adding color tags. For a list of supported colors look at the default config file.

```
             ${bright_white}...-:::::-...
${bright_white}          .-MMMMMMMMMMMMMMM-.
${bright_white}      .-MMMM${green}`..-:::::::-..`${bright_white}MMMM-.
${bright_white}    .:MMMM${green}.:MMMMMMMMMMMMMMM:.${bright_white}MMMM:.
${bright_white}   -MMM${green}-M---MMMMMMMMMMMMMMMMMMM.${bright_white}MMM-
${bright_white}  :MMM${green}:MM`  :MMMM:....::-...-MMMM:${bright_white}MMM:
${bright_white} :MMM${green}:MMM`  :MM:`  ``    ``  `:MMM:${bright_white}MMM:
${bright_white}.MMM${green}.MMMM`  :MM.  -MM.  .MM-  `MMMM.${bright_white}MMM.
${bright_white}:MMM${green}:MMMM`  :MM.  -MM-  .MM:  `MMMM-${bright_white}MMM:
${bright_white}:MMM${green}:MMMM`  :MM.  -MM-  .MM:  `MMMM:${bright_white}MMM:
${bright_white}:MMM${green}:MMMM`  :MM.  -MM-  .MM:  `MMMM-${bright_white}MMM:
${bright_white}.MMM${green}.MMMM`  :MM:--:MM:--:MM:  `MMMM.${bright_white}MMM.
${bright_white} :MMM${green}:MMM-  `-MMMMMMMMMMMM-`  -MMM-${bright_white}MMM:
${bright_white}  :MMM${green}:MMM:`                `:MMM:${bright_white}MMM:
${bright_white}   .MMM${green}.MMMM:--------------:MMMM.${bright_white}MMM.
${bright_white}     '-MMMM${green}.-MMMMMMMMMMMMMMM-.${bright_white}MMMM-'
${bright_white}       '.-MMMM${green}``--:::::--``${bright_white}MMMM-.'
${bright_white}            '-MMMMMMMMMMMMM-'
${bright_white}               ``-:::::-``
accentcolor: green
```

At the bottom of the file add an accentcolor: `accentcolor: green`. This is the color given to the info module labels.

Now in the config file add / edit `customascii: PATH_TO_FILE`. Dfetch should now be using your ASCII art.