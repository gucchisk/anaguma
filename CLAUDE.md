# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Anaguma is a CLI tool for accessing BadgerDB databases, supporting both BadgerDB v1 and v2. The project is written in Go and uses the Cobra CLI framework.

## Build and Development Commands

### Building
```bash
make build          # Build binary to bin/anaguma
go build -v -o bin/anaguma github.com/gucchisk/anaguma
```

### Testing
```bash
go test ./...       # Run all tests
go test ./cmd       # Run tests for cmd package
```

### Release Management
```bash
make v0.3.6         # Set version and prepare release (calls scripts/release/main.go)
make pack           # Package release (calls scripts/package/main.go)
```

### Cleanup
```bash
make clean          # Remove bin/ directory
```

## Code Architecture

### Core Structure
- `main.go` - Entry point that delegates to cmd package
- `cmd/` - CLI commands using Cobra framework
  - `root.go` - Root command with global flags and DB initialization
  - `get.go`, `set.go`, `keys.go`, `values.go` - Individual commands
  - `version.go` - Version command
- `common/` - Shared interfaces and utilities
  - `db.go` - DB interface abstraction
  - `format.go` - Input/output format handling
  - `version.go` - Version management
- `db/` - Database implementations
  - `v1/db.go` - BadgerDB v1 implementation
  - `v2/db.go` - BadgerDB v2 implementation

### Database Abstraction
The project uses a common interface (`common.DB`) to abstract BadgerDB operations:
- `Open(dir string, log bool) error` - Open database
- `Close()` - Close database
- `Get(key []byte, fn func(value []byte) error) error` - Get value by key
- `Set(key, value []byte) error` - Set key-value pair
- `Keys(fn func(keys [][]byte)) error` - Iterate over keys
- `Values(fn func(keys, values [][]byte)) error` - Iterate over key-value pairs

### CLI Flags and Configuration
Global flags (available on all commands):
- `-b, --bversion uint8` - BadgerDB version [1, 2] (default: 1)
- `-o, --out string` - Output format [ascii, hex, base64] (default: "ascii")
- `-l, --log` - Enable log output
- `-v, --version` - Show version

Command-specific flags:
- `get` and `set` commands: `-i, --in string` - Input format (default: "ascii")

### Database Version Handling
The application selects the appropriate BadgerDB implementation based on the `--bversion` flag:
- Version 1: Uses `github.com/dgraph-io/badger` (v1 API)
- Version 2: Uses `github.com/dgraph-io/badger/v2` (v2 API)

The main differences in implementation:
- V1: Uses `opts.WithLogger(nil)` to disable logging
- V2: Uses `opts.WithLoggingLevel(badger.ERROR)` to set logging level

### Testing
- Tests are located alongside source files with `_test.go` suffix
- Currently has basic CLI command testing in `cmd/root_test.go`
- Tests verify help command behavior and output consistency

### Version Management
Version is managed through:
- `scripts/release/main.go` - Sets version in source code
- `common/version.go` - Stores and retrieves version information
- Makefile targets for version bumping (`make v<version>`)