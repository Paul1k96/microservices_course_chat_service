package mapper

import (
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	desc "github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
)

// ToCreateResponseFromChat converts chatID to CreateResponse.
func ToCreateResponseFromChat(chatID model.ChatID) *desc.CreateResponse {
	return &desc.CreateResponse{
		Id: int64(chatID),
	}
}

// ToUserIDsFromCreateRequest converts userIDs to UserIDs.
func ToUserIDsFromCreateRequest(req *desc.CreateRequest) []model.UserID {
	result := make([]model.UserID, 0, len(req.UserIds))
	for _, id := range req.UserIds {
		result = append(result, model.UserID(id))
	}

	return result
}
