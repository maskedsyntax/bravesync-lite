# BraveSync Lite

A local-first, security-focused backup and restore tool for Brave browser data on Linux. This tool allows you to securely archive your bookmarks and password database, encrypt them locally, and sync the encrypted blobs to a private GitHub repository.

## Features

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

Before using the tool, you must configure your GitHub repository settings and Personal Access Token (PAT).

```bash
./bravesynclite config --repo <your-private-repo-url> --local ~/.local/share/brave-sync-repo --pat <your-github-pat>
```

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
