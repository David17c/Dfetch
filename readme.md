Minimal lighweight tool inspired by [Neofetch](https://github.com/dylanaraps/neofetch) displaying your system information. 

Dfetch currently only works on Linux based operating systems.
Its far from done and not yet ready for use.

### To do
* [ ] Add configuration system
* [ ] Add colors to ASCII art
* [ ] Add more types of system info
* [ ] Test on more distro's
* [ ] Add support for more distro's
* [ ] Improve error handeling
* [ ] Write better README.md
* [ ] Add command-line flags

```
Dfetch
├── getsysinfo
│   ├── getcpu.go       # Fetch cpu info
│   ├── getdistro.go    # Fetch distro information
│   ├── gethostname.go  # Fetch hostname
│   ├── getkernel.go    # Fetch info about kernel version
│   ├── getlocalip.go   # Fetch your local IP
│   ├── getmemory.go    # Fetch memory related info
│   ├── getuptime.go    # Fetch computer uptime
│   └── getusername.go  # Fetch current users name
│
├── go.mod              # Go module file
├── images              # Some images used in readme
├── LICENSE
│
├── logo                # ASCII art distro logos
│   ├── arch.txt
│   ├── debian.txt
│   ├── linuxmint.txt
│   ├── ubuntu.txt
│   └── ...
│
├── main.go             # Start and end of the program
└── readme.md           # Repo descripting
```