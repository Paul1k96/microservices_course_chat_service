package v1

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Paul1k96/microservices_course_chat_service/internal/mapper"
	desc "github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
)

// Create creates a chat.
func (i *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	logger := i.logger.
		With("method", "Create").
		With("userIDs", request.UserIds)

	chatID, err := i.chatService.Create(ctx, mapper.ToUserIDsFromCreateRequest(request))
	if err != nil {
		logger.Error("failed to create chat", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return mapper.ToCreateResponseFromChat(chatID), nil
}
