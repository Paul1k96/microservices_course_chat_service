package chat

import (
	"time"

	"github.com/google/uuid"
)

// Chat represents chat model.
type Chat struct {
	ID       int64
	Users    []User
	Messages []Message
}

// Message represents chat message.
type Message struct {
	ID        uuid.UUID
	User      string
	Text      string
	CreatedAt time.Time
}

// User represents user model.
type User struct {
	Name string
}
