package main

import (
	"log/slog"
	"net"

	"github.com/Paul1k96/microservices_course_chat_service/internal/chat"
	"github.com/Paul1k96/microservices_course_chat_service/internal/config/env"
	"github.com/Paul1k96/microservices_course_chat_service/pkg/proto/gen/chat_v1"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := slog.Default()

	grpcConfig := env.NewGRPCConfig()

	listen, err := net.Listen("tcp", grpcConfig.GetAddress())
	if err != nil {
		logger.Error("failed to listen", slog.String("error", err.Error()))
		return
	}

	pgConfig := env.NewPGConfig()

	db, err := sqlx.Connect("postgres", pgConfig.GetDSN())
	if err != nil {
		logger.Error("failed to connect to database", slog.String("error", err.Error()))
		return
	}

	chatDB := chat.NewChatRepository(db)

	chatAPIv1 := chat.NewChatAPI(logger, chatDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat_v1.RegisterChatServer(grpcServer, chatAPIv1)

	logger.Info("server listening at", slog.Any("addr", listen.Addr()))

	if err = grpcServer.Serve(listen); err != nil {
		logger.Error("failed to serve", slog.String("error", err.Error()))
		return
	}
}
