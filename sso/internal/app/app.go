package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	auth2 "sso/internal/services/auth"
	"sso/internal/storage/sqlite"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: инициализировать хранилище
	storage, err := sqlite.New(storagePath)
	if err != nil {
		return nil
	}
	// TODO: инициализировать auth service (auth)
	auth := auth2.New()

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
