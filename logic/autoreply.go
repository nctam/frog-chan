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
	communityLogTag = "Community"
)

func HandleMessage(ctx context.Context, config *server.Config, _ *server.TargetUser, s *discord.Session, r *discord.MessageCreate) {
	logger := zerolog.Ctx(ctx)
	if !utils.StringContains(config.Channels, r.ChannelID) {
		logger.Info().Str(constant.AutoRepLogTag, communityLogTag).Msgf("Not in appropriate channel")
		return
	}

	prob := 0.4
	//if r.Author.ID == constant.EchChanID {
	//	prob = 1.0
	//}

	if !utils.ShouldReply(ctx, prob, 100) {
		logger.Info().Str(constant.AutoRepLogTag, communityLogTag).Msgf("You're lucky this time %s", r.Author.Username)
		return
	}

	logger.Info().Str(constant.AutoRepLogTag, communityLogTag).Msg("Community detected, get messages")
	msg := message.PickMessage(message.DetectMessage(constant.CommunityMessages, r))
	if msg == nil {
		logger.Warn().Str(constant.AutoRepLogTag, communityLogTag).Msg("Don't know what to say, skip")
		return
	}

	if msg.ReactEmoji != "" {
		if err := s.MessageReactionAdd(r.ChannelID, r.ID, msg.ReactEmoji); err != nil {
			logger.Error().Str(constant.AutoRepLogTag, communityLogTag).Msgf("Unable to react message", err.Error())
		}
	}

	var sendMsgErr error
	if msg.HasRef {
		_, sendMsgErr = s.ChannelMessageSendReply(r.ChannelID, msg.Build(), r.Reference())
	} else {
		_, sendMsgErr = s.ChannelMessageSend(r.ChannelID, msg.Build())
	}

	if sendMsgErr != nil {
		logger.Error().Str(constant.AutoRepLogTag, communityLogTag).Msgf("Unable to reply message", sendMsgErr.Error())
	}
}
