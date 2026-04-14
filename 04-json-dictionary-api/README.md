# 📖 JSON Dictionary API

A simple HTTP API that serves dictionary definitions as JSON. Learn `net/http`, JSON encoding, and struct tags.

## Core Concepts
- `net/http` — Building HTTP servers
- `encoding/json` — Marshaling and unmarshaling JSON
- Struct tags — Controlling JSON serialization

## Running
```bash
go run main.go
```

Then visit `http://localhost:8080/define?word=golang`

## Endpoints
- `GET /define?word=<word>` — Get definition of a word
