package mapper

import (
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	desc "github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
)

// ToMessageFromSendMessageRequest converts message from send message request to message service model.
func ToMessageFromSendMessageRequest(req *desc.SendMessageRequest) *model.Message {
	var message model.Message
	message.ChatID = model.ChatID(req.ChatId)
	message.UserID = model.UserID(req.From)
	message.Text = req.Text

	if req.Timestamp.IsValid() {
		message.CreatedAt = req.Timestamp.AsTime()
	}

	return &message
}
