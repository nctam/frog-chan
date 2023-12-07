package logic

import (
	"context"
	"fmt"
	"time"

	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"kaeru.chan/voz/server"

	constant "kaeru.chan/voz/constant"
	internal "kaeru.chan/voz/message"
)

const (
	communityLogRename = "GeneralRename"
)

func init() {
	config = server.AppConfig
}

type GeneralAutoRename struct {
}

type NickName struct {
	oldNick string
	newNick string
}

func (g *GeneralAutoRename) SendMsgRename(ctx context.Context, s *discord.Session, r *discord.MessageCreate, msg internal.Template) {
	_, err := s.ChannelMessageSend(r.ChannelID, msg.Build())
	if err != nil {
		log.Err(err).Msgf("Unable to send message %s to channel %s", msg.Message, r.ChannelID)
	}
}

func (g *GeneralAutoRename) Rename(ctx context.Context, s *discord.Session, r *discord.MessageCreate) {
	log := zerolog.Ctx(ctx).With().Str(communityLogRename, "Rename").Logger()
	log.Info().Msgf("Prepare rename user: %s", r.Author.Username)

	oldName := r.Message.Member.Nick
	newName := constant.MemberNickNames[0]

	nick := NickName{
		oldNick: oldName,
		newNick: newName,
	}

	msg := internal.Template{
		Message: fmt.Sprintf("<@%s> đổi tên <@%s> thành %s", config.Bots[0], nick.oldNick, nick.newNick),
	}

	g.SendMsgRename(ctx, s, r, msg)
	log.Info().Msgf("Rename user: %s", r.Author.Username)

	go g.RollbackName(ctx, s, r, nick)
}

func (g *GeneralAutoRename) RollbackName(ctx context.Context, s *discord.Session, r *discord.MessageCreate, nick NickName) {
	log := zerolog.Ctx(ctx).With().Str(communityLogRename, "Rename").Logger()
	log.Info().Msgf("Prepare rename user: %s", r.Author.Username)

	msg := internal.Template{
		Message: fmt.Sprintf("<@%s> đổi tên <@%s> thành %s", config.Bots[0], nick.newNick, nick.oldNick),
	}
	time.Sleep(5 * time.Minute)
	g.SendMsgRename(ctx, s, r, msg)
	log.Info().Msgf("Rollback rename user: %s", nick.oldNick)
}
