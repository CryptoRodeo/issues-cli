# Konflux Issues CLI

A command-line interface for managing Konflux issues, written in Go. This CLI tool provides a simple and efficient way to interact with the Konflux Issues API.

## Features

- List, filter, and search issues
- Get detailed information about specific issues
- Works as a standalone CLI and as a kubectl plugin
- Automatically detects and uses the current kubectl namespace
- Configurable API endpoint

## Installation

### Option 1: Build from source

```bash
# Requires Go 1.16+
git clone https://github.com/CryptoRodeo/issues-cli.git
cd issues-cli
go build -o konflux-issues main.go
chmod +x konflux-issues
sudo mv konflux-issues /usr/local/bin/
```

### Setup as kubectl plugin

```bash
# Create a symlink with the kubectl- prefix
ln -s /usr/local/bin/konflux-issues /usr/local/bin/kubectl-issues

# Verify plugin installation
kubectl plugin list
```

## Usage

### Standalone CLI

```bash
# List issues in a namespace
konflux-issues list -n team-alpha

# Filter issues by type
konflux-issues list -n team-alpha -t build

# Filter issues by severity
konflux-issues list -n team-alpha -s critical

# Get details for a specific issue
konflux-issues details -i failed-build-frontend -n team-alpha

# Configure the API URL
konflux-issues config set-api-url http://localhost:3000/api/v1

# Show current configuration
konflux-issues config

# Reset configuration to defaults
konflux-issues config reset
```

### As a kubectl plugin

```bash
# List issues in current kubectl context namespace
kubectl issues list

# List issues in a specific namespace
kubectl issues list -n team-alpha

# Get details for an issue
kubectl issues details -i failed-build-frontend
```

## Configuration

The CLI uses a configuration file stored at `~/.konflux-issues/config.yaml`. You can modify settings using the `config` command or by directly editing this file.

Default configuration:
```yaml
api_url: http://localhost:3000/api/v1
```

You can also set the API URL using the `KONFLUX_API_URL` environment variable.

## Development

### Prerequisites

- Go 1.16 or later
- Make (optional)

### Building

```bash
# Build the binary
go build -o konflux-issues main.go

# Run tests
go test ./...

# Install locally
go install
```

## License

MIT
