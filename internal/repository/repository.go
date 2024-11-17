package repository

import (
	"context"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
)

//go:generate ../../bin/mockgen -source $GOFILE -destination "mocks/repository.go" -package "mocks"

// ChatRepository represents chat repository.
type ChatRepository interface {
	Create(ctx context.Context) (model.ChatID, error)
	AddUsers(ctx context.Context, users model.Users) error
	Delete(ctx context.Context, chatID model.ChatID) error
}

// MessageRepository represents message repository.
type MessageRepository interface {
	Create(ctx context.Context, message *model.Message) error
}

// UserRepository represents user repository.
type UserRepository interface {
	Get(ctx context.Context, id model.UserID) (*model.User, error)
	List(ctx context.Context, ids model.UserIDs) (model.Users, error)
}
