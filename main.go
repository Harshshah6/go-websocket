package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

type Message struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	System    bool   `json:"system"` // true if it's a system message
}

var (
	clients   = make(map[*websocket.Conn]string)
	broadcast = make(chan Message)

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer ws.Close()

	// First message is the user's name
	_, nameMsg, err := ws.ReadMessage()
	if err != nil {
		log.Println("Failed to read name:", err)
		return
	}
	name := string(nameMsg)
	clients[ws] = name

	log.Printf("User joined: %s\n", name)

	// Broadcast system "user joined" message
	broadcast <- Message{
		Name:      name,
		Message:   "joined the chat 👋",
		Timestamp: time.Now().Format("1:04:05 AM"),
		System:    true,
	}

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("User disconnected: %s\n", clients[ws])

			// Send "user left" system message
			broadcast <- Message{
				Name:      clients[ws],
				Message:   "left the chat 👋",
				Timestamp: time.Now().Format("1:04:05 AM"),
				System:    true,
			}

			delete(clients, ws)
			break
		}

		msg.Name = clients[ws]
		msg.Timestamp = time.Now().Format("1:04:05 AM")
		msg.System = false

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		data, _ := json.Marshal(msg)

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("Write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
