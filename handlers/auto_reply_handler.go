package handlers

import (
	"context"
	discord "github.com/bwmarrin/discordgo"

	"kaeru.chan/voz/decorator"
	"kaeru.chan/voz/logic"
)

var (
	autoRep               = logic.GeneralAutoReply{}
	channelValidator      = decorator.ValidateChannel
	messageValidator      = decorator.ValidateMessage
	excludedUserValidator = decorator.ValidateExcludedUser
	curseRequestValidator = decorator.ValidateCurseRequest
	validate              = decorator.Validate
)

func AutoReply(ctx context.Context) func(s *discord.Session, r *discord.MessageCreate) {
	return func(s *discord.Session, r *discord.MessageCreate) {
		validations := []decorator.AutoReplyValidation{
			channelValidator,
			messageValidator,
			excludedUserValidator,
			curseRequestValidator,
		}
		f := validate(autoRep.SendReply, validations...)
		f(ctx, s, r)
	}
}
