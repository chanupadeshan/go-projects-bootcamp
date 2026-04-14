# 🌐 Concurrent Web Pinger

Ping multiple websites concurrently using goroutines. Demonstrates Go's concurrency primitives.

## Core Concepts
- `goroutine` — Lightweight threads (with `go` keyword)
- `WaitGroup` — Synchronization primitive
- `channel` — Communication between goroutines

## Running
```bash
go run main.go
```

## Features
- Ping multiple URLs concurrently
- Timeout handling
- Summary report of results
