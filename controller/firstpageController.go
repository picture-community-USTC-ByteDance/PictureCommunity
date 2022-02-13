package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func GetIdListController(c *gin.Context) {
	var u _request.Firstpage
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}
	isOk, IdList := service.GetPostIdList(u.Uid, u.Page, u.PageSize)
	if isOk {
		response.Success(c, IdList, "获取成功")
	} else {
		response.Success(c, nil, "该用户没有关注任何人")
	}
}
func GetDetailController(c *gin.Context) {
	var u _request.Firstpage
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}
	isOk, postList := service.GetPostDetail(u.Uid, u.Page, u.PageSize)
	if isOk {
		response.Success(c, postList, "获取成功")
	} else {
		response.Success(c, nil, "该用户没有关注任何人")
	}
}
