package model

import (
	"database/sql"
	"time"
)

// Chat represents chat model.
type Chat struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
