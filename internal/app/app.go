package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/Paul1k96/microservices_course_chat_service/internal/closer"
	chatv1 "github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// App is the main application structure.
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	logger          *slog.Logger
}

// NewApp creates a new instance of the App.
func NewApp(ctx context.Context, logger *slog.Logger) (*App, error) {
	a := &App{
		logger: logger,
	}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to init dependencies: %w", err)
	}

	return a, nil
}

// Run runs the application.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.logger)
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	chatImplV1, err := a.serviceProvider.ChatV1Impl(ctx)
	if err != nil {
		return fmt.Errorf("failed to create chat v1 implementation: %w", err)
	}

	reflection.Register(a.grpcServer)
	chatv1.RegisterChatServer(a.grpcServer, chatImplV1)

	return nil
}

func (a *App) runGRPCServer() error {
	a.logger.Info(
		"GRPC server is running on",
		slog.Any("addr", a.serviceProvider.GRPCConfig().GetServerAddress()))

	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().GetServerAddress())
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	if err = a.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
