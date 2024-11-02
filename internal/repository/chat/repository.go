package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	modelRepoChat "github.com/Paul1k96/microservices_course_chat_service/internal/repository/chat/model"
	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/db"
)

const (
	chatTableName = "chats"

	chatIDColumn    = "id"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

const (
	chatUsersTableName = "chat_users"

	chatUsersID  = "chat_id"
	userIDColumn = "user_id"
)

// Repository represents chat repository.
type Repository struct {
	db db.DB
}

// NewRepository creates a new instance of repository.ChatRepository.
func NewRepository(pg db.DB) *Repository {
	return &Repository{db: pg}
}

// Create chat.
func (r *Repository) Create(ctx context.Context) (model.ChatID, error) {
	queryBuilder := sq.Insert(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn).
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING id")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("build query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var chatID model.ChatID
	err = r.db.QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	return chatID, nil
}

// AddUsers to chat.
func (r *Repository) AddUsers(ctx context.Context, userChat []*modelRepoChat.User) error {
	queryBuilder := sq.Insert(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatUsersID, userIDColumn)

	for _, user := range userChat {
		queryBuilder = queryBuilder.Values(user.ChatID, user.ID)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repository.AddUsers",
		QueryRaw: query,
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

// Delete chat by id.
func (r *Repository) Delete(ctx context.Context, id int64) error {
	queryBuilder := sq.Delete(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: id})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
