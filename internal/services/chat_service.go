package services

import (
	"time"
	"github.com/harshshah6/go-websocket/internal/models"
)

type ChatService struct{}

func (s *ChatService) NewMessage(name, text string, system bool) models.Message {
	return models.Message{
		Name:      name,
		Message:   text,
		Timestamp: time.Now().Format("3:04:05 PM"),
		System:    system,
	}
}
