package logic

import (
	"context"
	"fmt"
	"time"

	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"istio.io/pkg/cache"
	"kaeru.chan/voz/server"

	"golang.org/x/exp/rand"
	constant "kaeru.chan/voz/constant"
	internal "kaeru.chan/voz/message"
)

const (
	communityLogRename = "GeneralRename"
	timeTTL            = 5
)

var (
	cacheNickName cache.ExpiringCache
	_             AutoRename = &GeneralAutoRename{}
)

func init() {
	config = server.AppConfig
	cacheNickName = cache.NewTTL(timeTTL*time.Minute, timeTTL*time.Minute)
}

type GeneralAutoRename struct {
}

type NickName struct {
	oldNick string
	newNick string
	userId  string
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

	if r.Author.ID == constant.Bot {
		return
	}

	oldName := r.Member.Nick
	newName := constant.MemberNickNames[rand.Intn(len(constant.MsgReplyTagged))]

	nick := NickName{
		oldNick: oldName,
		newNick: newName,
		userId:  r.Author.ID,
	}

	value, _ := cacheNickName.Get(nick.userId)
	if value != nil && value.(int) > 0 {
		log.Info().Msgf("Rename from %s has processed", r.Author.Username)
		return
	}

	cacheNickName.Set(nick.userId, 1)
	err := s.GuildMemberNickname(r.GuildID, nick.userId, nick.newNick)
	if err != nil {
		log.Error().Msgf("Rename %s failed", r.Author.Username)
		return
	}

	msg := internal.Template{
		Message: fmt.Sprintf("vì tội phỉ báng nên bị đổi tên <@%s> ", nick.userId),
	}

	g.SendMsgRename(ctx, s, r, msg)
	log.Info().Msgf("Rename user: %s", r.Author.Username)

	go g.RollbackName(ctx, s, r, nick)
}

func (g *GeneralAutoRename) RollbackName(ctx context.Context, s *discord.Session, r *discord.MessageCreate, nick NickName) {
	log := zerolog.Ctx(ctx).With().Str(communityLogRename, "Rename").Logger()
	log.Info().Msgf("Prepare rename user: %s", r.Author.Username)

	err := s.GuildMemberNickname(r.GuildID, nick.userId, nick.newNick)
	if err != nil {
		log.Error().Msgf("Rename %s failed", r.Author.Username)
		return
	}

	msg := internal.Template{
		Message: fmt.Sprintf("Tha thứ cho lần này <@%s> ", nick.userId),
	}
	time.Sleep(timeTTL * time.Minute)
	g.SendMsgRename(ctx, s, r, msg)
	log.Info().Msgf("Rollback rename user: %s", nick.oldNick)
	value, _ := cacheNickName.Get(nick.userId)
	if value != nil {
		cacheNickName.Remove(nick.userId)
	}
}
