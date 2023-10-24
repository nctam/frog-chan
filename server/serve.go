package server

import (
	"context"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

var (
	AppConfig *Config
)

func init() {
	AppConfig = readConfig()
}

func readConfig() *Config {
	ctx := context.TODO()
	ctx = ConfigLogger().WithContext(ctx)
	logger := zerolog.Ctx(ctx).With().Str("Server", "ReadConfig").Logger()
	logger.Debug().Msg("Start loading appConfig")
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

	if config.Env == "test" {
		config.Probability = 0.3
		config.UniversalSet = 100
		config.Channels = config.Channels[:1] // use test channel only
		config.ExcludedUsers = make([]string, 0)
	}

	logger.Debug().Msg("End loading appConfig")
	return config
}

func ConfigLogger() zerolog.Logger {
	loggerOutPut := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	return zerolog.New(loggerOutPut).With().Timestamp().Logger()
}
