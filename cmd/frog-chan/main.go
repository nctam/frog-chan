package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"

	"kaeru.chan/voz/handlers"
	"kaeru.chan/voz/server"
)

var (
	config       *server.Config
	readyHandler = handlers.Ready
	autoRep      = handlers.AutoReply
	autoReact    = handlers.AutoReact
	autoRename   = handlers.AutoRename
)

const (
	panicMsg = "Some problem when start app!"
)

func init() {
	config = server.AppConfig
}

func main() {
	ctx := context.TODO()
	ctx = server.ConfigLogger().WithContext(ctx)
	logger := zerolog.Ctx(ctx).With().Str("Main", "main()").Logger()
	session, _ := discord.New("Bot " + config.BotToken)
	session.AddHandler(readyHandler(ctx))
	session.AddHandler(autoRep(ctx))
	session.AddHandler(autoReact(ctx))
	session.AddHandler(autoRename(ctx))

	if err := session.Open(); err != nil {
		logger.Error().Msg("Failed to open session")
		panic(panicMsg)
	}

	handlers.RegisterHealthCheck()
	defer func(sessionHandler *discord.Session) {
		err := sessionHandler.Close()
		if err != nil {
			logger.Error().Msg("Failed to close session")
		}
	}(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop
	logger.Debug().Msg("Graceful shutdown!!")
}
