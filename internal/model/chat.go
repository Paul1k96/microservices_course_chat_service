package model

import "time"

// Chat represents chat model.
type Chat struct {
	ID        ChatID
	CreatedAt time.Time
	UpdatedAt *time.Time
	Users     []*User
}

// ChatID represents chat ID.
type ChatID int64
