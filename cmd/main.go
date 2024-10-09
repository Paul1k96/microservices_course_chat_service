package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

// Chat represents chat model.
type Chat struct {
	ID       int64
	Users    []string
	Messages []Message
}

// Message represents chat message.
type Message struct {
	ID        uuid.UUID
	User      string
	Text      string
	CreatedAt time.Time
}

// ChatMap represents chat repository.
type ChatMap struct {
	chats map[int64]Chat
	mu    sync.RWMutex
}

// NewChatMap creates a new chat repository.
func NewChatMap() *ChatMap {
	return &ChatMap{chats: make(map[int64]Chat)}
}

// Create creates a chat.
func (c *ChatMap) Create(_ context.Context, chat Chat) (int64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	chat.ID = int64(len(c.chats) + 1)
	c.chats[chat.ID] = chat

	return chat.ID, nil
}

// SendMessage sends a message to a chat.
func (c *ChatMap) SendMessage(_ context.Context, chatID int64, message Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	chat, ok := c.chats[chatID]
	if !ok {
		return fmt.Errorf("chat with id %d not found", chatID)
	}

	chat.Messages = append(chat.Messages, message)
	c.chats[chatID] = chat

	return nil
}

// Delete deletes a chat.
func (c *ChatMap) Delete(_ context.Context, chatID int64) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.chats, chatID)

	return nil
}

// ChatRepository represents chat repository.
type ChatRepository interface {
	Create(ctx context.Context, chat Chat) (int64, error)
	SendMessage(ctx context.Context, chatID int64, message Message) error
	Delete(ctx context.Context, chatID int64) error
}

// ChatAPI represents chat API.
type ChatAPI struct {
	logger   *slog.Logger
	chatRepo ChatRepository
	chat_v1.UnimplementedChatServer
}

// NewChatAPI creates a new chat API.
func NewChatAPI(logger *slog.Logger, chatRepo ChatRepository) *ChatAPI {
	return &ChatAPI{logger: logger, chatRepo: chatRepo}
}

// Create creates a chat with the given usernames.
func (c *ChatAPI) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	logger := c.logger.
		With("method", "Create").
		With("usernames", req.Usernames)

	chat := Chat{
		ID:    1,
		Users: req.Usernames,
	}

	chatID, err := c.chatRepo.Create(ctx, chat)
	if err != nil {
		logger.Error("failed to create chat", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return &chat_v1.CreateResponse{Id: chatID}, nil
}

// SendMessage sends a message to a chat.
func (c *ChatAPI) SendMessage(
	ctx context.Context,
	req *chat_v1.SendMessageRequest,
) (*chat_v1.SendMessageResponse, error) {
	logger := c.logger.With("method", "SendMessage")

	message := Message{
		ID:        uuid.New(),
		User:      req.From,
		Text:      req.Text,
		CreatedAt: req.Timestamp.AsTime(),
	}

	err := c.chatRepo.SendMessage(ctx, req.ChatId, message)
	if err != nil {
		logger.Error("failed to send message", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	return &chat_v1.SendMessageResponse{}, nil
}

// Delete deletes a chat.
func (c *ChatAPI) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*chat_v1.DeleteResponse, error) {
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

func main() {
	logger := slog.Default()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		logger.Error("failed to listen", slog.String("error", err.Error()))
		return
	}

	chatDB := NewChatMap()
	chatAPIv1 := NewChatAPI(logger, chatDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat_v1.RegisterChatServer(grpcServer, chatAPIv1)

	logger.Info("server listening at", slog.Any("addr", listen.Addr()))

	if err = grpcServer.Serve(listen); err != nil {
		logger.Error("failed to serve", slog.String("error", err.Error()))
		return
	}
}
