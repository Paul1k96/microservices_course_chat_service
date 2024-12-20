package chat

import (
	"github.com/Paul1k96/microservices_course_chat_service/internal/repository"
	svc "github.com/Paul1k96/microservices_course_chat_service/internal/service"
	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/db"
)

type service struct {
	chatRepository repository.ChatRepository
	userRepo       repository.UserRepository
	messageRepo    repository.MessageRepository
	txManager      db.TxManager
}

// NewService creates a new service.
func NewService(
	chatRepository repository.ChatRepository,
	userRepo repository.UserRepository,
	messageRepo repository.MessageRepository,
	txManager db.TxManager,
) svc.ChatService {
	return &service{
		chatRepository: chatRepository,
		userRepo:       userRepo,
		messageRepo:    messageRepo,
		txManager:      txManager,
	}
}
