package config

import (
	"Crash-Currency-service/internal/transport/server"
	"Crash-Currency-service/pkg/db"
	"Crash-Currency-service/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(
		LoadConfig,
		func(cfg *Config) *db.DBConfig { return cfg.DBConfig },
		func(cfg *Config) *server.GrpcConfig { return cfg.GrpcConfig },
		func(cfg *Config) *logger.Config { return cfg.LoggerConfig },
	),
)
