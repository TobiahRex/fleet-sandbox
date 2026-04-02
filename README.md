# Fleet Sandbox: Conway's Game of Life (Go + WASM)

A browser-playable Conway's Game of Life built with Go, compiled to WebAssembly.

## Architecture

```
cmd/life/main.go          - CLI entry point (prints to terminal)
cmd/life-wasm/main.go     - WASM entry point (renders to canvas)
pkg/game/board.go         - Core game logic (Board, Cell, Step)
pkg/game/patterns.go      - Classic patterns (glider, blinker, pulsar, etc.)
pkg/game/board_test.go    - Unit tests
web/index.html            - Browser UI
web/wasm_exec.js          - Go WASM support
```

## Quick Start

```bash
# Run in terminal
go run cmd/life/main.go

# Build WASM
GOOS=js GOARCH=wasm go build -o web/life.wasm cmd/life-wasm/main.go

# Run tests
go test ./...
```

## Goals

This is a test project for validating Overlord's multi-agent fleet. Agents will:
- Add new patterns (developer)
- Review code quality and add tests (reviewer)
- Resolve any conflicts (merger)
