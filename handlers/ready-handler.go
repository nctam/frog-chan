package handlers

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func Ready(ctx context.Context) func(s *discord.Session, r *discord.Ready) {
	logger := zerolog.Ctx(ctx)
	return func(s *discord.Session, r *discord.Ready) {
		logger.Info().Msg("Bot is ready")
	}
}
