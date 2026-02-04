# Networked Tic-Tac-Toe in Go

A simple client–server implementation of the classic Tic-Tac-Toe game written in **Go**. The server runs as a single instance, listens for incoming TCP connections, pairs clients automatically, and hosts multiple concurrent games. Clients connect via TCP and play using a minimal command-line interface.

---

##  Key Features

* Single-instance TCP server
* Automatic client matchmaking
* Multiple concurrent games using goroutines
* Server-authoritative game state (prevents cheating)
* Simple JSON-based protocol
* Minimal CLI client (no UI dependencies)

---

##  Architecture Overview

### High-Level Design

```
Client ──┐
         ├── TCP ──> Server ──> Game Loop (goroutine)
Client ──┘
```

### Server Responsibilities

* Accept incoming TCP connections
* Pair clients using a matchmaking queue
* Maintain the authoritative game state
* Validate moves
* Broadcast board updates and game results

### Client Responsibilities

* Connect to the server
* Display game state in the terminal
* Send move input to the server
* Receive opponent moves and game outcome

---

##  Server-Authoritative Model

Clients only send **intent** (the move index). The server:

* Validates whether the move is legal
* Applies the move to the board
* Determines win or draw conditions
* Broadcasts updates to both players

This ensures fairness and simplifies client logic.

---

##  Concurrency Model

* A buffered channel is used for matchmaking
* Each game runs in its own goroutine
* Shared state within a game is protected using a mutex

This allows the server to safely handle many games concurrently.

---

##  Project Structure

```
.
├── client/
│   └── main.go      # CLI client
├── server/
│   └── main.go      # TCP server & game loop
├── shared/
│   └── game.go      # Shared protocol & game logic
├── go.mod
└── README.md
```

---

## ▶ How to Run

### Prerequisites

* Go 1.20 or newer

### Start the Server

```bash
go run server/main.go
```

### Start Clients (in separate terminals)

```bash
go run client/main.go
```

* Start at least **2 clients** to begin a game
* Additional clients will be automatically paired

---

##  How to Play

* Board positions are indexed from **0 to 8**:

```
 0 | 1 | 2
---+---+---
 3 | 4 | 5
---+---+---
 6 | 7 | 8
```

* Enter a number (0–8) to make a move
* The server enforces turns and validates moves

---

##  Error Handling

* Invalid moves are rejected by the server
* Clients receive error messages without crashing
* Games terminate cleanly on win or draw

---

##  Possible Enhancements

These were intentionally left out to keep the scope reasonable:

* Move timeouts
* Client reconnection handling
* WebSocket or HTTP transport
* Web-based UI
* Persistent game history
* Automated tests

---

##  Why This Design

This solution prioritizes:

* Correctness over aesthetics
* Simplicity over over-engineering
* Clear separation of concerns
* Concurrency safety

The design scales naturally with Go’s concurrency model and is easy to reason about during code review or interviews.

---

##  Author

Jagriti Srivastava

---

##  License

This project is provided for educational and evaluation purposes.
