package logic

import (
    "context"
    "fmt"
    "net/url"
    "regexp"
    "strings"

    "golang.org/x/exp/rand"

    discord "github.com/bwmarrin/discordgo"
    "github.com/rs/zerolog"

    "kaeru.chan/voz/constant"
    "kaeru.chan/voz/message"
    "kaeru.chan/voz/server"
    "kaeru.chan/voz/utils"
)

const (
    communityLogTag = "GeneralReply"
)

var (
    fbRegex = regexp.MustCompile(
        `https?://(?:www\.)?facebook\.com/(?:watch/?\?v=|share/[rv]/|reel/|[^/]+/videos/|story\.php\?story_fbid=)[a-zA-Z0-9_\-\&\=\?]+`,
    )

    embedDomain = "facecot.com"
)

type GeneralAutoReply struct{}

var (
    config *server.Config
    _      AutoReply = &GeneralAutoReply{}
)

func init() {
    config = server.AppConfig
}

func (g *GeneralAutoReply) SendReply(ctx context.Context, s *discord.Session, r *discord.MessageCreate) {
    log := zerolog.Ctx(ctx).With().Str(communityLogTag, "SendReply").Logger()
    if r.Author.ID == s.State.User.ID {
        return
    }

    fixedURL, platform, found := extractURL(r.Content)
    if found {
        if sendEmbeddedErr := sendEmbedMessage(s, r.ChannelID, r.Author.ID, platform, fixedURL); sendEmbeddedErr == nil {
            _ = s.ChannelMessageDelete(r.ChannelID, r.ID)
            return
        } else {
            return
        }
    }

    if !utils.ShouldReply(ctx, r.Author.ID, config) {
        log.Warn().Msgf("Reject reply user: %s", r.Author.Username)
        return
    }

    log.Info().Msgf("Accept reply user: %s", r.Author.Username)
    msg := message.PickMessage(message.DetectMessage(constant.CommunityMessages, r))

    if msg == nil {
        log.Warn().Msg("Message template not found")
        return
    }

    if taggedUsers := utils.ExtractTaggedUserID(ctx, r.Content, config); taggedUsers != nil {
        msg.TagUsers = taggedUsers
    }

    if msg.ReactEmoji != "" {
        if err := s.MessageReactionAdd(r.ChannelID, r.ID, msg.ReactEmoji); err != nil {
            log.Error().Err(err).Msgf("Unable to react message")
        }
    }

    var sendMsgErr error
    msgToSend := &discord.MessageEmbed{
        Description: msg.Build(),
        Image: &discord.MessageEmbedImage{
            URL: msg.Url,
        },
    }

    if strings.Contains(r.Content, "chửi") && msg.Build() == "" {
        msg.TagUsers = []string{r.Author.ID}
        msg.Message = constant.MsgReplyTagged[rand.Intn(len(constant.MsgReplyTagged))]
        msgToSend.Description = msg.Build()
    }

    if strings.Contains(r.Content, "sạc") {
        if r.Author.ID == constant.PongChanID || r.Author.ID == constant.SevenBabyID || r.Author.ID == constant.DucToID {
            msg.TagUsers = []string{r.Author.ID}
            msgToSend.Description = msg.Build()
        } else {
            log.Warn().Msgf("Reject reply user: %s with %s", r.Author.Username, "sạc")
            return
        }
    }

    if msg.HasRef {
        _, sendMsgErr = s.ChannelMessageSendEmbedReply(r.ChannelID, msgToSend, r.Reference())
    } else {
        _, sendMsgErr = s.ChannelMessageSendEmbed(r.ChannelID, msgToSend)
    }

    if sendMsgErr != nil {
        log.Error().Err(sendMsgErr).Msg("Unable to reply message")
    }
}

func extractURL(content string) (string, string, bool) {
    // Facebook
    if fbRegex.MatchString(content) {
        urlStr := fbRegex.FindString(content)
        u, err := url.Parse(urlStr)
        if err != nil {
            return "", "", false
        }

        if u.Host == "www.facebook.com" {
            u.Host = "www." + embedDomain
        } else if u.Host == "facebook.com" {
            u.Host = embedDomain
        } else {
            return "", "", false
        }
        return u.String(), "Facebook", true
    }

    return "", "", false
}

func sendEmbedMessage(
    s *discord.Session,
    channelID string,
    authorID string,
    platform string,
    url string,
) error {
    msg := fmt.Sprintf(
        "\n%s\n\nsauce: <@%s>",
        url,
        authorID,
    )

    _, err := s.ChannelMessageSend(channelID, msg)
    return err
}
