package chat

import (
	"context"
	"fmt"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/Paul1k96/microservices_course_chat_service/internal/repository/chat/mapper"
	repoUserMapper "github.com/Paul1k96/microservices_course_chat_service/internal/repository/user/mapper"
)

// Create creates a chat with users.
func (s *service) Create(ctx context.Context, userIDs []model.UserID) (model.ChatID, error) {
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

		err = s.chatRepository.AddUsers(ctx, mapper.ToRepoCreateFromUsersService(users))
		if err != nil {
			return fmt.Errorf("add users: %w", err)
		}

		return nil
	}); txErr != nil {
		return 0, fmt.Errorf("transaction error: %w", txErr)
	}

	return chatID, nil
}

func (s *service) getUsers(ctx context.Context, userIDs []model.UserID) (model.Users, error) {
	users := make(model.Users, 0, len(userIDs))
	for _, userID := range userIDs {
		user, err := s.userRepo.Get(ctx, repoUserMapper.ToGetRequest(userID))
		if err != nil {
			return nil, fmt.Errorf("get user %d: %w", userID, err)
		}

		users = append(users, user)
	}

	return users, nil
}
