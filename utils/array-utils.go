package utils

import "kaeru.chan/voz/server"

func FindUserByID(ID string, users []server.TargetUser) *server.TargetUser {
	for _, user := range users {
		if user.ID == ID {
			return &user
		}
	}
	return nil
}

func StringContains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
