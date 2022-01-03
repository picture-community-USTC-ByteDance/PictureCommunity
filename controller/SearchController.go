package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func SearchUsers(c *gin.Context) {
	var u _request.SearchUsers

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}

	count, totalPage, searchUsers := service.SearchService(u)

	if count == 0 {
		response.Success(c, nil, "搜索用户不存在")
	} else {
		response.Success(c, gin.H{"totalpage": totalPage, "searchusers": searchUsers}, "ok")
	}

}
