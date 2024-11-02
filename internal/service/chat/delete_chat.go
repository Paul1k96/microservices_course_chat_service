package chat

import (
	"context"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
)

// Delete deletes chat.
func (s *service) Delete(ctx context.Context, chatID model.ChatID) error {
	err := s.chatRepository.Delete(ctx, int64(chatID))
	if err != nil {
		return err
	}

	return nil
}
