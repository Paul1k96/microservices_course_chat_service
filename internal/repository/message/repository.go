package message

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	modelRepoMessage "github.com/Paul1k96/microservices_course_chat_service/internal/repository/message/model"
	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/db"
)

const (
	messageTable = "messages"

	idColumn        = "id"
	chatIDColumn    = "chat_id"
	userIDColumn    = "user_id"
	textColumn      = "text"
	createdAtColumn = "created_at"
)

// Repository represents message repository.
type Repository struct {
	db db.DB
}

// NewRepository creates a new instance of repository.MessageRepository.
func NewRepository(pg db.DB) *Repository {
	return &Repository{db: pg}
}

// Create message.
func (r *Repository) Create(ctx context.Context, message *modelRepoMessage.Message) error {
	queryBuilder := sq.Insert(messageTable).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, chatIDColumn, userIDColumn, textColumn, createdAtColumn).
		Values(message.ID, message.ChatID, message.UserID, message.Text, message.CreatedAt)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	q := db.Query{
		Name:     "message_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
