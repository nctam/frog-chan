package decorator

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"regexp"

	"kaeru.chan/voz/constant"
	"kaeru.chan/voz/logic"
	"kaeru.chan/voz/server"
	"kaeru.chan/voz/utils"
)

const (
	logTag = "Validation"
)

var (
	config *server.Config
)

func init() {
	config = server.AppConfig
}

type AutoReplyValidation func(logic.AutoReplyFunc) logic.AutoReplyFunc

func Validate(f logic.AutoReplyFunc, v ...AutoReplyValidation) logic.AutoReplyFunc {
	if len(v) == 0 {
		return f
	}

	return v[0](Validate(f, v[1:]...))
}

func ValidateMessage(f logic.AutoReplyFunc) logic.AutoReplyFunc {
	return func(ctx context.Context, s *discord.Session, c *discord.MessageCreate) {
		log := zerolog.Ctx(ctx).With().Str(logTag, "ValidateMessage").Logger()
		match, err := regexp.MatchString(constant.MessagePattern, c.Content)
		if err != nil || !match {
			log.Warn().Msgf("Unsupported message: %s", c.Content)
			return
		}
		log.Info().Msgf("Supported message: %s", c.Content)
		f(ctx, s, c)
	}
}

func ValidateChannel(f logic.AutoReplyFunc) logic.AutoReplyFunc {
	return func(ctx context.Context, s *discord.Session, c *discord.MessageCreate) {
		log := zerolog.Ctx(ctx).With().Str(logTag, "ValidateChannel").Logger()
		if !utils.StringContains(config.Channels, c.ChannelID) {
			log.Warn().Msgf("Unsupported channel: %s", c.ChannelID)
			return
		}
		log.Info().Msgf("Supported channel: %s", c.ChannelID)
		f(ctx, s, c)
	}
}

func ValidateExcludedUser(f logic.AutoReplyFunc) logic.AutoReplyFunc {
	return func(ctx context.Context, s *discord.Session, c *discord.MessageCreate) {
		log := zerolog.Ctx(ctx).With().Str(logTag, "ValidateExcludedUser").Logger()
		if utils.StringContains(config.ExcludedUsers, c.Author.ID) {
			log.Warn().Msgf("Unsupported user: %s", c.Author.Username)
			return
		}
		log.Info().Msgf("Supported user: %s", c.Author.Username)
		f(ctx, s, c)
	}
}
