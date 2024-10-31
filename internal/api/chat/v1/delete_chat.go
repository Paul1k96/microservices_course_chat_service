package v1

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	desc "github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
)

// Delete deletes a chat.
func (i *Implementation) Delete(ctx context.Context, request *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	logger := i.logger.
		With("method", "Delete").
		With("chatID", request.Id)

	err := i.chatService.Delete(ctx, model.ChatID(request.Id))
	if err != nil {
		logger.Error("failed to create chat", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return &desc.DeleteResponse{}, nil
}
