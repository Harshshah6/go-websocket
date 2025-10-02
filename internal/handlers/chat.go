package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/harshshah6/go-websocket/internal/models"
	"github.com/harshshah6/go-websocket/internal/services"
	"github.com/harshshah6/go-websocket/pkg/logger"
)

type ChatHandler struct {
	Clients   map[*websocket.Conn]string
	Broadcast chan models.Message
	Upgrader  websocket.Upgrader
	Service   *services.ChatService
	Logger    *logger.Logger
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		Clients:   make(map[*websocket.Conn]string),
		Broadcast: make(chan models.Message),
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		Service: &services.ChatService{},
		Logger: logger.New("[CHAT]"),
	}
}

func (h *ChatHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Logger.Fatal("WebSocket upgrade error:", err)
		return
	}
	defer ws.Close()

	// First message is the user's name
	_, nameMsg, err := ws.ReadMessage()
	if err != nil {
		return
	}
	name := string(nameMsg)
	h.Clients[ws] = name

	h.Broadcast <- h.Service.NewMessage(name, "joined the chat 👋", true)

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			h.Broadcast <- h.Service.NewMessage(h.Clients[ws], "left the chat 👋", true)
			delete(h.Clients, ws)
			break
		}
		msg.Name = h.Clients[ws]
		msg.System = false
		h.Broadcast <- h.Service.NewMessage(msg.Name, msg.Message, false)
	}
}

func (h *ChatHandler) HandleMessages() {
	for {
		msg := <-h.Broadcast
		data, _ := json.Marshal(msg)
		for client := range h.Clients {
			err := client.WriteMessage(websocket.TextMessage, data)
			h.Logger.Logger.Println(string(data))
			if err != nil {
				client.Close()
				delete(h.Clients, client)
			}
		}
	}
}
