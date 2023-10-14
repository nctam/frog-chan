package server

import (
	"context"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

func ReadConfig(ctx context.Context) *Config {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Start loading appConfig")
	configData, err := os.ReadFile("./config_files/app.yaml")
	if err != nil {
		logger.Warn().Msg("Failed to read config file")
		return &Config{}
	}

	configString := os.ExpandEnv(string(configData))

	config := &Config{}

	if err = yaml.Unmarshal([]byte(configString), config); err != nil {
		logger.Warn().Msg("Failed to resolve env from config file")
		return &Config{}
	}

	logger.Info().Msg("End loading appConfig")
	return config
}

func ConfigLogger(ctx context.Context) context.Context {
	loggerOutPut := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	logger := zerolog.New(loggerOutPut).With().Timestamp().Logger()
	ctx = logger.WithContext(ctx)

	return ctx
}
