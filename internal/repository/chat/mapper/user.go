package mapper

import (
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	modelRepo "github.com/Paul1k96/microservices_course_chat_service/internal/repository/chat/model"
)

// ToUserFromRepo converts user from repository to user service model.
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:     model.UserID(user.ID),
		ChatID: model.ChatID(user.ChatID),
	}
}

// ToRepoCreateFromUsersService converts users from service to repository model.
func ToRepoCreateFromUsersService(user model.Users) []*modelRepo.User {
	users := make([]*modelRepo.User, 0, len(user))
	for _, u := range user {
		users = append(users, &modelRepo.User{
			ID:     int64(u.ID),
			ChatID: int64(u.ChatID),
		})
	}
	return users
}
