package model

import (
	"time"

	"github.com/google/uuid"
)

// Message represents chat message.
type Message struct {
	ID        MessageID
	UserID    UserID
	ChatID    ChatID
	Text      string
	CreatedAt time.Time
}

// MessageID represents message ID.
type MessageID uuid.UUID

// Messages represents a list of messages.
type Messages []Message
