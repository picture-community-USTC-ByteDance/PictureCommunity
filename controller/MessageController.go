package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
	"strconv"
)

func GetUserChatList(c *gin.Context) {
	var u _request.MessageRequest
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}
	uid, _ := c.Get("uid")
	count, totalPage, chatList := service.GetUserChatList(u, uid)

	if count == 0 {
		response.Success(c, nil, "没有与其他用户聊天")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "chatList": chatList}, "ok")
	}
}

func GetHistoryMsg(c *gin.Context) {
	ToId, _ := strconv.Atoi(c.Query("toid"))
	if uint(ToId) <= 0 {
		response.Fail(c, nil, "请求错误")
		return
	}

	uid, _ := c.Get("uid")
	mux.RLock()
	_, ok := Manager.Clients[uid.(uint)]
	mux.RUnlock()
	if !ok {
		response.Fail(c, nil, "未进行WS连接")
		return
	}

	mux.Lock()
	Manager.Clients[uid.(uint)].ToId = uint(ToId)
	mux.Unlock()
	count, userInfo, HistoryMsgList := service.GetHistoryMsg(uid.(uint), uint(ToId))

	if count == 0 {
		response.Success(c, nil, "没有与其他用户聊天")
	} else {
		response.Success(c, gin.H{"nickname": userInfo.Nickname, "profile": userInfo.Profile, "HistoryMsgList": HistoryMsgList}, "ok")
	}
}
func GetUnreadMsg(c *gin.Context) {

	uid, _ := c.Get("uid")
	totalUnreadMsg := service.GetUnreadMsg(uid.(uint))

	if totalUnreadMsg == 0 {
		response.Success(c, nil, "没有未读消息")
	} else {
		response.Success(c, gin.H{"totalUnreadMsg": totalUnreadMsg}, "ok")
	}
}
