# Biller - A Go-powered Application

## Overview

This is a simple billing application built with Go, a powerful and efficient programming language. This project serves as a starting point for exploring Go's capabilities in building apps.

## Prerequisites

- Go 1.20 or later
- Make (optional, for using Makefile commands)

## Installation

1. Install Go from [golang.org](https://golang.org/dl/)

2. Clone the repository:

   ```bash
   git clone https://github.com/sam3177/go-biller.git
   ```

3. Install dependencies:
   ```bash
   make deps
   ```
   Or without Make:
   ```bash
   go mod download
   go mod tidy
   ```

## Usage

The project includes several Make commands for common tasks:

```bash
make build    # Build the application (output to bin/bill)
make run      # Run the application
make run-terminal-printer # Run the application with terminal printer
make run-epson-printer # Run the application with a physical epson printer
make test     # Run all tests
make clean    # Clean build artifacts
make lint     # Run the linter
make deps     # Install dependencies
```

I did both types of printer implementations to show the power of interfaces in Go. The terminal printer is a simple implementation that prints to the terminal, while the Epson printer implementation uses the Epson ePOS SDK to print to a physical printer.

If you don't have Make installed, you can use the equivalent Go commands:

```bash
go build -o bin/bill cmd/bill/main.go  # Build
go run cmd/bill/main.go                # Run
go test ./...                          # Test
go clean                               # Clean
```
