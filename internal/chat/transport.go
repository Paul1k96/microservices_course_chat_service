package chat

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
	"github.com/google/uuid"
)

// ChatsRepository represents chat repository.
type ChatsRepository interface {
	Create(ctx context.Context) (*int64, error)
	AddBatchUsers(ctx context.Context, chatID int64, users []User) error
	SendMessage(ctx context.Context, chatID int64, message Message) error
	Delete(ctx context.Context, chatID int64) error
}

// API represents chat API.
type API struct {
	logger   *slog.Logger
	chatRepo ChatsRepository
	chat_v1.UnimplementedChatServer
}

// NewChatAPI creates a new chat API.
func NewChatAPI(logger *slog.Logger, chatRepo ChatsRepository) *API {
	return &API{logger: logger, chatRepo: chatRepo}
}

// Create creates a chat with the given usernames.
func (c *API) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	logger := c.logger.
		With("method", "Create").
		With("usernames", req.Usernames)

	usersForChat := make([]User, 0, len(req.Usernames))

	for _, username := range req.Usernames {
		usersForChat = append(usersForChat, User{Name: username})
	}

	chatID, err := c.chatRepo.Create(ctx)
	if err != nil {
		logger.Error("failed to create chat", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	err = c.chatRepo.AddBatchUsers(ctx, *chatID, usersForChat)
	if err != nil {
		logger.Error("failed to add users to chat", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to add users to chat: %w", err)
	}

	return &chat_v1.CreateResponse{Id: *chatID}, nil
}

// SendMessage sends a message to a chat.
func (c *API) SendMessage(
	ctx context.Context,
	req *chat_v1.SendMessageRequest,
) (*chat_v1.SendMessageResponse, error) {
	logger := c.logger.With("method", "SendMessage")

	message := Message{
		ID:   uuid.New(),
		User: req.From,
		Text: req.Text,
	}

	err := c.chatRepo.SendMessage(ctx, req.ChatId, message)
	if err != nil {
		logger.Error("failed to send message", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	return &chat_v1.SendMessageResponse{}, nil
}

// Delete deletes a chat.
func (c *API) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*chat_v1.DeleteResponse, error) {
	logger := c.logger.
		With("method", "Delete").
		With("chat_id", req.Id)

	err := c.chatRepo.Delete(ctx, req.Id)
	if err != nil {
		logger.Error("failed to delete chat", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to delete chat: %w", err)
	}

	return &chat_v1.DeleteResponse{}, nil
}
