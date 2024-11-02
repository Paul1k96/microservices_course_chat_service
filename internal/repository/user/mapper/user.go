package mapper

import (
	modelRepoUser "github.com/Paul1k96/microservices_course_auth/pkg/proto/gen/user_v1"
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
)

// ToUserFromGetResponse converts modelRepoUser.GetResponse to User.
func ToUserFromGetResponse(resp *modelRepoUser.GetResponse) *model.User {
	return &model.User{
		ID:   model.UserID(resp.Id),
		Name: resp.Name,
	}
}

// ToGetRequest converts UserID to modelRepoUser.GetRequest.
func ToGetRequest(id model.UserID) *modelRepoUser.GetRequest {
	return &modelRepoUser.GetRequest{
		Id: int64(id),
	}
}
