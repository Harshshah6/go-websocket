# Go Chat App

A simple **real-time chat application** built with **Go** and [Gorilla WebSocket](https://github.com/gorilla/websocket) for the backend and a lightweight **HTML/JS frontend**.

---

## Features

* Real-time chat using WebSockets
* User joins/leaves are broadcast as system messages
* Supports **bold**, *italic*, and `inline code` markdown
* Preserves **multiline messages** with line breaks
* Auto-scroll to newest messages
* Simple and responsive UI

---

## Project Structure

```
.
├── main.go              # Go server (WebSocket + HTTP handler)
├── Dockerfile           # Docker build file
├── docker-compose.yml   # Compose for running/scaling easily
└── static/
    └── index.html       # Chat frontend 
```

---

## Requirements

* [Go 1.18+](https://go.dev/dl/)
* Gorilla WebSocket package
* Docker & Docker Compose (optional)

Install Gorilla WebSocket:

```bash
go get github.com/gorilla/websocket
```

---

## Running Locally (without Docker)

1. Clone this repository:

    ```bash
    git clone https://github.com/yourusername/go-chat.git
    cd go-chat
    ```

2. Run the server:

    ```bash
    go run main.go
    ```

3. Open the app in your browser:

    ```
    http://localhost:8080
    ```

---

## Running with Docker

1. Build the Docker image:

    ```bash
    docker build -t go-chat .
    ```

2. Run the container:

    ```bash
    docker run -p 8080:8080 go-chat
    ```

3. Open the app:

    ```
    http://localhost:8080
    ```

---

## Running with Docker Compose

1. Start the service:

    ```bash
    docker compose up --build
    ```

2. Open the app:

    ```
    http://localhost:8080
    ```

3. To scale multiple instances (requires sticky sessions in a reverse proxy for WebSockets):

    ```bash
    docker compose up --scale go-chat=3
    ```

---

## Usage

* On page load, enter your **name**.
* Type messages in the input box and press **Enter** to send.
* Use **Shift+Enter** to create a new line.
* Basic Markdown supported:

  * `**bold**` → **bold**
  * `*italic*` → *italic*
  * `` `code` `` → `code`

---

## License

MIT License.
You are free to use and modify this project.
