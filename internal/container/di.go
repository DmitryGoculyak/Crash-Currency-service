package container

import (
	"Crash-Currency-service/internal/config"
	repo "Crash-Currency-service/internal/repository/pgsql"
	"Crash-Currency-service/internal/service"
	"Crash-Currency-service/internal/transport/rpc/handlers"
	"Crash-Currency-service/internal/transport/server"
	"Crash-Currency-service/pkg/db"
	"Crash-Currency-service/pkg/logger"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		db.Module,
		config.Module,
		logger.Module,
		repo.Module,
		service.Module,
		handlers.Module,
		server.Module,
	)
}
