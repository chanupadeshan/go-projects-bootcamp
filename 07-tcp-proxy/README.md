# 🔌 TCP Proxy / Middleware

A TCP proxy that intercepts and logs traffic. Demonstrates `io.Reader/Writer`, middleware patterns, and `MultiWriter`.

## Core Concepts
- `io.Reader` / `io.Writer` — Universal I/O interfaces
- Middleware pattern — Wrapping and intercepting data flow
- `io.MultiWriter` — Writing to multiple destinations
- Raw TCP — How proxies, VPNs, and load balancers work

## Running
```bash
go run main.go
```

## Features
- Forward TCP connections
- Log all traffic
- Middleware layer for data interception
