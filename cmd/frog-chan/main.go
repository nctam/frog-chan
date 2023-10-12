package main

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"kaeru.chan/voz/server"
	"os"
	"os/signal"
)

var (
	appConfig *server.Config
)

func main() {
	session, _ := discord.New("Bot " + appConfig.BotToken)
	session.AddHandler(func(s *discord.Session, r *discord.Ready) {
		log.Info().Msg("Bot is ready")
	})
	if err := session.Open(); err != nil {
		log.Error().Msg("Failed to open session")
	}

	defer func(sessionHandler *discord.Session) {
		err := sessionHandler.Close()
		if err != nil {
			log.Error().Msg("Failed to close session")
		}
	}(session)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Debug().Msg("Graceful shutdown")
}

func init() {
	ctx := context.Background()
	conf, err := server.ReadConfig(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Msg("Failed to load appConfig")
		return
	}
	appConfig = conf
}
