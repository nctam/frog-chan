package session

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

type CommandDispatcher struct {
	SessionManager *Manager
}

func NewCommandHandler() *CommandDispatcher {
	return &CommandDispatcher{
		SessionManager: NewManager(),
	}
}

func (c *CommandDispatcher) Handle(ctx context.Context, msg *discordgo.MessageCreate) bool {
	log := zerolog.Ctx(ctx).With().Str("Command", "SendReply").Logger()
	handler, ok := c.SessionManager.Get(msg.GuildID, msg.Author.ID)
	if !ok {
		return false
	}

	done, err := handler.Handle(msg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to handle message")
	}

	if done {
		c.SessionManager.Delete(msg.GuildID, msg.Author.ID)
	}

	return true
}

func (c *CommandDispatcher) CreateSession(
	msg *discordgo.MessageCreate,
	targetID string,
	duration time.Duration,
	execute CommandExecutor,
	reply func(channelID, content string) error,
) {
	c.SessionManager.Set(
		msg.GuildID,
		msg.Author.ID,
		&CommandSession{
			TargetID: targetID,
			Duration: duration,
			Execute:  execute,
			Reply:    reply,
		},
		30*time.Second,
	)
}
