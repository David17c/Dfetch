# Dfetch

Dfetch is a lightweight system information tool focused on clean output, fast startup times, and simple configuration. It provides useful system details without the complexity of heavily customizable alternatives.

<table>
  <tr>
    <td><img src="image/output_example_1.png" alt="Output example Pop os" width="100%"></td>
    <td><img src="image/output_example_2.png" alt="output example Linux MInt" width="100%"></td>
  </tr>
  <tr>
    <td><img src="image/output_example_3.png" alt="Output example Debian" width="100%"></td>
    <td><img src="image/output_example_4.png" alt="Output example Fedora" width="100%"></td>
  </tr>
</table>

## Why use this?

Dfetch is designed for those who want a simple system information tool with sensible defaults, clean output, and fast startup times. Rather than prioritizing extensive customization, Dfetch focuses on being lightweight, easy to configure, and pleasant to use.

## Features

```
- Fast startup time
- Simple configuration file
- Useful, clutter-free system information
- Custom ASCII art support
- Configurable modules
- No external dependencies
- Clean default look
```

## Installation

To install Dfetch, visit [the releases page](https://github.com/David17c/Dfetch/releases) and either download the package for your operating system, download a prebuilt binary, or build Dfetch from source.

## Customization

`~/.config/Dfetch/Dfetch.conf`

```
// Lines starting with `//` are comments and are ignored by Dfetch.
// In the modules section, you can change which information is displayed and in what order.

// Insert empty lines in the modules block to get empty lines in the final output.
modules {
    userinfo
    os
    host
    kernel
    uptime
    shell
    terminal
    de
    packages
    cpu
    ram
    swap
    disk
    motherboard
    local_ip
    // battery
    // time
    // date
}

custom_ascii: default
// Set a custom ASCII logo by providing the path to a text file containing it.

accent_color: default
// Color used for the information labels.

// Available colors:
// black, red, green, yellow, blue,
// magenta, cyan, white,
// bright_black, bright_red,
// bright_green, bright_yellow,
// bright_blue, bright_magenta,
// bright_cyan, bright_white
```

## Supported Linux distros

```
- Arch
- Bazzite
- CachyOS
- Debian
- Fedora
- Linux Mint
- Manjaro
- OpenSUSE Leap
- OpenSUSE Tumbleweed
- Pop!_OS
- Ubuntu
- Zorin OS
- EndeavourOS
- Nix OS
```

If your favorite distro isn't listed, it may still be supported. This list only includes distributions that Dfetch has built-in ASCII art for and that have been tested.

## How to make a custom ascii art

Save your ASCII art in a text file. It should look something like this.

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

You can then optionally add colors by using color tags. For a list of supported colors look at the default config file.

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
accent_color: green
```

At the bottom of the file add an accent color: `accent_color: green`. This is the color given to the info module labels.

In your config file, set: `custom_ascii: PATH_TO_FILE`. Dfetch should now be using your ASCII art.