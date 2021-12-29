package controller

import (
	"picture_community/utils"
)

func VerifyIDByToken(ID int64, token string) bool {
	user, err := utils.ParserToken(token)

	if user == nil || user.ID != ID || err != nil {
		return false
	}
	return true
}
