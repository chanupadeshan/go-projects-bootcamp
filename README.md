# 🐹 Go Projects Bootcamp

> A hands-on journey through Go — from the basics to TCP networking — built entirely through projects.

![Go](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go&logoColor=white)
![Projects](https://img.shields.io/badge/Projects-7-success?style=flat)
![Level](https://img.shields.io/badge/Level-Beginner%20→%20Intermediate-blue?style=flat)
![Status](https://img.shields.io/badge/Status-In%20Progress-yellow?style=flat)

---

## 🎯 Goal

Instead of reading theory, I learned Go by building real projects — each one introducing a core concept of the language. This repo is my learning log, going from "Hello World" thinking to writing concurrent servers and TCP proxies.

---

## 🗂️ Project List

| # | Project | Core Concept | Difficulty |
|---|---------|-------------|------------|
| [01](./01-smart-guessing-game/) | 🎯 Smart Guessing Game | `package main`, `:=`, `for` loop | ⭐ Beginner |
| [02](./02-inventory-manager/) | 📦 Inventory Manager | `struct`, `map`, `ok` pattern | ⭐ Beginner |
| [03](./03-shape-area-calculator/) | 📐 Shape Area Calculator | `interface`, implicit implementation | ⭐⭐ Easy |
| [04](./04-json-dictionary-api/) | 📖 JSON Dictionary API | `net/http`, `encoding/json`, struct tags | ⭐⭐ Easy |
| [05](./05-concurrent-web-pinger/) | 🌐 Concurrent Web Pinger | `goroutine`, `WaitGroup`, `channel` | ⭐⭐⭐ Intermediate |
| [06](./06-image-worker-pool/) | 🏭 Image Worker Pool | `context`, `select`, buffered channels | ⭐⭐⭐ Intermediate |
| [07](./07-tcp-proxy/) | 🔌 TCP Proxy / Middleware | `io.Reader/Writer`, middleware, `MultiWriter` | ⭐⭐⭐⭐ Advanced |

---

## 💡 What I Learned

### The Basics
- Go's **single loop** (`for`) handles while, for, and foreach
- **`:=`** lets Go infer types — no boilerplate declarations
- **Structs + Methods** replace classes without needing OOP

### Go's Data Structures
- **Maps** (`map[string]Item{}`) — key/value stores with safe lookup using the `ok` pattern
- **Slices** — dynamic arrays that grow with `append()`
- **Interfaces** — define behavior, not data. No `implements` keyword needed

### Concurrency (Go's Superpower)
- **Goroutines** — lightweight threads launched with `go`. Thousands can run at once
- **Channels** — pipes for safely passing data between goroutines
- **WaitGroups** — coordinate when all goroutines have finished
- **Context** — a kill switch that cancels goroutines across your entire program

### Systems Programming
- **`io.Reader` / `io.Writer`** — universal interfaces for anything that reads or writes data
- **Middleware pattern** — wrap a writer to intercept data without changing the flow
- **Raw TCP** — how VPNs, load balancers, and proxies work under the hood

---

## 🚀 Running Any Project

Each project is fully self-contained with its own Go module.

```bash
# Clone the repo
git clone https://github.com/YOUR_USERNAME/go-projects-bootcamp.git
cd go-projects-bootcamp

# Run any project
cd 01-smart-guessing-game
go run main.go
```

> **Requirement:** Go 1.21+ → [Download here](https://go.dev/dl/)

---

## 📁 Folder Structure

```
go-projects-bootcamp/
├── README.md
├── .gitignore
├── 01-smart-guessing-game/
│   ├── main.go
│   └── README.md
├── 02-inventory-manager/
│   ├── main.go
│   └── README.md
├── 03-shape-area-calculator/
│   ├── main.go
│   └── README.md
├── 04-json-dictionary-api/
│   ├── main.go
│   └── README.md
├── 05-concurrent-web-pinger/
│   ├── main.go
│   └── README.md
├── 06-image-worker-pool/
│   ├── main.go
│   └── output/           ← generated at runtime
│   └── README.md
└── 07-tcp-proxy/
    ├── main.go
    └── README.md
```

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Standard Library only** — no external frameworks or dependencies
- Packages used: `fmt`, `net/http`, `encoding/json`, `sync`, `context`, `io`, `image`, `net`, `math/rand`, `time`

---

## 📈 Progress

- [x] Project 1 — Smart Guessing Game
- [x] Project 2 — Inventory Manager
- [x] Project 3 — Shape Area Calculator
- [x] Project 4 — JSON Dictionary API
- [x] Project 5 — Concurrent Web Pinger
- [x] Project 6 — Image Worker Pool
- [x] Project 7 — TCP Proxy / Middleware

---

## 📚 Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Tour](https://go.dev/tour/)

---

> *"Don't communicate by sharing memory; share memory by communicating."* — Go Proverb