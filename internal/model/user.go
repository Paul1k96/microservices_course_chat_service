package model

// User represents user model.
type User struct {
	ID     UserID
	ChatID ChatID
	Name   string
}

// UserID represents user ID.
type UserID int64

// ToInt64 converts user ID to int64.
func (id UserID) ToInt64() int64 {
	return int64(id)
}

// UserIDs represents a list of user IDs.
type UserIDs []UserID

// ToInt64 converts user IDs to int64.
func (ids UserIDs) ToInt64() []int64 {
	intIDs := make([]int64, 0, len(ids))
	for _, id := range ids {
		intIDs = append(intIDs, id.ToInt64())
	}

	return intIDs
}

// Users represents a list of users.
type Users []*User

// SetChatID sets chat ID to all users.
func (u *Users) SetChatID(chatID ChatID) {
	for _, user := range *u {
		user.ChatID = chatID
	}
}
