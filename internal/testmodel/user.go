package testmodel

import (
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/brianvoe/gofakeit/v7"
)

// NewUsers creates a slice of User instances
func NewUsers(quantity int) model.Users {
	res := make([]*model.User, 0, quantity)
	for i := 0; i < quantity; i++ {
		res = append(res, NewUser())
	}

	return res
}

// NewUser creates a new User instance
func NewUser() *model.User {
	m := struct {
		ID     model.UserID
		ChatID model.ChatID
		Name   string
	}{}

	_ = gofakeit.Struct(&m)
	res := model.User(m)
	return &res
}
