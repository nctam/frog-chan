package server

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os"
)

func ReadConfig(ctx context.Context) (*Config, error) {
	log.Info().Ctx(ctx).Msg("Start loading appConfig")
	configData, err := os.ReadFile("./config_files/keys.json")
	if err != nil {
		return nil, err
	}

	configString := os.ExpandEnv(string(configData))

	config := &Config{}
	if err = json.Unmarshal([]byte(configString), config); err != nil {
		return nil, err
	}

	log.Info().Ctx(ctx).Msg("End loading appConfig")
	return config, nil
}
