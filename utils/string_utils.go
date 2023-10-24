package utils

import (
	"context"
	"github.com/rs/zerolog"
	"kaeru.chan/voz/server"
	"regexp"
	"strings"
)

const (
	taggedUserRegex = "(\\w?)<@[0-9]+>(\\w?)"
)

func ExtractTaggedUserID(ctx context.Context, msg string, config *server.Config) string {
	log := zerolog.Ctx(ctx).With().Str("Utils", "ExtractTaggedUserID").Logger()
	match, err := regexp.MatchString(taggedUserRegex, msg)
	if err != nil || !match {
		log.Debug().Msgf("No tagged user detected: '%s'", msg)
		return ""
	}
	atIndex := strings.Index(msg, "@")
	closeTagIndex := strings.Index(msg, ">")

	userID := msg[atIndex+1 : closeTagIndex]
	if StringContains(config.ExcludedUsers, userID) {
		log.Debug().Msgf("User: %s is excluded", userID)
		return ""
	}

	return userID
}
