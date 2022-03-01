package service

import (
	"math"
	"picture_community/dao/user"
	"picture_community/entity/_request"
	"picture_community/entity/_response"
	"picture_community/entity/db"
)

func GetUserChatList(u _request.MessageRequest, uid interface{}) (int64, int, []_response.ResponseChatUsers) {
	count, userChatList := user.QueryChatListByUID(uid.(uint), u.Page, u.PageSize)

	totalPage := int(math.Ceil(float64(count) / float64(u.PageSize)))

	return count, totalPage, userChatList
}

func GetHistoryMsg(uid uint, ToId uint) (int64, db.UserDetail, []_response.ResponseHistoryMsg) {
	count, userInfo, HistoryMsgList := user.QueryHistoryMsg(uid, ToId)

	return count, userInfo, HistoryMsgList

}

func GetUnreadMsg(uid uint) int64 {
	count := user.QueryUnreadMsg(uid)

	return count

}
