package handlers

import (
	"context"

	discord "github.com/bwmarrin/discordgo"
	fp "github.com/repeale/fp-go"
	"github.com/rs/zerolog"
	"kaeru.chan/voz/decorator"
	"kaeru.chan/voz/logic"
	"kaeru.chan/voz/server"
	"kaeru.chan/voz/utils"
)

var (
	autoRep               = logic.GeneralAutoReply{}
	channelValidator      = decorator.ValidateChannel
	messageValidator      = decorator.ValidateMessage
	excludedUserValidator = decorator.ValidateExcludedUser
	curseRequestValidator = decorator.ValidateCurseRequest
)

func init() {
	config = server.AppConfig
}

func AutoReply(ctx context.Context) func(s *discord.Session, r *discord.MessageCreate) {

	return func(s *discord.Session, r *discord.MessageCreate) {
		// default validate
		validate := fp.Compose2(channelValidator, excludedUserValidator)
		pattern := "chửi"
		log := zerolog.Ctx(ctx).With().Str(logTag, "AutoReply").Logger()
		logMsg := "Unmatching with pattern --> '%v' with message --> '%v'"
		if utils.ExtractMessage(pattern, r.Content) {
			logMsg = "Matching with pattern --> '%v' with message --> '%v'"
			validate = fp.Compose2(validate, curseRequestValidator)
		}
		if config.Env == "debug" {
			log.Info().Msgf(logMsg, pattern, r.Content)
		}
		validate(autoRep.SendReply)(ctx, s, r)
	}
}
