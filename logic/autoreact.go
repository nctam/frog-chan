package logic

import (
    "context"
    discord "github.com/bwmarrin/discordgo"
    "github.com/rs/zerolog"

    "kaeru.chan/voz/constant"
    "kaeru.chan/voz/utils"
)

const (
    autoReactLogTag = "GeneralAutoReact"
)

type GeneralAutoReact struct{}

func (g *GeneralAutoReact) SendReact(ctx context.Context, s *discord.Session, r *discord.MessageReactionAdd) {
    log := zerolog.Ctx(ctx).With().Str(autoReactLogTag, "SendReact").Logger()
    message, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
    msgAuthor := message.Author.ID
    reactAuthor := r.Member.User.ID
    if utils.StringContains(config.Bots, reactAuthor) && msgAuthor == constant.Bot {
        log.Debug().Msgf("Reaction from %s found, removing", r.Member.User.Username)
        if err := s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.APIName(), reactAuthor); err != nil { // fucking naming emoji params
            log.Err(err).Msgf("Unable to remove emoji %s from user %s", r.Emoji.ID, r.Member.User.Username)
            return
        }
    }
}
