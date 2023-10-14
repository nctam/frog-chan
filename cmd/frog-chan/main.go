package main

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"kaeru.chan/voz/handlers"
	"kaeru.chan/voz/server"
	"os"
	"os/signal"
)

var (
	appConfig    *server.Config
	logger       *zerolog.Logger
	readyHandler = handlers.Ready
	autoRep      = handlers.AutoReply
)

func main() {
	ctx := context.TODO()
	ctx = server.ConfigLogger(ctx)
	appConfig = server.ReadConfig(ctx)
	session, _ := discord.New("Bot " + appConfig.BotToken)
	session.AddHandler(readyHandler(ctx))
	session.AddHandler(autoRep(ctx, appConfig))
	if err := session.Open(); err != nil {
		logger.Error().Msg("Failed to open session")
	}

	handlers.RegisterHealthCheck()
	defer func(sessionHandler *discord.Session) {
		err := sessionHandler.Close()
		if err != nil {
			logger.Error().Msg("Failed to close session")
		}
	}(session)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	logger.Debug().Msg("Graceful shutdown")
}
