# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**clist** is a CLI tool for to-do lists and notes, written in Go.

## Development Setup

This is a Go project. The standard Go development workflow applies:

```bash
# Initialize Go module (if not already done)
go mod init github.com/your-username/clist

# Build the project
go build

# Run the project
go run .

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Install dependencies
go mod tidy

# Format code
go fmt ./...

# Lint code (requires golangci-lint)
golangci-lint run
```

## Project Structure

This is an early-stage project. The typical Go CLI project structure would include:

- `main.go` - Entry point for the CLI application
- `cmd/` - Command definitions (if using a library like Cobra)
- `internal/` - Internal packages not meant for external use
- `pkg/` - Public packages that could be imported by other projects
- `go.mod` - Go module definition
- `go.sum` - Dependency checksums

## Architecture Notes

As this project develops, consider:

- Using a CLI library like [Cobra](https://github.com/spf13/cobra) for command structure
- Implementing data persistence (file-based storage, SQLite, etc.)
- Following Go project layout standards from https://github.com/golang-standards/project-layout