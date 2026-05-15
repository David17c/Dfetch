![Dfetch banner](images/Dfetch_banner.png)

Minimal lighweight tool written in Go displaying your system information. 

Dfetch currently only works on Linux based operating systems.
Its far from done and not yet ready for use.

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
* [ ] Add more types of system info
* [ ] Test on more distro's
* [ ] Add support for more distro's
* [ ] Improve error handeling
* [ ] Write better README.md
* [ ] Add command-line flags