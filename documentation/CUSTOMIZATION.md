# Customization

Dfetch stores its configuration in `~/.config/Dfetch`. Every time Dfetch starts, it checks if this file exists. If it doesn't, a default configuration file is created automatically.

The configuration file contains two sections:

- `ascii` - Controls the ASCII art displayed by Dfetch.
- `modules` - Controls which modules are displayed and how they are formatted.

## ASCII

```json
"ascii": {
    "enabled": true,
    "path": "builtin",
    "padding_top": 1,
    "padding_bottom": 1
}
```

### `enabled`

Enables or disables the ASCII art.

### `path`

Sets which ASCII art Dfetch should use.

- `"builtin"` uses the built-in ASCII art.
- Any other value should be the path to a custom ASCII art file.

### `padding_top`

Adds empty lines above the ASCII art.

### `padding_bottom`

Adds empty lines below the ASCII art.

## Modules

The `modules` array controls what Dfetch displays and in what order.

For example,

```json
{
    "name": "cpu",
    "label": "CPU",
    "color": "green",
    "format": "short",
    "separator": ":"
}
```

displays the CPU module.

Modules are displayed in the same order they appear in the configuration file.

### Common options

#### `name`

The module Dfetch should display.

#### `label`

The text displayed before the module's information.

#### `color`

Sets the label color.

#### `separator`

Sets the text between the label and the information.

For example,

```
CPU: AMD Ryzen 7 250
```

uses `:` as the separator.

#### `format`

Changes how a module formats its output.

Only supported by:

- `kernel` supported format options: `short`
- `battery` supported format options: `short`
- `cpu` supported format options: `short`
- `memory` supported format options: `short`
- `swap` supported format options: `short`
- `uptime`supported format options: `short`
- `bios`supported format options: `short`
- `desktop` supported format options: `short`
- `shell` supported format options: `short`
- `terminal` supported format options: `short`
- `disk` supported format options: `short`
- `datetime` supported format options: `time`, `date`
- `packages` supported format options: `short`
- `host` supported format options: `short`

#### `mount`

Only used by the `disk` module.

Sets which mount point or file system should be displayed.

#### `text`

Used by the text module to set a custom text.

## Available modules

| Module        | Description                           |
| ---------------| ---------------------------------------|
| `userinfo`    | Username and hostname                 |
| `os`          | Operating system                      |
| `kernel`      | Current kernel                        |
| `cpu`         | Processor information                 |
| `memory`      | Memory usage                          |
| `swap`        | Swap usage                            |
| `local_ip`    | Local IP address                      |
| `uptime`      | System uptime                         |
| `battery`     | Battery information                   |
| `bios`        | BIOS information                      |
| `desktop`     | Desktop environment or window manager |
| `shell`       | Current shell                         |
| `terminal`    | Current terminal                      |
| `disk`        | Disk usage                            |
| `time`        | Current time                          |
| `date`        | Current date                          |
| `packages`    | Installed packages                    |
| `host`        | Device model                          |
| `motherboard` | Motherboard name                      |
| `emptyline`   | Inserts a blank line                  |
| `text`        | Custom text                           |

## Custom ASCII art

To use a custom ASCII art, create a text file containing your ASCII art.

You can also add color tags such as `${bright_blue}` inside the file.

```
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

Change

```json
"path": "builtin"
```

to

```json
"path": "/path/to/ascii.txt"
```

to use it.

## Supported colors

```
black
red
green
yellow
blue
magenta
cyan
white

bright_black
bright_red
bright_green
bright_yellow
bright_blue
bright_magenta
bright_cyan
bright_white
```