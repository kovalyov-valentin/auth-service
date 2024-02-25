package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	auth "sso/internal/services/auth"
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
		panic(err)
	}
	// TODO: инициализировать auth service (auth)
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
