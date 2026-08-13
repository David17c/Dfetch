# Customization

The configuration file is stored in `~/.config/Dfetch`. Every time Dfetch is started, it checks if this file exists. If it doesn't, a default configuration file is created automatically.

The configuration file can be reset by running Dfetch with the `--reset-config`.

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

This section allows you to set: if a logo is displayed, which is displayed (builtin, distro name or path to custom) and the distance between the logo and the top or bottom (padding).

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

### Available modules

| Module      | Description             |
| ------------| ------------------------|
|  userinfo   | Username and hostname   |
|  os         | Operating system        |
|  kernel     | Current kernel          |
|  cpu        | Processor information   |
|  memory     | Memory usage            |
|  swap       | Swap usage              |
|  local_ip   | Local IP address        |
|  locale     | Systems locale settings |
|  uptime     | System uptime           |
|  battery    | Battery information     |
|  bios       | BIOS information        |
|  de         | Desktop environment     |
|  wm         | Window manager          |
|  shell      | Current shell           |
|  terminal   | Current terminal        |
|  disk       | Disk usage              |
|  time       | Current time            |
|  date       | Current date            |
|  packages   | Installed packages      |
|  host       | Device model            |
|  board      | Motherboard name        |
|  emptyline  | Inserts a blank line    |
|  text       | Custom text             |
|  color      | Terminal color pallet   |

### Common options

| Option     | Description                                               |
|------------|-----------------------------------------------------------|
| name       | The module that should be displayed                       |
| label      | The name displayed in output                              |
| color      | Color of the label                                        |
| separator  | Text seperating the label and the info                    |
| format     | Changes how a module formats its output                   |
| mount      | Sets which mount point or file system should be displayed |
| text       | Used by the text module to set a custom text              |

### Notes

Format is only supported by

- `battery` supported format options: `short`
- `cpu` supported format options: `short`
- `memory` supported format options: `long`
- `swap` supported format options: `long`
- `wm` supported format options: `short`
- `shell` supported format options: `short`
- `disk` supported format options: `long`
- `datetime` supported format options: `time`, `date`
- `packages` supported format options: `short`

and mount only works on the disk module

## Custom ASCII art

To use a custom ASCII art, create a text file containing your ASCII art.

You can then optionally add color tags such as `${bright_blue}` inside the file.

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

In the configuration file change

```json
"path": "builtin"
```

to

```json
"path": "/path/to/ascii.txt"
```

to use it.

## Supported colors

| Color     | Bright color      | Bold color      | Bright bold color      |
|-----------|-------------------|-----------------|------------------------|
| 1 black   | 9 bright_black    | 17 bold_black   | 25 bold_bright_black   |
| 2 red     | 10 bright_red     | 18 bold_red     | 26 bold_bright_red     |
| 3 green   | 11 bright_green   | 19 bold_green   | 27 bold_bright_green   |
| 4 yellow  | 12 bright_yellow  | 20 bold_yellow  | 28 bold_bright_yellow  |
| 5 blue    | 13 bright_blue    | 21 bold_blue    | 29 bold_bright_blue    |
| 6 magenta | 14 bright_magenta | 22 bold_magenta | 30 bold_bright_magenta |
| 7 cyan    | 15 bright_cyan    | 23 bold_cyan    | 31 bold_bright_cyan    |
| 8 white   | 16 bright_white   | 24 bold_white   | 32 bold_bright_white   |

The number in front of the color can be used as a substitute for the color.