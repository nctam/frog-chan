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
	// default validate
	validate = fp.Compose3(channelValidator, messageValidator, excludedUserValidator)
)

func AutoReply(ctx context.Context) func(s *discord.Session, r *discord.MessageCreate) {
	return func(s *discord.Session, r *discord.MessageCreate) {

		if strings.Contains(r.Content, "chửi") {
			validate = fp.Compose2(validate, curseRequestValidator)
		}

		validate(autoRep.SendReply)(ctx, s, r)
	}
}
