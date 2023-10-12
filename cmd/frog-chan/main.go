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
	readyHandler = handlers.Ready
	logger       *zerolog.Logger
)

func main() {
	ctx := context.TODO()
	ctx = logger.WithContext(ctx)
	session, _ := discord.New("Bot " + appConfig.BotToken)
	session.AddHandler(readyHandler(ctx))
	if err := session.Open(); err != nil {
		logger.Error().Msg("Failed to open session")
	}

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

func init() {
	ctx := context.TODO()
	ctx = server.ConfigLogger(ctx)
	conf, err := server.ReadConfig(ctx)
	logger = zerolog.Ctx(ctx)
	if err != nil {
		logger.Error().Msg("Failed to load appConfig")
		return
	}
	appConfig = conf
}
