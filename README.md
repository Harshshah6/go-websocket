# Go Chat (Structured Go Project)

A real-time **chat application** built with **Go** and [Gorilla WebSocket](https://github.com/gorilla/websocket).
This version is structured for **scalability** and **maintainability** using a layered architecture.

---

## Features

* Real-time chat with WebSockets
* User join/leave notifications
* Markdown support (`**bold**`, `*italic*`, `` `code` ``)
* Multiline messages (Shift+Enter for new line)
* Clean project structure (`internal/`, `pkg/`, `cmd/`)
* Custom logger

---

## Project Structure

```
go-chat/
├── cmd/
│   └── server/
│       └── main.go       # Entry point
│
├── internal/
│   ├── app/
│   │   └── app.go        # Wire dependencies + run server
│   ├── config/           # (future) env/config loader
│   ├── handlers/
│   │   └── chat.go       # WebSocket handler
│   ├── models/
│   │   └── message.go    # Data models
│   ├── services/
│   │   └── chat_service.go
│   └── store/
│       └── memory_store.go
│
├── pkg/
│   └── logger/
│       └── logger.go     # Custom logger wrapper
│
├── static/
│   └── index.html        # Frontend chat UI
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## Requirements

* Go 1.18+
* Gorilla WebSocket package

Install dependencies:

```bash
go mod tidy
```

---

## Running Locally

Start the app:

```bash
go run ./cmd/server
```

Visit in your browser:

```
http://localhost:8080
```

---

## Running with Docker

1. Build the image:

```bash
docker build -t go-chat .
```

2. Run the container:

```bash
docker run -p 8080:8080 go-chat
```

Open in browser:

```
http://localhost:8080
```

---

## Running with Docker Compose

```bash
docker compose up --build
```

To scale:

```bash
docker compose up --scale go-chat=3
```

*(requires sticky sessions for WebSockets when load balancing)*

---

## License

MIT License.
Free to use, modify, and distribute.
