package mapper

import (
	modelRepoUser "github.com/Paul1k96/microservices_course_auth/pkg/proto/gen/user_v1"
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
)

// ToUsersFromListResponse converts modelRepoUser.ListResponse to []*model.User.
func ToUsersFromListResponse(resp *modelRepoUser.GetListResponse) []*model.User {
	users := make([]*model.User, 0, len(resp.Users))
	for _, u := range resp.Users {
		users = append(users, ToUserFromGetResponse(u))
	}

	return users
}

// ToUserFromGetResponse converts modelRepoUser.GetResponse to User.
func ToUserFromGetResponse(resp *modelRepoUser.GetResponse) *model.User {
	return &model.User{
		ID:   model.UserID(resp.Id),
		Name: resp.Name,
	}
}
