# My CLI

This is a command-line interface (CLI) tool built with Go to
streamline your SRE/DevOps workflows.
It provides a set of commands to manage your Git configurations,
Kubernetes contexts, and VPN connections.

## Features

- **Git:** Switch between different Git configurations and get the status of your repositories.
- **Kubernetes:** Switch between different Kubernetes contexts.
- **VPN:** Switch between different WireGuard VPN connections.

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

## Usage

### Git

- `cli git switch-config --name [config name]`
- `cli git get-config`
- `cli git status`

### Kubernetes

- `cli k8s switch-context`

### VPN

- `cli vpn switch-connect`

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
