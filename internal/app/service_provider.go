package app

import (
	"context"
	"fmt"
	"log/slog"

	userV1Client "github.com/Paul1k96/microservices_course_auth/pkg/proto/gen/user_v1"
	chatv1 "github.com/Paul1k96/microservices_course_chat_service/internal/api/chat/v1"
	"github.com/Paul1k96/microservices_course_chat_service/internal/client/db"
	"github.com/Paul1k96/microservices_course_chat_service/internal/client/db/pg"
	"github.com/Paul1k96/microservices_course_chat_service/internal/client/db/transaction"
	"github.com/Paul1k96/microservices_course_chat_service/internal/closer"
	"github.com/Paul1k96/microservices_course_chat_service/internal/config"
	"github.com/Paul1k96/microservices_course_chat_service/internal/config/env"
	"github.com/Paul1k96/microservices_course_chat_service/internal/repository"
	chatRepo "github.com/Paul1k96/microservices_course_chat_service/internal/repository/chat"
	messageRepo "github.com/Paul1k96/microservices_course_chat_service/internal/repository/message"
	"github.com/Paul1k96/microservices_course_chat_service/internal/repository/user"
	"github.com/Paul1k96/microservices_course_chat_service/internal/service"
	chatSvc "github.com/Paul1k96/microservices_course_chat_service/internal/service/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	grpcServerConfig config.GRPCConfig
	logger           *slog.Logger

	dbClient          db.Client
	txManager         db.TxManager
	chatRepository    repository.ChatRepository
	userRepository    repository.UserRepository
	messageRepository repository.MessageRepository

	chatService service.ChatService

	chatV1Impl *chatv1.Implementation
	userClient userV1Client.UserClient
}

func newServiceProvider(logger *slog.Logger) *serviceProvider {
	return &serviceProvider{logger: logger}
}

// PGConfig returns an instance of config.PGConfig.
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		s.pgConfig = env.NewPGConfig()
	}

	return s.pgConfig
}

// GRPCConfig returns an instance of config.GRPCConfig.
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcServerConfig == nil {
		s.grpcServerConfig = env.NewGRPCConfig()
	}

	return s.grpcServerConfig
}

// DBClient returns an instance of db.Client.
func (s *serviceProvider) DBClient(ctx context.Context) (db.Client, error) {
	if s.dbClient == nil {
		s.logger.Info(
			"creating db client, dsn:",
			slog.String("dsn", s.PGConfig().GetDSN()),
		)

		cl, err := pg.New(ctx, s.PGConfig().GetDSN())
		if err != nil {
			return nil, fmt.Errorf("failed to create db client: %w", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			return nil, fmt.Errorf("ping error: %w", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient, nil
}

// UserClient returns an instance of userClient.UserClient.
func (s *serviceProvider) UserClient() (userV1Client.UserClient, error) {
	if s.userClient == nil {
		s.logger.Info(
			"creating GRPC auth client, address:",
			slog.String("address", s.GRPCConfig().GetAuthAddress()))

		conn, err := grpc.NewClient(
			s.GRPCConfig().GetAuthAddress(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create grpc client: %w", err)
		}

		s.userClient = userV1Client.NewUserClient(conn)
	}

	return s.userClient, nil
}

// TxManager returns an instance of db.TxManager.
func (s *serviceProvider) TxManager(ctx context.Context) (db.TxManager, error) {
	if s.txManager == nil {
		dbClient, err := s.DBClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get db client: %w", err)
		}
		s.txManager = transaction.NewTransactionManager(dbClient.DB())
	}

	return s.txManager, nil
}

// ChatRepository returns an instance of repository.ChatRepository.
func (s *serviceProvider) ChatRepository(ctx context.Context) (repository.ChatRepository, error) {
	if s.chatRepository == nil {
		dbClient, err := s.DBClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get db client: %w", err)
		}
		s.chatRepository = chatRepo.NewRepository(dbClient.DB())
	}

	return s.chatRepository, nil
}

// MessageRepository returns an instance of repository.MessageRepository.
func (s *serviceProvider) MessageRepository(ctx context.Context) (repository.MessageRepository, error) {
	if s.messageRepository == nil {
		dbClient, err := s.DBClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get db client: %w", err)
		}
		s.messageRepository = messageRepo.NewRepository(dbClient.DB())
	}

	return s.messageRepository, nil
}

// UserRepository returns an instance of repository.UserRepository.
func (s *serviceProvider) UserRepository(_ context.Context) (repository.UserRepository, error) {
	if s.userRepository == nil {
		userClient, err := s.UserClient()
		if err != nil {
			return nil, fmt.Errorf("failed to get user client: %w", err)
		}
		s.userRepository = user.NewRepository(userClient)
	}

	return s.userRepository, nil
}

// ChatService returns an instance of service.ChatService.
func (s *serviceProvider) ChatService(ctx context.Context) (service.ChatService, error) {
	if s.chatService == nil {
		txManager, err := s.TxManager(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get tx manager: %w", err)
		}
		chatRepository, err := s.ChatRepository(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get chat repository: %w", err)
		}
		messageRepository, err := s.MessageRepository(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get message repository: %w", err)
		}
		userRepository, err := s.UserRepository(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get user repository: %w", err)
		}
		s.chatService = chatSvc.NewService(chatRepository, userRepository, messageRepository, txManager)
	}

	return s.chatService, nil
}

// ChatV1Impl returns an instance of chatV1.Implementation.
func (s *serviceProvider) ChatV1Impl(ctx context.Context) (*chatv1.Implementation, error) {
	if s.chatV1Impl == nil {
		chatService, err := s.ChatService(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get chat service: %w", err)
		}
		s.chatV1Impl = chatv1.NewImplementation(s.logger, chatService)
	}

	return s.chatV1Impl, nil
}
