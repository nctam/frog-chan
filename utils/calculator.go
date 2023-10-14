package utils

import (
	"context"
	"github.com/rs/zerolog"
	"golang.org/x/exp/rand"
	"time"
)

func ShouldReply(ctx context.Context, prob float64, universalSet int) bool {
	logger := zerolog.Ctx(ctx)
	counter := 0
	for i := 0; i < universalSet; i++ {
		rand.Seed(uint64(time.Now().UnixNano()))
		randValue := rand.Int()
		if randValue%2 == 0 {
			counter += 1
		}
	}
	logger.Info().Msgf("Total votes: %d", counter)

	decision := float64(counter)/float64(universalSet) < prob
	logger.Info().Msgf("Reply decision: %v", decision)

	return decision
}
