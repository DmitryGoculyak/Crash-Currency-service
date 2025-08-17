package config

import (
	"Crash-Currency-service/internal/transport/server"
	"Crash-Currency-service/pkg/db"
	"Crash-Currency-service/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	err    error
	config *Config
	s      sync.Once
)

type Config struct {
	DBConfig     *db.DBConfig
	GrpcConfig   *server.GrpcConfig
	LoggerConfig *logger.Config
}

func LoadConfig() (*Config, error) {
	s.Do(func() {
		config = &Config{}

		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		DBConfig := viper.Sub("database")
		GrpcConfig := viper.Sub("server")
		LoggerConfig := viper.Sub("logger")

		if err = parseSubConfig(DBConfig, &config.DBConfig); err != nil {
			return
		}
		if err = parseSubConfig(GrpcConfig, &config.GrpcConfig); err != nil {
			return
		}
		if err = parseSubConfig(LoggerConfig, &config.LoggerConfig); err != nil {
			return
		}
	})
	return config, err
}

func parseSubConfig[T any](subConfig *viper.Viper, parseTo *T) error {
	if subConfig == nil {
		return fmt.Errorf("can not read %T config: subconfig is nil", parseTo)
	}

	if err = subConfig.Unmarshal(parseTo); err != nil {
		return err
	}
	return nil
}
