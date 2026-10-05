package session

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"kaeru.chan/voz/server"
)

type CommandExecutor func(guildID, targetID string, duration time.Duration) error

type CommandSession struct {
	TargetID string
	Duration time.Duration
	Execute  CommandExecutor
	Reply    func(channelID, content string) error
}

type Intent int
type Command int

const (
	IntentUnknown Intent = iota
	IntentConfirm
	IntentCancel
)

const (
	CommandUnknown Command = iota
	CommandTimeout
	CommandRemoveTimeout
	CommandNickname
	CommandRole
	CommandPileOn
)

var (
	config          *server.Config
	intentPatterns  map[Intent]*regexp.Regexp
	commandPatterns map[Command][]*regexp.Regexp
)

func init() {
	config = server.AppConfig
	initIntentPatterns()
	initCommandPatterns()
}

func (s *CommandSession) Handle(msg *discordgo.MessageCreate) (bool, error) {
	var exeRes bool
	content := strings.ToLower(strings.TrimSpace(msg.Content))
	intent := parseIntent(content)
	_ = s.Execute(msg.GuildID, s.TargetID, s.Duration)
	exeRes = true

	//switch cmd {
	//case CommandPileOn:
	//	_ = s.Execute(msg.GuildID, s.TargetID, s.Duration)
	//	exeRes = true
	//default: // try to test pile on first
	//	exeRes = false
	//}

	if exeRes {
		return exeRes, nil
	}

	switch intent {
	case IntentConfirm:
		err := s.Execute(msg.GuildID, s.TargetID, s.Duration)
		if err != nil {
			_ = s.Reply(
				msg.ChannelID,
				"timeout failed 💀",
			)
			return true, err
		}

		err = s.Reply(
			msg.ChannelID,
			fmt.Sprintf(
				"ok, timeout <@%s> %s",
				s.TargetID,
				formatDuration(s.Duration),
			),
		)

		return true, err

	case IntentCancel:
		return true, s.Reply(
			msg.ChannelID,
			"thứ gì, rảnh háng",
		)
	default:
		return false, nil
	}
}

func formatDuration(d time.Duration) string {
	switch {
	case d%time.Hour == 0:
		return fmt.Sprintf("%dh", int(d/time.Hour))
	case d%time.Minute == 0:
		return fmt.Sprintf("%dm", int(d/time.Minute))
	default:
		return d.String()
	}
}

func parseIntent(content string) Intent {
	for intent, pattern := range intentPatterns {
		if pattern.MatchString(content) {
			return intent
		}
	}
	return IntentUnknown
}

func initIntentPatterns() {
	intentPatterns = make(map[Intent]*regexp.Regexp)

	for intent, pattern := range config.Session.IntentPatterns {
		var intentType Intent

		switch intent {
		case "confirm":
			intentType = IntentConfirm
		case "cancel":
			intentType = IntentCancel
		default:
			continue
		}

		intentPatterns[intentType] = regexp.MustCompile(pattern)
	}
}

func initCommandPatterns() {
	commandPatterns = make(map[Command][]*regexp.Regexp)

	for name, patterns := range config.Session.CommandPatterns {
		cmd := commandFromString(name)
		if cmd == CommandUnknown {
			continue
		}

		for _, pattern := range patterns {
			commandPatterns[cmd] = append(
				commandPatterns[cmd],
				regexp.MustCompile(pattern),
			)
		}
	}
}

func commandFromString(value string) Command {
	switch value {
	case "timeout":
		return CommandTimeout
	case "remove-timeout":
		return CommandRemoveTimeout
	case "nickname":
		return CommandNickname
	case "role":
		return CommandRole
	case "pile-on":
		return CommandPileOn
	default:
		return CommandUnknown
	}
}
