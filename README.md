```
██████  ███████ ███████ ████████  ██████ ██   ██ 
██   ██ ██      ██         ██    ██      ██   ██ 
██   ██ █████   █████      ██    ██      ███████ 
██   ██ ██      ██         ██    ██      ██   ██ 
██████  ██      ███████    ██     ██████ ██   ██ 
```

A minimal lighweight command-line tool inspired by [Neofetch](https://github.com/dylanaraps/neofetch). Dfetch shows information relating to your OS, hardware and software and has a focus on being minimal, simple and lightweight.

### Example output

```
 ___________    David17c@Thinkpad
|_          \   --------------
  | | _____ |   OS: LMDE 7 (gigi)
  | | | | | |   Kernel: 6.12.86+deb13-amd64
  | | | | | |   CPU: AMD Ryzen 7 250 w/ Radeon 780M Graphics
  | \_____/ |   Memory: 3.55 / 29.03 GB (12%)
  \_________/   Local IP (IPv4): 192.168.1.107
                Uptime: 6h 59m 16s
```

### File structure

```
Dfetch
├── getsysinfo
│   ├── cpu.go       # CPU information
│   ├── distro.go    # Linux distribution
│   ├── hostname.go  # System hostname
│   ├── kernel.go    # Kernel version
│   ├── localip.go   # Local IP address
│   ├── memory.go    # Memory usage
│   ├── uptime.go    # System uptime
│   └── username.go  # Current username
│
├── go.mod           # Go module config
├── LICENSE          # Project license
│
├── logo             # ASCII logos
│   ├── arch.txt
│   ├── debian.txt
│   ├── linuxmint.txt
│   ├── ubuntu.txt
│   └── ...
│
├── main.go          # Start / end of program
└── README.md        # Project overview
```

### To do

* [ ] Add configuration system
* [ ] Add colors to ASCII art
* [ ] Test on more distro's
* [ ] Add support for more distro's
* [ ] Improve error handeling
* [ ] Add command-line flags