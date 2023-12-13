package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/exp/rand"

	"kaeru.chan/voz/server"
)

const (
	logTag = "Utils"
)

func makeDecision(probability float64, universal int) (int, bool) {
	counter := 0
	for i := 0; i < universal; i++ {
		rand.Seed(uint64(time.Now().UnixNano()))
		randValue := rand.Int()
		fmt.Println(randValue)
		if randValue%5 == 0 {
			counter += 1
		}
	}
	decision := float64(counter)/float64(universal) >= probability
	return counter, decision
}

func ShouldReply(ctx context.Context, userID string, config *server.Config) bool {
	logger := zerolog.Ctx(ctx).With().Str(logTag, "ShouldReply").Logger()
	if StringContains(config.PriorUsers, userID) {
		logger.Info().Msgf("Prior user: %s found, start reply", userID)
		return true
	}
	counter, decision := makeDecision(config.Probability, config.UniversalSet)
	logger.Info().Msgf("Total votes: %d ==> Reply decision: %v", counter, decision)
	return config.Env == "test" || decision
}

func ShouldRename(ctx context.Context, config *server.Config) bool {
	logger := zerolog.Ctx(ctx).With().Str(logTag, "ShouldRename").Logger()
	counter, decision := makeDecision(config.Probability, config.UniversalSet)
	logger.Info().Msgf("Total votes: %d ==> Rename decision: %v", counter, decision)
	return config.Env == "test" || decision
}
