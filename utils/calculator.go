package utils

import (
    "context"
    "github.com/rs/zerolog"
    "golang.org/x/exp/rand"
    "time"

    "kaeru.chan/voz/server"
)

const (
    logTag = "Utils"
)

func ShouldReply(ctx context.Context, config *server.Config) bool {
    logger := zerolog.Ctx(ctx).With().Str(logTag, "ShouldReply").Logger()
    counter := 0
    for i := 0; i < config.UniversalSet; i++ {
        rand.Seed(uint64(time.Now().UnixNano()))
        randValue := rand.Int()
        if randValue%5 == 0 {
            counter += 1
        }
    }
    logger.Info().Msgf("Total votes: %d", counter)

    decision := float64(counter)/float64(config.UniversalSet) >= config.Probability
    logger.Info().Msgf("Reply decision: %v", decision)

    return config.Env == "test" || decision
}
