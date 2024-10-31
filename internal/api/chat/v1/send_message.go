package v1

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Paul1k96/microservices_course_chat_service/internal/mapper"
	desc "github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
)

// SendMessage creates a chat.
func (i *Implementation) SendMessage(
	ctx context.Context,
	request *desc.SendMessageRequest,
) (*desc.SendMessageResponse, error) {
	logger := i.logger.
		With("method", "Create").
		With("userID", request.From).
		With("chatID", request.ChatId)

	err := i.chatService.SendMessage(ctx, mapper.ToMessageFromSendMessageRequest(request))
	if err != nil {
		logger.Error("failed to create chat", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return &desc.SendMessageResponse{}, nil
}
