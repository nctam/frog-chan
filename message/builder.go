package message

import (
	"fmt"
	discord "github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

type TargetMessage struct {
	Message    string
	Emoji      string
	ReactEmoji string
	Url        string
	HasRef     bool
	TagUser    string
}

func (m *TargetMessage) Build() string {
	res := m.Message
	if m.TagUser != "" {
		res = fmt.Sprintf("<@%s> %s", m.TagUser, res)
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

func PickMessage(messages []TargetMessage) *TargetMessage {
	if messages == nil || len(messages) == 0 {
		return nil
	}
	now := time.Now().Unix()
	messageIndex := now % int64(len(messages))
	return &messages[messageIndex]
}

func DetectMessage(message map[string][]TargetMessage, r *discord.MessageCreate) []TargetMessage {
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
