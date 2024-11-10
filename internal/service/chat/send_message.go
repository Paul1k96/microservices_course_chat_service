package chat

import (
	"context"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/google/uuid"
)

// SendMessage sends message.
func (s *service) SendMessage(ctx context.Context, message *model.Message) error {
	message.ID = model.MessageID(uuid.New())
	return s.messageRepo.Create(ctx, message)
}
