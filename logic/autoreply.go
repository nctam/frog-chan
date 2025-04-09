package logic

import (
	"context"
	"strings"

	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"

	"kaeru.chan/voz/constant"
	"kaeru.chan/voz/message"
	"kaeru.chan/voz/server"
	"kaeru.chan/voz/utils"
)

const (
	communityLogTag = "GeneralReply"
)

type GeneralAutoReply struct{}

var (
	config *server.Config
	_      AutoReply = &GeneralAutoReply{}
)

func init() {
	config = server.AppConfig
}

func (g *GeneralAutoReply) SendReply(ctx context.Context, s *discord.Session, r *discord.MessageCreate) {
	log := zerolog.Ctx(ctx).With().Str(communityLogTag, "SendReply").Logger()
	if !utils.ShouldReply(ctx, r.Author.ID, config) {
		log.Warn().Msgf("Reject reply user: %s", r.Author.Username)
		return
	}

	log.Info().Msgf("Accept reply user: %s", r.Author.Username)
	msg := message.PickMessage(message.DetectMessage(constant.CommunityMessages, r))

	if msg == nil {
		log.Warn().Msg("Message template not found")
		return
	}

	if taggedUsers := utils.ExtractTaggedUserID(ctx, r.Content, config); taggedUsers != nil {
		msg.TagUsers = taggedUsers
	}

	if msg.ReactEmoji != "" {
		if err := s.MessageReactionAdd(r.ChannelID, r.ID, msg.ReactEmoji); err != nil {
			log.Error().Err(err).Msgf("Unable to react message")
		}
	}

	var sendMsgErr error
	msgToSend := &discord.MessageEmbed{
		Description: msg.Build(),
		Image: &discord.MessageEmbedImage{
			URL: msg.Url,
		},
	}

	if strings.Contains(r.Content, "sạc") && r.Author.ID == constant.PongChanID {
		msg.TagUsers = []string{r.Author.ID}
		msgToSend.Description = msg.Build()
	} else {
		log.Warn().Msgf("Reject reply user: %s with %s", r.Author.Username, "sạc")
		return
	}

	if msg.HasRef {
		_, sendMsgErr = s.ChannelMessageSendEmbedReply(r.ChannelID, msgToSend, r.Reference())
	} else {
		_, sendMsgErr = s.ChannelMessageSendEmbed(r.ChannelID, msgToSend)
	}

	if sendMsgErr != nil {
		log.Error().Err(sendMsgErr).Msg("Unable to reply message")
	}
}
