package handlers

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

const (
	logTag = "ReadyHandler"
)

func Ready(ctx context.Context) func(s *discord.Session, r *discord.Ready) {
	logger := zerolog.Ctx(ctx).With().Str(logTag, "Ready").Logger()
	return func(s *discord.Session, r *discord.Ready) {
		logger.Info().Msg("Bot is ready")
	}
}
