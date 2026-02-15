# qfm

A blazing-fast file manager for [Quickshell](https://github.com/quickshell-mirror/quickshell)-based desktop environments.

Linux-only. Wayland-only. Built with Go and Qt6/QML via [MIQT](https://github.com/mappu/miqt).

## Status

Early development.

## Requirements

- Linux kernel 6.6+ (LTS)
- Go 1.25+
- [mise](https://mise.jdx.dev) for tool management

## Quick Start

```bash
# Install dev tools
mise install

# Build
task build

# Run
./bin/qfm --version

# Run full CI checks
task ci
```

## Development

```bash
# Live reload
task dev

# Run tests
task test

# Lint
task lint

# See all tasks
task --list
```

## Architecture

Three-layer stack: **QML Frontend** → **Go Backend** → **io_uring / Linux Kernel**

## License

TBD
