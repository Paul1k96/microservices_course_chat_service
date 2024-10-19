package chat

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const (
	chatTable     = "chats"
	chatUserTable = "chat_users"
	messageTable  = "messages"
)

// Repository represents chat repository.
type Repository struct {
	pg *sqlx.DB
}

// NewChatRepository creates a new instance of Repository.
func NewChatRepository(pg *sqlx.DB) *Repository {
	return &Repository{pg: pg}
}

// Create chat.
func (r *Repository) Create(ctx context.Context) (*int64, error) {
	queryBuilder := sq.Insert(chatTable).
		PlaceholderFormat(sq.Dollar).
		Columns("id").
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING id")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var chatID int64
	err = r.pg.QueryRowContext(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &chatID, nil
}

// AddBatchUsers adds users to a chat.
func (r *Repository) AddBatchUsers(ctx context.Context, chatID int64, users []User) error {
	tx, txErr := r.pg.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if txErr != nil {
		return fmt.Errorf("failed to begin transaction: %w", txErr)
	}

	for _, user := range users {
		queryBuilder := sq.Insert(chatUserTable).
			PlaceholderFormat(sq.Dollar).
			Columns("chat_id", "user_name").
			Values(chatID, user.Name)

		query, args, err := queryBuilder.ToSql()
		if err != nil {
			return fmt.Errorf("failed to build query: %w", err)
		}

		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("failed to exec query: %w", err)
		}
	}

	txErr = tx.Commit()
	if txErr != nil {
		return fmt.Errorf("failed to commit transaction: %w", txErr)
	}
	return nil
}

// SendMessage sends a message to a chat.
func (r *Repository) SendMessage(ctx context.Context, chatID int64, message Message) error {
	queryBuilder := sq.Insert(messageTable).
		PlaceholderFormat(sq.Dollar).
		Columns("id", "chat_id", "user_name", "text").
		Values(message.ID, chatID, message.User, message.Text)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.pg.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}

// Delete chat.
func (r *Repository) Delete(ctx context.Context, chatID int64) error {
	queryBuilder := sq.Delete(chatTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": chatID})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.pg.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}
