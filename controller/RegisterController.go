package controller

import (
	"github.com/gin-gonic/gin"
	"picture_community/entity/_request"
	"picture_community/response"
	"picture_community/service"
)

func RegisterController(c *gin.Context) {
	var u _request.RegisterUserInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "注册参数错误")
		return
	}
	isOK, message := service.RegisterService(u)
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}

func UsernameIsUniqueController(c *gin.Context) {
	var u _request.UsernameIsUniqueInfo
	if err := c.ShouldBind(&u); err != nil {
		response.CheckFail(c, nil, "检查用户名可用性参数错误")
		return
	}

	isOK, message := service.UsernameIsUniqueService(u)
	if isOK {
		response.Success(c, nil, message)
	} else {
		response.Fail(c, nil, message)
	}
	return
}
