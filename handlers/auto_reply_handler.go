package handlers

import (
	"context"
	"strings"

	discord "github.com/bwmarrin/discordgo"

	fp "github.com/repeale/fp-go"
	"kaeru.chan/voz/decorator"
	"kaeru.chan/voz/logic"
)

var (
	autoRep               = logic.GeneralAutoReply{}
	channelValidator      = decorator.ValidateChannel
	messageValidator      = decorator.ValidateMessage
	excludedUserValidator = decorator.ValidateExcludedUser
	curseRequestValidator = decorator.ValidateCurseRequest
	// validate              = decorator.Validate
)

func AutoReply(ctx context.Context) func(s *discord.Session, r *discord.MessageCreate) {
	return func(s *discord.Session, r *discord.MessageCreate) {
		var f func(logic.AutoReplyFunc) logic.AutoReplyFunc
		if strings.Contains(r.Content, "chửi") {
			f = fp.Compose4(channelValidator, messageValidator, excludedUserValidator, curseRequestValidator)
		} else {
			f = fp.Compose3(channelValidator, messageValidator, excludedUserValidator)
		}

		f(autoRep.SendReply)(ctx, s, r)
	}
}
