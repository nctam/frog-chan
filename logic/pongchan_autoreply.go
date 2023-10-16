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
	pongLogTag = "PongChan"
)

func HandlePongChanMessage(ctx context.Context, config *server.Config, user *server.TargetUser, s *discord.Session, r *discord.MessageCreate) {
	logger := zerolog.Ctx(ctx)
	if !utils.StringContains(user.ChannelIDs, r.ChannelID) {
		logger.Info().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("Not in appropriate channel, you're lucky this time %s", user.Nickname)
		return
	}

	if !utils.ShouldReply(ctx, user.Probability, user.UniversalSet) {
		logger.Info().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("You're lucky this time %s", user.Nickname)
		return
	}

	logger.Info().Str(constant.AutoRepLogTag, pongLogTag).Msg("Pong chan detected, get messages")
	msg := message.PickMessage(message.DetectMessage(constant.PongMessages, r))
	if msg == nil {
		logger.Warn().Str(constant.AutoRepLogTag, pongLogTag).Msg("Don't know what to say, skip")
		return
	}

	if msg.ReactEmoji != "" {
		if err := s.MessageReactionAdd(r.ChannelID, r.ID, msg.ReactEmoji); err != nil {
			logger.Error().Str(constant.AutoRepLogTag, pongLogTag).Msgf("Unable to react message", err.Error())
		}
	}

	if msg.Message == "đồ đĩ" {
		taggedUserID := utils.ExtractTaggedUserID(r.Content, config)
		if taggedUserID != "" {
			msg.TagUser = taggedUserID
		}
		_, err := s.ChannelMessageSend(r.ChannelID, msg.Build())
		if err != nil {
			logger.Error().Str(constant.AutoRepLogTag, pongLogTag).Msgf("Unable to reply message", err.Error())
		}
	} else {
		_, err := s.ChannelMessageSendReply(r.ChannelID, msg.Build(), r.Reference())
		if err != nil {
			logger.Error().Str(constant.AutoRepLogTag, pongLogTag).Msgf("Unable to reply message", err.Error())
		}
	}
}
