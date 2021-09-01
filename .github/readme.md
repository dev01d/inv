# Inv

A formatted investigation tool.

`inv` helps you `inv`estigate IPs/hosts and DNS records then returns the results in a copy/pasteable format.

## Usage

<p align="center">
    <img src=".github/assets/demo.gif" max-width="100%" alt="inv command usage GIF" />
</p>

```shell
Usage: inv [OPTIONS] COMMAND

-h, --help                   Show this help
-u, --up=IP                  Check liveness stats
-d, --dig=domain             Dig DNS records
-w, --whois=domain           Whois domain information
    --version                Print version
```

## Install

### Mac

via Homebrew:

```shell
brew install dev01d/tap/inv
```

## Linux

- DEB

```shell
sudo echo "deb [trusted=yes] https://apt.fury.io/dev01d/ /" > /etc/apt/sources.list.d/fury.list

sudo apt-get update; sudo apt install inv
```

- YUM

```shell
sudo echo """\
[fury]
name=Gemfury Private Repo
baseurl=https://yum.fury.io/dev01d/
enabled=1
gpgcheck=0
""" > /etc/yum.repos.d/fury.repo

yum install inv
```

### Go

```shell
go install github.com/dev01d/inv@latest
```

Download from [release page](https://github.com/dev01d/inv/releases)
