package logic

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"kaeru.chan/voz/constant"
	"kaeru.chan/voz/message"
	"kaeru.chan/voz/server"
	"kaeru.chan/voz/utils"
)

const (
	fonfonLogTag = "FonFon"
)

func HandleFonFonMessage(ctx context.Context, _ *server.Config, user *server.TargetUser, s *discord.Session, r *discord.MessageCreate) {
	logger := zerolog.Ctx(ctx)
	if !utils.StringContains(user.ChannelIDs, r.ChannelID) {
		logger.Info().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("Not in appropriate channel, you're lucky this time %s", user.Nickname)
		return
	}

	if !utils.ShouldReply(ctx, user.Probability, user.UniversalSet) {
		logger.Info().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("You're lucky this time %s", user.Nickname)
		return
	}

	logger.Info().Str(constant.AutoRepLogTag, fonfonLogTag).Msg("FonFon detected, get messages")
	msg := message.PickMessage(message.DetectMessage(constant.FonfonMessages, r))
	if msg == nil {
		logger.Warn().Str(constant.AutoRepLogTag, fonfonLogTag).Msg("Don't know what to say, skip")
		return
	}

	if msg.ReactEmoji != "" {
		if err := s.MessageReactionAdd(r.ChannelID, r.ID, msg.ReactEmoji); err != nil {
			logger.Error().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("Unable to react message", err.Error())
		}
	}

	_, err := s.ChannelMessageSendReply(r.ChannelID, msg.Build(), r.Reference())
	if err != nil {
		logger.Error().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("Unable to reply message", err.Error())
	}

}
