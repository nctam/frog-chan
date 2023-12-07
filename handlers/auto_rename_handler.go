package handlers

import (
	"context"

	discord "github.com/bwmarrin/discordgo"
	"kaeru.chan/voz/logic"
)

var (
	autoRename = logic.GeneralAutoRename{}
)

func AutoRename(ctx context.Context) func(s *discord.Session, r *discord.MessageCreate) {
	return func(s *discord.Session, r *discord.MessageCreate) {
		autoRename.Rename(ctx, s, r)
	}
}
