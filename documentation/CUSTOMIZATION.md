# Customization

The configuration file is stored in `~/.config/Dfetch`. Every time Dfetch is started, it checks if this file exists. If it doesn't, it is created automatically.

The configuration file can be reset by running Dfetch with `--reset-config`.

The configuration file contains two sections:

* `ascii` - Controls the ASCII art displayed by Dfetch.
* `modules` - Controls which modules are displayed in what way and how info is formatted.

## ASCII

```json
"ascii": {
    "enabled": true,
    "path": "builtin",
    "padding_top": 1,
    "padding_bottom": 1
}
```

This section controls whether ASCII art is displayed, which ASCII art is used (`builtin`, distro name, or a custom path), and the padding above and below it.

## Modules

The `modules` array controls what Dfetch displays and in what order.

For example:

```json
{
    "name": "cpu",
    "label": "CPU",
    "color": "green",
    "format": "{short}",
    "separator": ":"
}
```

displays the CPU module using its short name.

### Available modules

| Module      | Description            |
| ----------- | ---------------------- |
| `userinfo`  | Username and hostname  |
| `os`        | Operating system       |
| `kernel`    | Current kernel         |
| `cpu`       | Processor information  |
| `memory`    | Memory usage           |
| `swap`      | Swap usage             |
| `local_ip`  | Local IP address       |
| `locale`    | System locale settings |
| `uptime`    | System uptime          |
| `battery`   | Battery information    |
| `bios`      | BIOS information       |
| `de`        | Desktop environment    |
| `wm`        | Window manager         |
| `shell`     | Current shell          |
| `terminal`  | Current terminal       |
| `disk`      | Disk usage             |
| `datetime`  | Current date and time  |
| `packages`  | Installed packages     |
| `host`      | Device model           |
| `board`     | Motherboard name       |
| `emptyline` | Inserts a blank line   |
| `text`      | Custom text            |
| `color`     | Terminal color palette |

### Common options

| Option      | Description                                              |
| ----------- | -------------------------------------------------------- |
| `name`      | The module that should be displayed                      |
| `label`     | The name displayed in output                             |
| `color`     | Color of the label default is set based on the distro    |
| `separator` | Text separating the label and the information            |
| `format`    | Controls how the module output is formatted              |
| `mount`     | Sets which mount point or filesystem should be displayed |
| `text`      | Used by the `text` module to set custom text             |

# Formatting

Module formatting uses `{placeholder}` syntax. Text outside of placeholders is preserved.

This is a list of the placeholders supported by every module:

### `userinfo`

* `{username}` - Current username
* `{hostname}` - System hostname

### `os`

* `{name}` - Operating system name

### `kernel`

* `{type}` - Kernel type
* `{release}` - Kernel release

### `cpu`

* `{name}` - Full CPU name
* `{short}` - Simplified CPU name

### `memory`

* `{memory}` - Used memory and total memory
* `{used}` - Used memory
* `{total}` - Total memory
* `{percent}` - Memory usage percentage
* `{unit}` - The unit in which the rest is displayed

### `swap`

* `{swap}` - Used swap and total swap
* `{used}` - Used swap
* `{total}` - Total swap
* `{percent}` - Swap usage percentage
* `{unit}` - The unit in which the rest is displayed

### `local_ip`

* `{ip}` - IP address without prefix
* `{prefix}` - subnet prefix
* `{address}` - IP address with prefix

### `uptime`

* `{uptime}` - Formatted uptime (e.g. `2 days, 4 hours, 12 mins`)
* `{centuries}` - Number of complete centuries in the uptime
* `{years}` - Number of complete years remaining after centuries
* `{months}` - Number of complete 30-day months remaining after years
* `{weeks}` - Number of complete weeks remaining after months
* `{days}` - Number of complete days remaining after weeks
* `{hours}` - Number of complete hours remaining after days
* `{minutes}` - Number of complete minutes remaining after hours


### `battery`

* `{percent}` - Battery percentage
* `{status}` - Battery status

### `bios`

* `{bios}` - Bios name

### `de`

* `{de}` - Desktop enviroment name

### `wm`

* `{name}` - Window manager name
* `{version}` - Window manager version
* `{sessiontype}` - Session type, such as `x11` or `wayland`

### `shell`

* `{name}` - Shell name
* `{version}` - Shell version

### `terminal`

* `{name}` - Terminal name
* `{version}` - Terminal version

### `disk`

* `{disk}` - Used space and total space
* `{used}` - Used space
* `{total}` - Total space
* `{unit}` - Storage unit
* `{percent}` - Disk usage percentage

The `mount` option can be used to select which filesystem is measured.

### `datetime`

* `{date}` - Current date
* `{time}` - Current time

### `packages`

* `{packages}` - Default formatted output
* `{total}` - Total number of installed packages
* `{dpkg, pacman, apk, eopkg, rpm, snap, flatpak}` - Total packages for specified package manager

### `host`

* `{name}` - Device model/family

### `board`

* `{board}` - Motherboard name

### `locale`

* `{locale}` - Language

### Modules without formatting

The following modules do not currently provide format placeholders:

* `text`
* `emptyline`
* `color`

The `text` module uses the `format` option for custom text but does not support any placeholders.

### Color in formatted information

Formatted information, such as

`{disk} ({percent}%)`

can be colored in the same way as ASCII art by providing color tags. For example, if we want to make the percent part of the information green, we can write:

`{disk} ${green}({percent}%)`

A [list of supported colors](#supported-colors) can be found below.


### mount

The `mount` option is only used by the `disk` module.

For example:

```json
{
    "name": "disk",
    "mount": "/home",
    "format": "{used} / {total} {unit}"
}
```

displays disk usage for `/home` instead of `/`.

## Custom ASCII art

To use custom ASCII art, create a text file containing your ASCII art.

You can optionally add color tags such as `${bright_blue}` inside the file.

```text
${bright_white}       _,met$$$$$gg.
${bright_white}   ,g$$$$$$$$$$$$$$$$P.
${bright_white} ,$$P'              `$$$.
${bright_white}',$$P       ,ggs.     `$$:
${bright_white}`d$$'     ,$P"'   ${bright_red}.    ${bright_white}$$$
${bright_white} $$P      d$'     ${bright_red},    ${bright_white}$$P
${bright_white} $$:      $$.   ${bright_red}-    ${bright_white},d$$'
${bright_white} $$;      Y$b._   _,d$P'
${bright_white} Y$$.    ${bright_red}`.${bright_white}`"Y$$$$P"'
${bright_white} `$$b      ${bright_red}"-.__
${bright_white}  `Y$$
${bright_white}   `Y$$.
${bright_white}     `$$b.
${bright_white}       `Y$$b.
${bright_white}          `"Y$b._
```

In the configuration file, change:

```json
"path": "builtin"
```

to:

```json
"path": "/path/to/ascii.txt"
```

to use the custom ASCII art.

## Supported colors

| Color     | Bright color      | Bold color      | Bright bold color      |
| --------- | ----------------- | --------------- | ---------------------- |
| 1 black   | 9 bright_black    | 17 bold_black   | 25 bold_bright_black   |
| 2 red     | 10 bright_red     | 18 bold_red     | 26 bold_bright_red     |
| 3 green   | 11 bright_green   | 19 bold_green   | 27 bold_bright_green   |
| 4 yellow  | 12 bright_yellow  | 20 bold_yellow  | 28 bold_bright_yellow  |
| 5 blue    | 13 bright_blue    | 21 bold_blue    | 29 bold_bright_blue    |
| 6 magenta | 14 bright_magenta | 22 bold_magenta | 30 bold_bright_magenta |
| 7 cyan    | 15 bright_cyan    | 23 bold_cyan    | 31 bold_bright_cyan    |
| 8 white   | 16 bright_white   | 24 bold_white   | 32 bold_bright_white   |

The number in front of a color can be used as a substitute for the color name.
