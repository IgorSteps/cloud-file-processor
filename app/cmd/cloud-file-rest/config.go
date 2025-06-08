package main

import (
	"app/internal/adapters/rest/middleware"
	"app/internal/drivers/restserver"

	"github.com/spf13/viper"
)

type Config struct {
	App     AppConfig
	Server  restserver.ServerConfig
	Tracing middleware.TracingMiddlewareConfig
}

// LoadConfig reads configuration from ./config.yaml file.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cmd/cloud-file-rest")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
