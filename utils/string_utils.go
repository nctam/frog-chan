package utils

import (
	"kaeru.chan/voz/server"
	"regexp"
	"strings"
)

const (
	taggedUserRegex = "<@[0-9]+>"
)

func ExtractTaggedUserID(msg string, config *server.Config) string {
	pattern := regexp.MustCompile(taggedUserRegex)
	if pattern == nil {
		return ""
	}
	atIndex := strings.Index(msg, "@")
	closeTagIndex := strings.Index(msg, ">")

	userID := msg[atIndex+1 : closeTagIndex]
	if StringContains(config.ExcludedUsers, userID) {
		return ""
	}

	return userID
}
