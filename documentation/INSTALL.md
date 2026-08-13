## Debian

Dfetch provides prebuilt packages for Debian and Debian-based distributions.

Download the `.deb` package from the [Releases page](https://github.com/David17c/Dfetch/releases) and install it. Alternatively, you can download the prebuilt binary from the Releases page, which works on Debian-based distributions as well.

## Fedora

Dfetch provides prebuilt packages for Fedora and Fedora-based distributions.

Download the `.rpm` package from the [Releases page](https://github.com/David17c/Dfetch/releases) and install it. Alternatively, you can download the prebuilt binary from the Releases page, which works on Fedora-based distributions as well.

## Arch

Dfetch does not currently provide a prebuilt package for Arch or Arch-based distributions. Instead, you can either:

* Download the prebuilt binary from the [Releases page](https://github.com/David17c/Dfetch/releases), or
* Build Dfetch from source by following the [build instructions](#build-from-source).

## Build from source

### Requirements

- Go 1.24.4+
- Git

### Build

Clone the repository and build the binary:

```bash
git clone https://github.com/David17c/Dfetch.git
cd Dfetch
go build .
```