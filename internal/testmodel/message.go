package testmodel

import (
	"time"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/brianvoe/gofakeit/v7"
)

// NewMessage creates a new Message instance
func NewMessage() *model.Message {
	m := struct {
		ID        model.MessageID
		UserID    model.UserID
		ChatID    model.ChatID
		Text      string
		CreatedAt time.Time
	}{}

	_ = gofakeit.Struct(&m)
	res := model.Message(m)
	return &res
}
