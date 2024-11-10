package model

import "time"

// Chat represents chat model.
type Chat struct {
	ID        ChatID
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// ChatID represents chat ID.
type ChatID int64

// ToInt64 converts ChatID to int64.
func (id ChatID) ToInt64() int64 {
	return int64(id)
}
