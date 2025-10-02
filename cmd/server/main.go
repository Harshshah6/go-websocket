package main

import (
	"github.com/harshshah6/go-websocket/internal/app"
)

func main() {
	application := app.NewApp()
	application.Run(":8080")
}
