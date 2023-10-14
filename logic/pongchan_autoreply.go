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
	pongLogTag = "PongChan"
)

var (
	pongMessage = map[string]message.TargetMessage{
		"chửi": {
			Message:    "đồ đĩ, làm gì chửi người ta",
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
			Message:    "póng póng cc, trú muốn tự chửi mình à?",
			Emoji:      constant.Tsk,
			ReactEmoji: "",
		},
		"trú": {
			Message:    "trú bận đi buscu rồi",
			Emoji:      constant.Tsk,
			ReactEmoji: "",
		},
	}
)

func HandlePongChanMessage(ctx context.Context, config *server.Config, user *server.TargetUser, s *discord.Session, r *discord.MessageCreate) {
	logger := zerolog.Ctx(ctx)
	if !utils.StringContains(user.ChannelIDs, r.ChannelID) {
		logger.Info().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("Not in appropriate channel, you're lucky this time %s", user.Nickname)
		return
	}

	if !utils.ShouldAimTo(ctx, user.Probability, user.UniversalSet) {
		logger.Info().Str(constant.AutoRepLogTag, fonfonLogTag).Msgf("You're lucky this time %s", user.Nickname)
		return
	}

	logger.Info().Str(constant.AutoRepLogTag, pongLogTag).Msg("Pong chan detected, start bullying")
	msg := detectPongMessage(r)
	if msg == nil {
		logger.Warn().Str(constant.AutoRepLogTag, pongLogTag).Msg("Don't know what to say, skip")
		return
	}

	if msg.ReactEmoji != "" {
		if err := s.MessageReactionAdd(r.ChannelID, r.ID, msg.ReactEmoji); err != nil {
			logger.Error().Str(constant.AutoRepLogTag, pongLogTag).Msgf("Unable to react message", err.Error())
		}
	}

	_, err := s.ChannelMessageSendReply(r.ChannelID, msg.Build(), r.Reference())
	if err != nil {
		logger.Error().Str(constant.AutoRepLogTag, pongLogTag).Msgf("Unable to reply message", err.Error())
	}
}

func detectPongMessage(r *discord.MessageCreate) *message.TargetMessage {
	for k, v := range pongMessage {
		if strings.Contains(r.Content, k) {
			return &v
		}
	}
	return nil
}
