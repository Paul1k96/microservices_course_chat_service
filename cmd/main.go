package main

import (
	"context"
	"log/slog"

	"github.com/Paul1k96/microservices_course_chat_service/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	logger := slog.Default()

	ctx := context.Background()
	a, err := app.NewApp(ctx, logger)
	if err != nil {
		logger.Error("failed to init app", slog.String("error", err.Error()))
		return
	}

	if err = a.Run(); err != nil {
		logger.Error("failed to run app", slog.String("error", err.Error()))
		return
	}
}
