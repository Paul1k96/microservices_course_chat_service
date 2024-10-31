package model

// User represents user model.
type User struct {
	ID     UserID
	ChatID ChatID
	Name   string
}

// UserID represents user ID.
type UserID int64

// Users represents a list of users.
type Users []*User

// SetChatID sets chat ID to all users.
func (u *Users) SetChatID(chatID ChatID) {
	for _, user := range *u {
		user.ChatID = chatID
	}
}
