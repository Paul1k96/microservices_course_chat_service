package repository

import (
	"context"

	modelRepoUser "github.com/Paul1k96/microservices_course_auth/pkg/proto/gen/user_v1"
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	modelRepoChat "github.com/Paul1k96/microservices_course_chat_service/internal/repository/chat/model"
	modelRepoMessage "github.com/Paul1k96/microservices_course_chat_service/internal/repository/message/model"
)

// ChatRepository represents chat repository.
type ChatRepository interface {
	Create(ctx context.Context) (model.ChatID, error)
	AddUsers(ctx context.Context, userChat []*modelRepoChat.User) error
	Delete(ctx context.Context, chatID int64) error
}

// MessageRepository represents message repository.
type MessageRepository interface {
	Create(ctx context.Context, message *modelRepoMessage.Message) error
}

// UserRepository represents user repository.
type UserRepository interface {
	Get(ctx context.Context, request *modelRepoUser.GetRequest) (*model.User, error)
}
