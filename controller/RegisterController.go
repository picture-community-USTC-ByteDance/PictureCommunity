package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func RegisterController(c *gin.Context) {
	var u _request.RegisterUser
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "Register invalid parameter")
		return
	}
	res := service.RegisterService(u)
	response.HandleResponse(c, res)
}

func IsUniqueController(c *gin.Context) {
	var u _request.IsUniqueInfo
	if err := c.ShouldBind(&u); err != nil {
		response.Fail(c, nil, "IsUnique invalid parameter")
	}

	//todo 验证用户名合法性

	res := service.IsUnique(u)
	response.HandleResponse(c, res)
}
