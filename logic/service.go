package logic

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
)

type AutoReply interface {
	SendReply(context.Context, *discord.Session, *discord.MessageCreate)
}

type AutoReact interface {
	SendReact(ctx context.Context, s *discord.Session, r *discord.MessageReactionAdd)
}

type AutoReplyFunc func(context.Context, *discord.Session, *discord.MessageCreate)
