package v1

import (
	"log/slog"

	"github.com/Paul1k96/microservices_course_chat_service/internal/service"
	desc "github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
)

// Implementation of the chat service.
type Implementation struct {
	logger      *slog.Logger
	chatService service.ChatService
	desc.UnimplementedChatServer
}

// NewImplementation creates a new chat service implementation.
func NewImplementation(logger *slog.Logger, chatService service.ChatService) *Implementation {
	return &Implementation{logger: logger, chatService: chatService}
}
