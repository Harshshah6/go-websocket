package app

import (
	"log"
	"net/http"

	"github.com/harshshah6/go-websocket/internal/handlers"
	"github.com/harshshah6/go-websocket/pkg/logger"
)

// App struct contains all top-level dependencies
type App struct {
	Handler *handlers.ChatHandler
	Logger  *logger.Logger
}

// NewApp initializes the application
func NewApp() *App {
	appLogger := logger.New("[CHAT]")

	handler := handlers.NewChatHandler()

	return &App{
		Handler: handler,
		Logger:  appLogger,
	}
}

// Run starts the HTTP/WebSocket server
func (a *App) Run(addr string) {
	go a.Handler.HandleMessages()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", a.Handler.HandleConnections)

	a.Logger.Println("Server starting on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
