package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func GetIdListController(c *gin.Context) {
	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}
	var u _request.Firstpage
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}
	isOk, IdList := service.GetPostIdList(uid.(uint), u.Page, u.PageSize)
	if isOk {
		response.Success(c, IdList, "获取成功")
	} else {
		response.Success(c, nil, "该用户没有关注任何人")
	}
}
func GetDetailController(c *gin.Context) {
	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}
	var u _request.Firstpage
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.Fail(c, nil, "页码或数量有误")
		return
	}
	isOk, postList := service.GetPostDetail(uid.(uint), u.Page, u.PageSize)
	if isOk {
		response.Success(c, postList, "获取成功")
	} else {
		response.Success(c, nil, "该用户没有关注任何人!")
	}
}

/*得到单个帖子*/
func GetSinglePost(c *gin.Context) {
	var u _request.SinglePost
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Pid <= 0 {
		response.CheckFail(c, nil, "帖子id出错")
		return
	}
	post := service.GetSingleDetail(u.Pid, c)
	response.Success(c, post, "获取成功")
}
func GetPossibleFriends(c *gin.Context) {
	uid, exists := c.Get("uid")
	if exists == false {
		response.Fail(c, nil, "用户不存在！")
	}
	res := service.GetPossibleFriends(uid.(uint))
	if len(res) == 0 {
		response.Success(c, nil, "你已经关注了可能关注的所有人")
	} else {
		response.Success(c, res, "获取成功")
	}

}
