package logic

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"kaeru.chan/voz/constant"
	"kaeru.chan/voz/message"
	"kaeru.chan/voz/server"
	"kaeru.chan/voz/utils"
	"strings"
)

const (
	communityLogTag = "Community"
)

var (
	commonMsg = map[string]message.TargetMessage{
		"đĩ": {
			Message:    "đồ đĩ",
			Emoji:      constant.Tsk,
			ReactEmoji: constant.Aragorn,
			Url:        "",
		},
		"bitch": {
			Message:    "",
			Emoji:      constant.Tsk,
			ReactEmoji: "",
		},
		"buscu": {
			Message:    "",
			Emoji:      "",
			ReactEmoji: "",
		},
		"ktv": {
			Message:    "",
			Emoji:      "",
			ReactEmoji: "",
		},
		"lấy": {
			Message:    "",
			Emoji:      "",
			ReactEmoji: "",
		},
		"mua": {
			Message:    "",
			Emoji:      "",
			ReactEmoji: "",
		},
		"póng": {
			Message:    "póng póng cc, hãy để trú yên, trú còn bận đi úp bô",
			Emoji:      constant.Tsk,
			ReactEmoji: "",
		},
		"trú": {
			Message:    "trú bận đi úp bô gồi",
			Emoji:      constant.Tsk,
			ReactEmoji: "",
		},
	}
)

func HandleMessage(ctx context.Context, config *server.Config, _ *server.TargetUser, s *discord.Session, r *discord.MessageCreate) {
	logger := zerolog.Ctx(ctx)
	if !utils.StringContains(config.Channels, r.ChannelID) {
		logger.Info().Str(constant.AutoRepLogTag, "Handler").Msgf("Not in appropriate channel")
		return
	}

	if !utils.ShouldAimTo(ctx, 0.2, 100) {
		logger.Info().Str(constant.AutoRepLogTag, communityLogTag).Msgf("You're lucky this time %s", r.Author.Username)
		return
	}

	logger.Info().Str(constant.AutoRepLogTag, communityLogTag).Msg("Community detected, start bullying")
	msg := detectCommunityMessage(r)
	if msg == nil {
		logger.Warn().Str(constant.AutoRepLogTag, communityLogTag).Msg("Don't know what to say, skip")
		return
	}

	if msg.ReactEmoji != "" {
		if err := s.MessageReactionAdd(r.ChannelID, r.ID, msg.ReactEmoji); err != nil {
			logger.Error().Str(constant.AutoRepLogTag, communityLogTag).Msgf("Unable to react message", err.Error())
		}
	}

	_, err := s.ChannelMessageSendReply(r.ChannelID, msg.Build(), r.Reference())
	if err != nil {
		logger.Error().Str(constant.AutoRepLogTag, communityLogTag).Msgf("Unable to reply message", err.Error())
	}
}

func detectCommunityMessage(r *discord.MessageCreate) *message.TargetMessage {
	for k, v := range commonMsg {
		if strings.Contains(r.Content, k) {
			return &v
		}
	}
	return nil
}
