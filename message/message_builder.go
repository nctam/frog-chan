package message

import (
    "fmt"
    discord "github.com/bwmarrin/discordgo"
    "strings"
    "time"
)

type Template struct {
    Message    string
    Emoji      string
    ReactEmoji string
    Url        string
    HasRef     bool
    TagUsers   []string
}

func (m *Template) Build() string {
    res := m.Message
    if m.TagUsers != nil {
        for _, user := range m.TagUsers {
            res = fmt.Sprintf("%s <@%s>", res, user)
        }
    }
    if m.Emoji != "" {
        emoji := m.Emoji
        if strings.HasPrefix(m.Emoji, "a:") { // custom emoji(animated)
            res = fmt.Sprintf("%s <%s>", res, emoji)
        } else {
            res = fmt.Sprintf("%s <:%s>", res, emoji)
        }
    }

    return res
}

func PickMessage(messages []Template) *Template {
    if messages == nil || len(messages) == 0 {
        return nil
    }
    now := time.Now().Unix()
    messageIndex := now % int64(len(messages))
    return &messages[messageIndex]
}

func DetectMessage(message map[string][]Template, r *discord.MessageCreate) []Template {
    for k, v := range message {
        keys := strings.Split(k, ",")
        for _, key := range keys {
            if strings.Contains(r.Content, key) {
                return v
            }
        }
    }
    return nil
}
