package message

import "fmt"

type TargetMessage struct {
	Message    string
	Emoji      string
	ReactEmoji string
	Url        string
}

func (m *TargetMessage) Build() string {
	res := m.Message
	if m.Emoji != "" {
		res = fmt.Sprintf("%s <:%s>", res, m.Emoji)
	}

	return res
}
