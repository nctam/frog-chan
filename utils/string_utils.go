package utils

import (
	"context"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	"kaeru.chan/voz/constant"
	"kaeru.chan/voz/server"
)

const (
	taggedUserRegex = "(\\w?)<@[0-9]+>(\\w?)"
)

func ExtractTaggedUserID(ctx context.Context, msg string, config *server.Config) []string {
	log := zerolog.Ctx(ctx).With().Str("Utils", "ExtractTaggedUserID").Logger()
	var users []string
	match, err := regexp.MatchString(taggedUserRegex, msg)
	if err != nil || !match {
		log.Debug().Msgf("No tagged user detected: '%s'", msg)
		return nil
	}

	for {
		atIndex := strings.Index(msg, "@")

		if atIndex < 0 {
			break
		}

		closeTagIndex := strings.Index(msg, ">")
		userID := msg[atIndex+1 : closeTagIndex]
		msg = msg[closeTagIndex+1:]
		if StringContains(config.ExcludedUsers, userID) || constant.EchChanID == userID {
			log.Debug().Msgf("User is excluded: %s", userID)
			continue
		}
		users = append(users, userID)
	}

	return users
}

func ExtractMessage(pattern, msg string) bool {
	match, err := regexp.MatchString(pattern, msg)
	if err != nil || !match {
		return false
	}

	return true
}
