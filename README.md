# Realtime App using Golang(Backend) and React(Frontend)

This repository contains a small real-time backend service written in Go.  
The goal of the project is to demonstrate how to structure a Go application properly while handling real-time data using WebSockets and deadline-aware processing.

The project is intentionally kept simple and explicit. It focuses on correctness, predictability, and clean structure rather than framework magic.

---

## Motivation

Many so-called “real-time” web applications are simply fast request/response systems. In practice, real-time systems are defined by timing guarantees and controlled behavior under load.

This project explores those ideas in a web context:
- bounded queues instead of unbounded buffering
- explicit deadlines and cancellation
- dropping late work instead of blocking the system
- graceful shutdown and resource cleanup

---

## Project Structure

```
realtime-app/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── app/
│   ├── domain/
│   ├── service/
│   └── transport/
├── configs/
├── scripts/
├── docs/
├── test/
├── go.mod
├── go.sum
└── README.md
```

---

## Requirements

- Go 1.21 or newer
- Git

---

## Setup

```bash
git clone https://github.com/<your-username>/realtime-app.git
cd realtime-app
go mod tidy
```

---

## Running the Application

```bash
go run ./cmd/server
```

Default address:
```
http://localhost:8080
```

---

## Endpoints

| Method | Path      | Description                   |
|------|-----------|-------------------------------|
| GET  | /health   | Health check                  |
| GET  | /ws       | WebSocket connection endpoint |

---

## WebSocket Behavior

Clients connect using:
```
ws://localhost:8080/ws
```

Events are streamed in real time. Each event has a deadline. If processing exceeds the deadline, the event is discarded.

---

## Configuration

Environment variables:

| Variable   | Default | Description             |
|-----------|---------|-------------------------|
| APP_PORT  | 8080    | HTTP server port        |
| APP_ENV   | dev     | Application environment |

---

## Design Principles

- Predictability over raw throughput
- Bounded channels
- Context-based cancellation
- Minimal dependencies
- Explicit error handling

---

## Testing

```bash
go test ./...
go test ./... -race
```

---

## Dependencies

- Go standard library
- github.com/gorilla/websocket

---

## License

MIT License.

---

## Author

Maintained by a Aziz AMMARI (A Go backend engineer).
