package testmodel

import (
	"time"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/brianvoe/gofakeit/v7"
)

// NewChat creates a new Chat instance
func NewChat() *model.Chat {
	m := struct {
		ID        model.ChatID
		CreatedAt time.Time
		UpdatedAt *time.Time
	}{}

	_ = gofakeit.Struct(&m)
	res := model.Chat(m)
	return &res
}
