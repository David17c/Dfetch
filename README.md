![Dfetch banner](images/Dfetch_banner.png)

A minimal lighweight command-line tool inspired by [Neofetch](https://github.com/dylanaraps/neofetch) written in Go. Dfetch shows information relating to your OS, hardware and software in a minimal and visually nice way.

![Output example](images/output_example.png)

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
├── images           # README images
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