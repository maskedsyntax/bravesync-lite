# BraveSync Lite

A local-first, security-focused backup and restore tool for Brave and Helium browser data on Linux. This tool allows you to securely archive your bookmarks and password database, encrypt them locally, and sync the encrypted blobs to a private GitHub repository.

## Features

- **Multi-Browser Support**: Supports both **Brave** and **Helium** browsers.
- **Local-First Security**: Data is encrypted on your machine before being uploaded.
- **AES-256-GCM Encryption**: Uses industry-standard encryption with Argon2id for key derivation.
- **Minimalist Sync**: Only syncs Bookmarks and Login Data (passwords).
- **TUI & CLI**: Offers both a terminal user interface and standard command-line operations.

## Installation

Ensure you have Go installed on your Linux system, then build the binary:

```bash
go build -o bravesynclite ./cmd/bravesynclite/main.go
```

## Setup

Before using the tool, you must configure your GitHub repository settings. You can do this interactively or via flags.

### Interactive (Recommended)
Simply run the config command and follow the prompts. You will be asked to choose between Brave and Helium:
```bash
./bravesynclite config
```

### Via Flags
```bash
./bravesynclite config --browser Helium --repo <your-private-repo-url> --pat <your-github-pat>
```
*(The `--local` path defaults to `~/.local/share/bravesync-repo` if omitted.)*

Alternatively, you can set the `BRAVE_SYNC_PAT` environment variable instead of storing it in the config file.

## Usage

### TUI Mode (Recommended)
Launch the interactive menu:
```bash
./bravesynclite tui
```

### CLI Mode
Perform operations directly via commands:

**Backup:**
```bash
./bravesynclite backup
```

**Restore:**
```bash
./bravesynclite restore
```

## Security Design

- **Zero-Knowledge**: The encryption password is never stored and is cleared from memory after use.
- **Encrypted Archives**: Only `.enc` files are stored in the remote repository.
- **Privacy**: No telemetry, analytics, or external network calls are made except to official GitHub APIs.
