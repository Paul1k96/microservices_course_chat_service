package mapper

import (
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	modelRepo "github.com/Paul1k96/microservices_course_chat_service/internal/repository/message/model"
	"github.com/google/uuid"
)

// ToMessageFromRepo converts message from repository to message service model.
func ToMessageFromRepo(message *modelRepo.Message) *model.Message {
	return &model.Message{
		ID:        model.MessageID(message.ID),
		UserID:    model.UserID(message.UserID),
		ChatID:    model.ChatID(message.ChatID),
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}
}

// ToRepoCreateFromMessage converts message from service to repository model.
func ToRepoCreateFromMessage(message *model.Message) *modelRepo.Message {
	return &modelRepo.Message{
		ID:        uuid.UUID(message.ID),
		UserID:    int64(message.UserID),
		ChatID:    int64(message.ChatID),
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}
}
