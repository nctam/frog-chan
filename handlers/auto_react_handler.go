package handlers

import (
	"context"
	discord "github.com/bwmarrin/discordgo"

	"kaeru.chan/voz/logic"
)

var (
	autoReact = logic.GeneralAutoReact{}
)

func AutoReact(ctx context.Context) func(s *discord.Session, r *discord.MessageReactionAdd) {
	return func(s *discord.Session, r *discord.MessageReactionAdd) {
		autoReact.SendReact(ctx, s, r)
	}
}
