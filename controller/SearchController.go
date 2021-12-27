package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity"
	"picture_community/response"
	"picture_community/service"
	"strings"
)

func Search(c *gin.Context) {
	var u entity.SearchUsers

	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	if u.Page <= 0 || u.PageSize <= 0 {
		response.CheckFail(c, nil, "页码或数量有误")
		return
	}
	/*支持多个关键词搜索，关键词用空格隔开*/
	keywords := "%" + strings.Replace(u.NickName, " ", "%", -1) + "%"

	res := service.SearchService(keywords, u.PageSize, u.Page)

	response.HandleResponse(c, res)
}
