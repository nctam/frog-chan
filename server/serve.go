package server

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func ReadConfig(ctx context.Context) (*Config, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Start loading appConfig")
	configData, err := os.ReadFile("./config_files/keys.json")
	if err != nil {
		return nil, err
	}

	configString := os.ExpandEnv(string(configData))

	config := &Config{}
	if err = json.Unmarshal([]byte(configString), config); err != nil {
		return nil, err
	}

	logger.Info().Msg("End loading appConfig")
	return config, nil
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
