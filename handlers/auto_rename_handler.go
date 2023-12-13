package handlers

import (
	"context"
	"regexp"

	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"kaeru.chan/voz/logic"
	"kaeru.chan/voz/server"
)

var (
	autoRename = logic.GeneralAutoRename{}
	config     *server.Config
)

func init() {
	config = server.AppConfig
}

func AutoRename(ctx context.Context) func(s *discord.Session, c *discord.MessageCreate) {
	return func(s *discord.Session, c *discord.MessageCreate) {
		match, err := regexp.MatchString(config.MessagePattern, c.Content)
		if err != nil || !match {
			log.Warn().Msgf("Unsupported message: %s", c.Content)
			return
		}
		log.Info().Msgf("Supported message: %s", c.Content)
		autoRename.Rename(ctx, s, c)
	}
}
