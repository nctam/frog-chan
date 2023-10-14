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
}

func (m *TargetMessage) Build() string {
	res := m.Message
	if m.Emoji != "" {
		res = fmt.Sprintf("%s <:%s>", res, m.Emoji)
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
