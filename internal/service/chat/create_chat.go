package chat

import (
	"context"
	"fmt"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
)

// Create creates a chat with users.
func (s *service) Create(ctx context.Context, userIDs model.UserIDs) (model.ChatID, error) {
	var chatID model.ChatID

	if txErr := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {

		users, err := s.getUsers(ctx, userIDs)
		if err != nil {
			return fmt.Errorf("get users: %w", err)
		}

		chatID, err = s.chatRepository.Create(ctx)
		if err != nil {
			return fmt.Errorf("create chat: %w", err)
		}

		users.SetChatID(chatID)

		err = s.chatRepository.AddUsers(ctx, users)
		if err != nil {
			return fmt.Errorf("add users: %w", err)
		}

		return nil
	}); txErr != nil {
		return 0, fmt.Errorf("transaction error: %w", txErr)
	}

	return chatID, nil
}

func (s *service) getUsers(ctx context.Context, userIDs model.UserIDs) (model.Users, error) {
	users, err := s.userRepo.List(ctx, userIDs)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	return users, nil
}
