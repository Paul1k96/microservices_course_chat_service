package model

import (
	"time"

	"github.com/google/uuid"
)

// Message represents message model.
type Message struct {
	ID        uuid.UUID `json:"id"`
	UserID    int64     `json:"user_id"`
	ChatID    int64     `json:"chat_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
