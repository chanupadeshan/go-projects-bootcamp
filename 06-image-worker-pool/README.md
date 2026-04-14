# 🏭 Image Worker Pool

Process images using a worker pool pattern. Learn `context`, `select`, and buffered channels.

## Core Concepts
- `context` — Request cancellation and timeouts
- `select` — Multiplexing on channels
- Buffered channels — Asynchronous communication
- Worker pool pattern — Efficient resource management

## Running
```bash
go run main.go
```

## Features
- Load images concurrently
- Apply filters
- Save results to `output/` directory
