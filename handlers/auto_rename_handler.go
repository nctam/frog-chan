package handlers

import (
	"context"
	"regexp"

	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
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
		log := zerolog.Ctx(ctx).With().Str(logTag, "AutoRename").Logger()
		match, err := regexp.MatchString(config.IotioPattern, c.Content)
		if err != nil || !match {
			log.Info().Msgf("Unsupported rename with message: %s", c.Content)
			return
		}

		log.Info().Msgf("Supported rename message: %s", c.Content)
		channelValidator(autoRename.Rename)(ctx, s, c)
	}
}
