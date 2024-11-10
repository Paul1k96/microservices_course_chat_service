package service

import (
	"context"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
)

// ChatService is a service for chat.
type ChatService interface {
	Create(ctx context.Context, userIDs model.UserIDs) (model.ChatID, error)
	SendMessage(ctx context.Context, message *model.Message) error
	Delete(ctx context.Context, chatID model.ChatID) error
}
