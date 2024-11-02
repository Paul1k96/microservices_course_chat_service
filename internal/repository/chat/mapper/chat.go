package mapper

import (
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	modelRepo "github.com/Paul1k96/microservices_course_chat_service/internal/repository/chat/model"
)

// ToChatFromRepo converts chat from repository to chat service model.
func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	var chatService model.Chat

	chatService.ID = model.ChatID(chat.ID)
	chatService.CreatedAt = chat.CreatedAt

	if chat.UpdatedAt.Valid {
		chatService.UpdatedAt = &chat.UpdatedAt.Time
	}

	return &chatService
}

// ToRepoCreateFromChat converts chat from service to repository model.
func ToRepoCreateFromChat(chat *model.Chat) *modelRepo.Chat {
	var chatRepo modelRepo.Chat

	chatRepo.ID = int64(chat.ID)
	chatRepo.CreatedAt = chat.CreatedAt
	if chat.UpdatedAt != nil {
		chatRepo.UpdatedAt.Time = *chat.UpdatedAt
		chatRepo.UpdatedAt.Valid = true
	}

	return &chatRepo
}
