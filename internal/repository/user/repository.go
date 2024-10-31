package user

import (
	"context"
	"fmt"

	userRepo "github.com/Paul1k96/microservices_course_auth/pkg/proto/gen/user_v1"
	"github.com/Paul1k96/microservices_course_chat_service/internal/model"
	"github.com/Paul1k96/microservices_course_chat_service/internal/repository/user/mapper"
)

// Repository represents user repository.
type Repository struct {
	grpcClient userRepo.UserClient
}

// NewRepository creates a new repository.
func NewRepository(grpcClient userRepo.UserClient) *Repository {
	return &Repository{grpcClient: grpcClient}
}

// Get gets user by id.
func (r *Repository) Get(ctx context.Context, request *userRepo.GetRequest) (*model.User, error) {
	resp, err := r.grpcClient.Get(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	return mapper.ToUserFromGetResponse(resp), nil
}
