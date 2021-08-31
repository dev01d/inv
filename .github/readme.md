# Inv

A formatted investigation tool.

`inv` helps you `inv`estigate IPs/hosts and DNS records then returns the results in a copy/pasteable format.

## Usage

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

### Go

```shell
go install github.com/dev01d/inv@latest
```

Download from [release page](https://github.com/dev01d/inv/releases)
