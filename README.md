# My CLI

This is a command-line interface (CLI) tool built with Go to
streamline your SRE/DevOps workflows.
It provides a set of commands to manage your Git configurations,
Kubernetes contexts, and VPN connections.

## Features

- **Git:** Switch between different Git configurations and get the status of your repositories.
- **Kubernetes:** Switch between different Kubernetes contexts.
- **VPN:** Switch between different WireGuard VPN connections.

## Prerequisites

This CLI tool relies on several external tools for its functionality.
Please ensure these are installed and configured on your system
before using the respective commands:

- **Git Commands (`cli git`):** Requires `git` to be installed and configured.
- **Kubernetes Commands (`cli k8s`):** Requires `kubectl` to be installed and
  a valid Kubernetes configuration (kubeconfig) in the path specified in `cliconfig.yaml`.
- **VPN Commands (`cli vpn`):** Requires `WireGuard` and `wg-quick`
  to be installed and configured. WireGuard configuration files should be present in the path specified in `cliconfig.yaml`.

## Installation

To install the CLI, you can use the provided `makefile`:

```bash
make install
```

This will build the binary and move it to `/usr/bin/cli`.

## Configuration

Before using the CLI, you need to create a configuration file.
You can do this by running:

```bash
cli init
```

This will create a `cliconfig.yaml` file in `~/.config/`.
You can also specify a different path using the `--path` or `-p` flag.

Here is an example of the `cliconfig.yaml` file:

```yaml
mainPath: /path/to/your/projects
toCheck:
  - path: ""
    is_repo: true
gitConfigs:
  work: /path/to/your/work/gitconfig
  perso: /path/to/your/perso/gitconfig
kubeconfigPath: /path/to/your/kube/configs
vpnConfigPath: /etc/wireguard
```

## Command Reference

### cli init

Initialize the CLI configuration file.

```bash
cli init [--path/-p <path-to-config>]
```

### cli git

Manage Git configurations and repository status.

#### cli git switch-config

Switch between different Git configurations.

```bash
cli git switch-config # Interactive selection
```

#### cli git get-config

Display the current Git configuration.

```bash
cli git get-config
```

#### cli git status

Show the status of Git repositories.

```bash
cli git status
```

### cli k8s

Work with Kubernetes, switch contexts, etc.

#### cli k8s context

Switch Kubernetes context.

```bash
cli k8s context [--name/-n <context-name>]
```

### cli vpn

Manage VPN connections.

#### cli vpn connect

Connect to a VPN endpoint.

```bash
cli vpn connect [--name/-n <vpn-name>]
```

#### cli vpn disconnect

Disconnect from the active VPN endpoint.

```bash
cli vpn disconnect
```

## Building from Source

To build the CLI from source, you can use the following command:

```bash
make build
```

This will create a binary in the `bin` directory.

## Running the CLI

To run the CLI after building it, you can use the following command:

```bash
make run
```
