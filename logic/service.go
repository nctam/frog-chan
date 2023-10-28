package logic

import (
	"context"
	discord "github.com/bwmarrin/discordgo"
)

type AutoReply interface {
	SendReply(context.Context, *discord.Session, *discord.MessageCreate)
}

type AutoReplyFunc func(context.Context, *discord.Session, *discord.MessageCreate)
