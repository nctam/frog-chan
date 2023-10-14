package handlers

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"kaeru.chan/voz/constant"
	"kaeru.chan/voz/logic"
	"kaeru.chan/voz/server"
	"kaeru.chan/voz/utils"
)

var (
	handleMsg      = logic.HandleMessage
	handleFonFon   = logic.HandleFonFonMessage
	handlePongChan = logic.HandlePongChanMessage
)

func AutoReply(ctx context.Context, config *server.Config) func(s *discord.Session, r *discord.MessageCreate) {
	logger := zerolog.Ctx(ctx)
	return func(s *discord.Session, r *discord.MessageCreate) {
		userName := r.Author.Username
		invokedUser := utils.FindUserByID(r.Author.ID, config.TargetUsers)
		if invokedUser != nil {
			logger.Info().Str(constant.AutoRepLogTag, "Handler").Msg("Target user found")
			userName = invokedUser.Nickname
		}

		switch userName {
		case constant.FonFonName, constant.EchChan:
			handleFonFon(ctx, config, invokedUser, s, r)
		case constant.PongChanName:
			handlePongChan(ctx, config, invokedUser, s, r)
		default:
			handleMsg(ctx, config, invokedUser, s, r)
		}
	}
}
